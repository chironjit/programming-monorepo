package main

import (
	"fmt"
	"net/http"
	"time"
)

type Login struct {
	HashedPassword string
	SessionToken   string
	CSRFToken      string
}

// Username is the key
type Users map[string]Login

var users = make(Users)

// Key functions
func register(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		err := http.StatusMethodNotAllowed
		http.Error(w, "Invalid method", err)
		return
	}

	username := r.FormValue("username")
	password := r.FormValue("password")
	if len(username) < 6 || len(password) < 16 {
		err := http.StatusNotAcceptable
		http.Error(w, "Invalid username/password", err)
		return
	}

	if _, ok := users[username]; ok {
		err := http.StatusConflict
		http.Error(w, "User already exists", err)
		return
	}

	hashedPassword, _ := hashPassword(password)
	users[username] = Login{
		HashedPassword: hashedPassword,
	}

	fmt.Fprint(w, "Registration successful")
}

func login(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		err := http.StatusMethodNotAllowed
		http.Error(w, "Invalid method", err)
		return
	}

	username := r.FormValue("username")
	password := r.FormValue("password")

	user, ok := users[username]
	if !ok || !checkPasswordMatch(user.HashedPassword, password) {
		err := http.StatusUnauthorized
		http.Error(w, "Invalid username or password", err)
		return
	}

	sessionToken := generateToken(64)
	csrfToken := generateToken(64)

	// Set session cookie
	http.SetCookie(w, &http.Cookie{
		Name:     "session_token",
		Value:    sessionToken,
		Expires:  time.Now().Add(24 * time.Hour),
		HttpOnly: true,
	})

	// Set CSRF token in a cookie
	http.SetCookie(w, &http.Cookie{
		Name:     "csrf_token",
		Value:    csrfToken,
		Expires:  time.Now().Add(24 * time.Hour),
		HttpOnly: false,
	})

	// Store session tokens
	user.SessionToken = sessionToken
	user.CSRFToken = csrfToken
	users[username] = user

	fmt.Fprintln(w, "Login successful")

}

func logout(w http.ResponseWriter, r *http.Request) {
	if err := Authorise(r); err != nil {
		err := http.StatusUnauthorized
		http.Error(w, "Unauthorised", err)
		return
	}

	// Clear cookie
	http.SetCookie(w, &http.Cookie{
		Name:     "session_token",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HttpOnly: true,
	})

	http.SetCookie(w, &http.Cookie{
		Name:     "csrf_token",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HttpOnly: false,
	})

	//Clear tokens from db
	username := r.FormValue("username")
	user, _ := users[username]
	user.SessionToken = ""
	user.CSRFToken = ""
	users[username] = user

	fmt.Fprintln(w, "Logged out successfully")
}

func protected(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		err := http.StatusMethodNotAllowed
		http.Error(w, "Invalid request method", err)
		return
	}

	if err := Authorise(r); err != nil {
		err := http.StatusUnauthorized
		http.Error(w, "Unauthorised", err)
		return
	}

	username := r.FormValue("username")
	fmt.Fprintf(w, "CSRF validated. Welcome, %s", username)
}

// Main function
func main() {

	// Endpoints
	http.HandleFunc("/register", register)
	http.HandleFunc("/login", login)
	http.HandleFunc("/logout", logout)
	http.HandleFunc("/protected", protected)
	http.ListenAndServe(":8100", nil)

}
