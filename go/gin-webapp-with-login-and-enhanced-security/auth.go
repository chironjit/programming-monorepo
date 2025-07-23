package main

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"net/http"
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

func registerUser(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	if username == "" || password == "" {
		c.Redirect(http.StatusSeeOther, "/register")
		return
	}

	usersMutex.Lock()
	defer usersMutex.Unlock()

	if _, exists := users[username]; exists {
		c.Redirect(http.StatusSeeOther, "/register")
		return
	}

	hashedPassword, err := hashPassword(password)
	if err != nil {
		logger.Error("Password hashing failed", "error", err.Error())
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

	// Set session cookie
	c.SetCookie("session_token", sessionToken, 3600*24, "/", "", false, true)
	c.SetCookie("csrf_token", csrfToken, 3600*24, "/", "", false, false)

	c.Redirect(http.StatusSeeOther, "/protected-page")
}

func loginUser(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	if username == "" || password == "" {
		c.Redirect(http.StatusSeeOther, "/login")
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

	// Set session cookies
	c.SetCookie("session_token", sessionToken, 3600*24, "/", "", false, true)
	c.SetCookie("csrf_token", csrfToken, 3600*24, "/", "", false, false)

	logger.Info("User logged in successfully", "username", username, "client_ip", c.ClientIP())
	c.Redirect(http.StatusSeeOther, "/protected-page")
}

func logoutUser(c *gin.Context) {
	sessionToken, err := c.Cookie("session_token")
	if err != nil {
		c.Redirect(http.StatusSeeOther, "/")
		return
	}

	// Find user and clear session
	usersMutex.Lock()
	for _, user := range users {
		if user.SessionToken == sessionToken {
			user.SessionToken = ""
			user.CSRFToken = ""
			break
		}
	}
	usersMutex.Unlock()

	// Clear cookies
	c.SetCookie("session_token", "", -1, "/", "", false, true)
	c.SetCookie("csrf_token", "", -1, "/", "", false, false)

	logger.Info("User logged out", "client_ip", c.ClientIP())
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
