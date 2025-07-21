package main

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"log"
	"net/http"
	"strings"
)

func generateToken(length int) string {
	bytes := make([]byte, length)
	if _, err := rand.Read(bytes); err != nil {
		log.Fatalf("Failed to generate token: %v", err)
	}

	return base64.URLEncoding.EncodeToString(bytes)
}

var AuthError = errors.New("unauthorised")

func Authorise(r *http.Request) error {
	username := sanitizeUsername(r.FormValue("username"))

	usersMutex.RLock()
	user, ok := users[username]
	usersMutex.RUnlock()

	if !ok {
		return AuthError
	}

	// Get and validate session token from cookie
	st, err := r.Cookie("session_token")
	if err != nil || st.Value != user.SessionToken || user.SessionToken == "" {
		return AuthError
	}

	// Get CSRF token from the header
	csrf := r.Header.Get("X-CSRF-Token")
	if csrf != user.CSRFToken || csrf == "" {
		return AuthError
	}

	return nil
}

func logSecurityEvent(event string, username string, r *http.Request) {
	logger.Warn("Security event",
		"event", event,
		"username", username,
		"client_ip", getClientIP(r),
		"user_agent", r.UserAgent(),
		"url", r.URL.String(),
		"method", r.Method,
	)
}

func getClientIP(r *http.Request) string {
	// Check X-Forwarded-For header first (for reverse proxies)
	xForwardedFor := r.Header.Get("X-Forwarded-For")
	if xForwardedFor != "" {
		// Take the first IP if multiple are present
		ips := strings.Split(xForwardedFor, ",")
		return strings.TrimSpace(ips[0])
	}

	// Check X-Real-IP header
	xRealIP := r.Header.Get("X-Real-IP")
	if xRealIP != "" {
		return xRealIP
	}

	// Fall back to RemoteAddr
	ip := r.RemoteAddr
	if colonPos := strings.LastIndex(ip, ":"); colonPos != -1 {
		ip = ip[:colonPos]
	}

	return ip
}
