package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"sync"
	"time"

	"golang.org/x/time/rate"
)

type Login struct {
	HashedPassword string
	SessionToken   string
	CSRFToken      string
}

// Username is the key
type Users map[string]Login

var (
	users      = make(Users)
	usersMutex sync.RWMutex
	logger     *slog.Logger
	limiter    = rate.NewLimiter(rate.Every(time.Minute), 5) // 5 requests per minute
)

func init() {
	// Initialize structured logging
	logger = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	}))
	slog.SetDefault(logger)
}

// Rate limiting middleware
func rateLimitMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !limiter.Allow() {
			logSecurityEvent("rate_limit_exceeded", "", r)
			http.Error(w, "Rate limit exceeded", http.StatusTooManyRequests)
			return
		}
		next.ServeHTTP(w, r)
	}
}

// HTTPS enforcement middleware
func requireHTTPS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check if request is HTTPS (works with reverse proxies too)
		if r.Header.Get("X-Forwarded-Proto") != "https" && r.TLS == nil {
			httpsURL := "https://" + r.Host + r.RequestURI
			logger.Warn("HTTP request redirected to HTTPS",
				"original_url", r.URL.String(),
				"redirect_url", httpsURL,
				"client_ip", getClientIP(r))
			http.Redirect(w, r, httpsURL, http.StatusPermanentRedirect)
			return
		}
		next.ServeHTTP(w, r)
	})
}

// Security headers middleware
func securityHeadersMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Content Security Policy (CSP) - strict policy for this app
		w.Header().Set("Content-Security-Policy",
			"default-src 'self'; "+
				"style-src 'unsafe-inline'; "+
				"script-src 'self'; "+
				"img-src 'self' data:; "+
				"font-src 'self'; "+
				"connect-src 'self'; "+
				"frame-ancestors 'none'; "+
				"base-uri 'self'; "+
				"form-action 'self'")

		// Referrer Policy
		w.Header().Set("Referrer-Policy", "strict-origin-when-cross-origin")

		// Cross Origin Resource Policy
		w.Header().Set("Cross-Origin-Resource-Policy", "same-origin")

		// Existing security headers
		w.Header().Set("X-Content-Type-Options", "nosniff")
		w.Header().Set("X-Frame-Options", "DENY")
		w.Header().Set("X-XSS-Protection", "1; mode=block")
		w.Header().Set("Strict-Transport-Security", "max-age=31536000; includeSubDomains")

		next.ServeHTTP(w, r)
	})
}

func register(w http.ResponseWriter, r *http.Request) {
	clientIP := getClientIP(r)

	if r.Method != http.MethodPost {
		logger.Warn("Invalid method attempted on register endpoint",
			"method", r.Method,
			"client_ip", clientIP)
		http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
		return
	}

	username := r.FormValue("username")
	password := r.FormValue("password")

	// Input validation and sanitization
	if err := validateRegistrationInput(username, password); err != nil {
		logger.Warn("Registration validation failed",
			"username", username,
			"error", err.Error(),
			"client_ip", clientIP)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Clean the username
	username = sanitizeUsername(username)

	usersMutex.Lock()
	defer usersMutex.Unlock()

	if _, ok := users[username]; ok {
		logger.Warn("Registration attempt for existing user",
			"username", username,
			"client_ip", clientIP)
		http.Error(w, "User already exists", http.StatusConflict)
		return
	}

	hashedPassword, err := hashPassword(password)
	if err != nil {
		logger.Error("Password hashing failed",
			"username", username,
			"error", err.Error(),
			"client_ip", clientIP)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	users[username] = Login{
		HashedPassword: hashedPassword,
	}

	logger.Info("User registered successfully",
		"username", username,
		"client_ip", clientIP)

	// Redirect to login page after successful registration
	http.Redirect(w, r, "/login-page", http.StatusSeeOther)
}

func login(w http.ResponseWriter, r *http.Request) {
	clientIP := getClientIP(r)

	if r.Method != http.MethodPost {
		logger.Warn("Invalid method attempted on login endpoint",
			"method", r.Method,
			"client_ip", clientIP)
		http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
		return
	}

	username := r.FormValue("username")
	password := r.FormValue("password")

	// Input validation
	if err := validateLoginInput(username, password); err != nil {
		logger.Warn("Login validation failed",
			"username", username,
			"error", err.Error(),
			"client_ip", clientIP)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	username = sanitizeUsername(username)

	usersMutex.RLock()
	user, ok := users[username]
	usersMutex.RUnlock()

	if !ok || !checkPasswordMatch(user.HashedPassword, password) {
		logSecurityEvent("login_failed", username, r)
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return
	}

	sessionToken := generateToken(64)
	csrfToken := generateToken(64)

	// Set secure session cookie
	http.SetCookie(w, &http.Cookie{
		Name:     "session_token",
		Value:    sessionToken,
		Expires:  time.Now().Add(24 * time.Hour),
		HttpOnly: true,
		Secure:   true, // HTTPS only
		SameSite: http.SameSiteStrictMode,
		Path:     "/",
	})

	// Set CSRF token in a cookie
	http.SetCookie(w, &http.Cookie{
		Name:     "csrf_token",
		Value:    csrfToken,
		Expires:  time.Now().Add(24 * time.Hour),
		HttpOnly: false,
		Secure:   true,
		SameSite: http.SameSiteStrictMode,
		Path:     "/",
	})

	// Store session tokens
	usersMutex.Lock()
	user.SessionToken = sessionToken
	user.CSRFToken = csrfToken
	users[username] = user
	usersMutex.Unlock()

	logger.Info("User logged in successfully",
		"username", username,
		"client_ip", clientIP)

	// Redirect to protected page after successful login
	http.Redirect(w, r, "/protected-page", http.StatusSeeOther)
}

func logout(w http.ResponseWriter, r *http.Request) {
	clientIP := getClientIP(r)

	// Validate CSRF token from form
	formCSRF := r.FormValue("csrf_token")
	sessionCSRF := getUserCSRFToken(r)

	if formCSRF == "" || sessionCSRF == "" || formCSRF != sessionCSRF {
		logSecurityEvent("invalid_csrf_logout", "", r)
		http.Error(w, "Invalid CSRF token", http.StatusForbidden)
		return
	}

	// Clear cookies
	http.SetCookie(w, &http.Cookie{
		Name:     "session_token",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteStrictMode,
		Path:     "/",
	})

	http.SetCookie(w, &http.Cookie{
		Name:     "csrf_token",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HttpOnly: false,
		Secure:   true,
		SameSite: http.SameSiteStrictMode,
		Path:     "/",
	})

	// Get username from session instead of form
	username := getUsernameFromSession(r)
	if username == "" {
		logSecurityEvent("logout_session_error", "", r)
		http.Error(w, "Session error", http.StatusInternalServerError)
		return
	}

	// Clear tokens from storage
	usersMutex.Lock()
	if user, exists := users[username]; exists {
		user.SessionToken = ""
		user.CSRFToken = ""
		users[username] = user
	}
	usersMutex.Unlock()

	logger.Info("User logged out successfully",
		"username", username,
		"client_ip", clientIP)

	// Redirect to homepage after successful logout
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func protected(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		logger.Warn("Invalid method attempted on protected endpoint",
			"method", r.Method,
			"client_ip", getClientIP(r))
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	if err := Authorise(r); err != nil {
		logSecurityEvent("unauthorized_protected_access", "", r)
		http.Error(w, "Unauthorised", http.StatusUnauthorized)
		return
	}

	username := r.FormValue("username")
	username = sanitizeUsername(username)

	logger.Info("Protected endpoint accessed",
		"username", username,
		"client_ip", getClientIP(r))

	fmt.Fprintf(w, "CSRF validated. Welcome, %s", username)
}

// Homepage handler that serves HTML from webpages.go
func homepageHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, homepage())
}

// Login page handler
func loginPageHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
		return
	}

	// If user is already authenticated, redirect to protected page
	if isAuthenticated(r) {
		http.Redirect(w, r, "/protected-page", http.StatusSeeOther)
		return
	}

	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, loginPage())
}

