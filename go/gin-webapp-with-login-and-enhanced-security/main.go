package main

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

var logger *slog.Logger

func init() {
	// Initialize structured logging
	logger = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	}))
	slog.SetDefault(logger)
}

// Security headers middleware
func securityHeadersMiddleware() gin.HandlerFunc {
	return gin.HandlerFunc(func(c *gin.Context) {
		// Content Security Policy (CSP)
		c.Header("Content-Security-Policy",
			"default-src 'self'; "+
				"style-src 'unsafe-inline'; "+
				"script-src 'self'; "+
				"img-src 'self' data:; "+
				"font-src 'self'; "+
				"connect-src 'self'; "+
				"frame-ancestors 'none'; "+
				"base-uri 'self'; "+
				"form-action 'self'")

		// Security headers
		c.Header("Referrer-Policy", "strict-origin-when-cross-origin")
		c.Header("Cross-Origin-Resource-Policy", "same-origin")
		c.Header("X-Content-Type-Options", "nosniff")
		c.Header("X-Frame-Options", "DENY")
		c.Header("X-XSS-Protection", "1; mode=block")
		c.Header("Strict-Transport-Security", "max-age=31536000; includeSubDomains")

		c.Next()
	})
}

func main() {
	logger.Info("Starting Gin server on port 8080")

	// Set Gin to release mode in production
	gin.SetMode(gin.ReleaseMode)

	r := gin.Default()

	// Request logging middleware
	r.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		logger.Info("Request",
			"method", param.Method,
			"path", param.Path,
			"status", param.StatusCode,
			"client_ip", param.ClientIP,
			"user_agent", param.Request.UserAgent(),
			"latency", param.Latency,
		)
		return ""
	}))

	// Apply security middleware
	r.Use(securityHeadersMiddleware())

	// Public routes
	r.GET("/", func(c *gin.Context) {
		c.Header("Content-Type", "text/html")
		c.String(http.StatusOK, homepage())
	})

	// Login route - GET shows form, POST processes it
	r.GET("/login", func(c *gin.Context) {
		c.Header("Content-Type", "text/html")
		c.String(http.StatusOK, loginPage())
	})
	r.POST("/login", loginUser)

	// Register route - GET shows form, POST processes it
	r.GET("/register", func(c *gin.Context) {
		c.Header("Content-Type", "text/html")
		c.String(http.StatusOK, registerPage())
	})
	r.POST("/register", registerUser)

	// Logout route
	r.POST("/logout", logoutUser)

	// Protected routes - Dashboard access
	r.GET("/protected-page", func(c *gin.Context) {
		// Check if user has valid session
		sessionToken, err := c.Cookie("session_token")
		if err != nil {
			// No session token, redirect to login
			c.Redirect(http.StatusSeeOther, "/login")
			return
		}

		// Find user with this session token
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
			// Invalid session, redirect to login
			c.Redirect(http.StatusSeeOther, "/login")
			return
		}

		// Valid session, show dashboard
		logger.Info("Dashboard accessed", "username", currentUser.Username, "client_ip", c.ClientIP())
		c.Header("Content-Type", "text/html")
		c.String(http.StatusOK, protectedPage(currentUser.Username, currentUser.CSRFToken))
	})

	// Start server
	logger.Info("Server starting on :8080")
	if err := r.Run(":8080"); err != nil {
		logger.Error("Server failed to start", "error", err.Error())
	}
}