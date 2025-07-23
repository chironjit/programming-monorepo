package main

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"net/http"
	"regexp"
	"strings"
	"sync"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Username     string
	PasswordHash string
	SessionToken string
	CSRFToken    string
}

var (
	users      = make(map[string]*User)
	usersMutex sync.RWMutex
)

func generateToken(length int) string {
	bytes := make([]byte, length)
	if _, err := rand.Read(bytes); err != nil {
		logger.Error("Failed to generate token", "error", err.Error())
		return ""
	}
	return base64.URLEncoding.EncodeToString(bytes)
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func checkPassword(hash, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// Input validation functions
func validateUsername(username string) error {
	username = strings.TrimSpace(username)
	
	if len(username) < 3 {
		return errors.New("username must be at least 3 characters long")
	}
	
	if len(username) > 50 {
		return errors.New("username must be no more than 50 characters long")
	}
	
	// Allow only alphanumeric characters, underscores, and hyphens
	validUsername := regexp.MustCompile(`^[a-zA-Z0-9_-]+$`)
	if !validUsername.MatchString(username) {
		return errors.New("username can only contain letters, numbers, underscores, and hyphens")
	}
	
	return nil
}

func validatePassword(password string) error {
	if len(password) < 8 {
		return errors.New("password must be at least 8 characters long")
	}
	
	if len(password) > 128 {
		return errors.New("password must be no more than 128 characters long")
	}
	
	// Check for at least one uppercase letter
	hasUpper := regexp.MustCompile(`[A-Z]`).MatchString(password)
	// Check for at least one lowercase letter
	hasLower := regexp.MustCompile(`[a-z]`).MatchString(password)
	// Check for at least one digit
	hasNumber := regexp.MustCompile(`[0-9]`).MatchString(password)
	
	if !hasUpper || !hasLower || !hasNumber {
		return errors.New("password must contain at least one uppercase letter, one lowercase letter, and one number")
	}
	
	return nil
}

func registerUser(c *gin.Context) {
	username := strings.TrimSpace(c.PostForm("username"))
	password := c.PostForm("password")

	if username == "" || password == "" {
		logger.Warn("Registration failed - missing credentials", "client_ip", c.ClientIP())
		c.String(http.StatusBadRequest, "Username and password are required")
		return
	}

	// Validate username
	if err := validateUsername(username); err != nil {
		logger.Warn("Registration failed - invalid username", "username", username, "error", err.Error(), "client_ip", c.ClientIP())
		c.String(http.StatusBadRequest, "Invalid username: "+err.Error())
		return
	}

	// Validate password
	if err := validatePassword(password); err != nil {
		logger.Warn("Registration failed - invalid password", "username", username, "error", err.Error(), "client_ip", c.ClientIP())
		c.String(http.StatusBadRequest, "Invalid password: "+err.Error())
		return
	}

	usersMutex.Lock()
	defer usersMutex.Unlock()

	if _, exists := users[username]; exists {
		logger.Warn("Registration failed - user already exists", "username", username, "client_ip", c.ClientIP())
		c.Redirect(http.StatusSeeOther, "/register")
		return
	}

	hashedPassword, err := hashPassword(password)
	if err != nil {
		logger.Error("Registration failed - password hashing error", "username", username, "client_ip", c.ClientIP(), "error", err.Error())
		c.Redirect(http.StatusSeeOther, "/register")
		return
	}

	users[username] = &User{
		Username:     username,
		PasswordHash: hashedPassword,
	}

	logger.Info("User registered successfully", "username", username, "client_ip", c.ClientIP())

	// After successful registration, log them in automatically
	sessionToken := generateToken(32)
	csrfToken := generateToken(32)

	users[username].SessionToken = sessionToken
	users[username].CSRFToken = csrfToken

	// Set session cookie with secure flag for HTTPS
	c.SetCookie("session_token", sessionToken, 3600*24, "/", "", true, true)
	c.SetCookie("csrf_token", csrfToken, 3600*24, "/", "", true, false)

	c.Redirect(http.StatusSeeOther, "/protected-page")
}

func loginUser(c *gin.Context) {
	username := strings.TrimSpace(c.PostForm("username"))
	password := c.PostForm("password")

	if username == "" || password == "" {
		logger.Warn("Login failed - missing credentials", "client_ip", c.ClientIP())
		c.String(http.StatusBadRequest, "Username and password are required")
		return
	}

	// Basic validation for login (less strict than registration)
	if len(username) > 50 || len(password) > 128 {
		logger.Warn("Login failed - input too long", "username", username, "client_ip", c.ClientIP())
		c.String(http.StatusBadRequest, "Invalid input length")
		return
	}

	usersMutex.RLock()
	user, exists := users[username]
	usersMutex.RUnlock()

	if !exists || !checkPassword(user.PasswordHash, password) {
		logger.Warn("Login failed", "username", username, "client_ip", c.ClientIP())
		c.Redirect(http.StatusSeeOther, "/login")
		return
	}

	sessionToken := generateToken(32)
	csrfToken := generateToken(32)

	usersMutex.Lock()
	user.SessionToken = sessionToken
	user.CSRFToken = csrfToken
	usersMutex.Unlock()

	// Set session cookies with secure flag for HTTPS
	c.SetCookie("session_token", sessionToken, 3600*24, "/", "", true, true)
	c.SetCookie("csrf_token", csrfToken, 3600*24, "/", "", true, false)

	logger.Info("User logged in successfully", "username", username, "client_ip", c.ClientIP())
	c.Redirect(http.StatusSeeOther, "/protected-page")
}

func logoutUser(c *gin.Context) {
	sessionToken, err := c.Cookie("session_token")
	if err != nil {
		c.Redirect(http.StatusSeeOther, "/")
		return
	}

	// Validate CSRF token
	submittedCSRFToken := c.PostForm("csrf_token")
	if submittedCSRFToken == "" {
		logger.Warn("Logout failed - missing CSRF token", "client_ip", c.ClientIP())
		c.String(http.StatusBadRequest, "Missing CSRF token")
		return
	}

	// Find user and validate CSRF token
	usersMutex.Lock()
	var username string
	var currentUser *User
	for _, user := range users {
		if user.SessionToken == sessionToken {
			currentUser = user
			username = user.Username
			break
		}
	}
	
	if currentUser == nil {
		usersMutex.Unlock()
		logger.Warn("Logout failed - invalid session", "client_ip", c.ClientIP())
		c.Redirect(http.StatusSeeOther, "/")
		return
	}
	
	// Validate CSRF token matches user's stored token
	if currentUser.CSRFToken != submittedCSRFToken {
		usersMutex.Unlock()
		logger.Warn("Logout failed - invalid CSRF token", "username", username, "client_ip", c.ClientIP())
		c.String(http.StatusForbidden, "Invalid CSRF token")
		return
	}
	
	// Clear session and CSRF token
	currentUser.SessionToken = ""
	currentUser.CSRFToken = ""
	usersMutex.Unlock()

	// Clear cookies with secure flag for HTTPS
	c.SetCookie("session_token", "", -1, "/", "", true, true)
	c.SetCookie("csrf_token", "", -1, "/", "", true, false)

	logger.Info("User logged out", "username", username, "client_ip", c.ClientIP())
	c.Redirect(http.StatusSeeOther, "/")
}

func requireAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		sessionToken, err := c.Cookie("session_token")
		if err != nil {
			c.Redirect(http.StatusSeeOther, "/login")
			c.Abort()
			return
		}

		usersMutex.RLock()
		var currentUser *User
		for _, user := range users {
			if user.SessionToken == sessionToken && sessionToken != "" {
				currentUser = user
				break
			}
		}
		usersMutex.RUnlock()

		if currentUser == nil {
			c.Redirect(http.StatusSeeOther, "/login")
			c.Abort()
			return
		}

		c.Set("user", currentUser)
		c.Next()
	}
}

func getCurrentUser(c *gin.Context) (*User, error) {
	user, exists := c.Get("user")
	if !exists {
		return nil, errors.New("user not found in context")
	}
	return user.(*User), nil
}
