package main

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"log"
	"net/http"
)

func generateToken(length int) string {
	bytes := make([]byte, length)
	if _, err := rand.Read(bytes); err != nil {
		log.Fatalf("Failed to generate token: %v", err)
	}

	return base64.URLEncoding.EncodeToString(bytes)
}

var AuthError = errors.New("Unauthorised")

func Authorise(r *http.Request) error {
	username := r.FormValue("username")
	user, ok := users[username]
	if !ok {
		return AuthError
	}

	// Get and validate session token from cookie
	st, err := r.Cookie("session_token")
	if err != nil || st.Value != user.SessionToken {
		return AuthError
	}

	// Get CSRF token from the header
	csrf := r.Header.Get("X-CSRF-Token")
	if csrf != user.CSRFToken || csrf == "" {
		return AuthError
	}

	return nil
}