// Register page handler
func registerPageHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
		return
	}

	// If user is already authenticated, redirect to protected page
	if isAuthenticated(r) {
		http.Redirect(w, r, "/protected-page", http.StatusSeeOther)
		return
	}

	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, registerPage())
}

// Check if user has a valid session (for page viewing, no CSRF required)
func isAuthenticated(r *http.Request) bool {
	// Get session token from cookie
	sessionCookie, err := r.Cookie("session_token")
	if err != nil || sessionCookie.Value == "" {
		return false
	}

	// Check if any user has this session token
	usersMutex.RLock()
	defer usersMutex.RUnlock()

	for _, user := range users {
		if user.SessionToken == sessionCookie.Value && user.SessionToken != "" {
			return true
		}
	}

	return false
}

// Get CSRF token for authenticated user
func getUserCSRFToken(r *http.Request) string {
	// Get session token from cookie
	sessionCookie, err := r.Cookie("session_token")
	if err != nil || sessionCookie.Value == "" {
		return ""
	}

	// Find user with this session token
	usersMutex.RLock()
	defer usersMutex.RUnlock()

	for _, user := range users {
		if user.SessionToken == sessionCookie.Value && user.SessionToken != "" {
			return user.CSRFToken
		}
	}

	return ""
}

// Get username from session token
func getUsernameFromSession(r *http.Request) string {
	// Get session token from cookie
	sessionCookie, err := r.Cookie("session_token")
	if err != nil || sessionCookie.Value == "" {
		return ""
	}

	// Find user with this session token
	usersMutex.RLock()
	defer usersMutex.RUnlock()

	for username, user := range users {
		if user.SessionToken == sessionCookie.Value && user.SessionToken != "" {
			return username
		}
	}

	return ""
}

// Protected page handler
func protectedPageHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
		return
	}

	// Check if user is authenticated
	if !isAuthenticated(r) {
		// Redirect to login page if not authenticated
		http.Redirect(w, r, "/login-page", http.StatusSeeOther)
		return
	}

	// Get CSRF token for the logged-in user
	csrfToken := getUserCSRFToken(r)
	if csrfToken == "" {
		http.Error(w, "Session error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, protectedPage(csrfToken))
}

func main() {
	logger.Info("Starting server on port 8101")

	// Create a new ServeMux for better control over routing
	mux := http.NewServeMux()

	// Homepage route
	mux.HandleFunc("/", homepageHandler)

	// Page routes (GET)
	mux.HandleFunc("/login-page", loginPageHandler)
	mux.HandleFunc("/register-page", registerPageHandler)
	mux.HandleFunc("/protected-page", protectedPageHandler)

	// Apply rate limiting to all endpoints
	mux.HandleFunc("/register", rateLimitMiddleware(register))
	mux.HandleFunc("/login", rateLimitMiddleware(login))
	mux.HandleFunc("/logout", rateLimitMiddleware(logout))
	mux.HandleFunc("/protected", rateLimitMiddleware(protected))

	// Apply security middleware
	handler := securityHeadersMiddleware(requireHTTPS(mux))

	server := &http.Server{
		Addr:         ":8101",
		Handler:      handler,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	logger.Info("Server starting on :8101")
	if err := server.ListenAndServe(); err != nil {
		logger.Error("Server failed to start", "error", err.Error())
	}
}
