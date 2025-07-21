package main

import (
	"errors"
	"regexp"
	"strings"
	"unicode"
)

var (
	usernameRegex      = regexp.MustCompile(`^[a-zA-Z0-9_]+$`)
	ErrInvalidUsername = errors.New("username must be 6-50 characters, alphanumeric and underscore only")
	ErrInvalidPassword = errors.New("password must be at least 12 characters with mixed case, numbers, and special characters")
	ErrEmptyInput      = errors.New("username and password cannot be empty")
)

func validateRegistrationInput(username, password string) error {
	if strings.TrimSpace(username) == "" || strings.TrimSpace(password) == "" {
		return ErrEmptyInput
	}

	if err := validateUsername(username); err != nil {
		return err
	}

	if err := validatePassword(password); err != nil {
		return err
	}

	return nil
}

func validateLoginInput(username, password string) error {
	if strings.TrimSpace(username) == "" || strings.TrimSpace(password) == "" {
		return ErrEmptyInput
	}

	// For login, we just check basic format, not complexity
	if len(username) < 6 || len(username) > 50 {
		return ErrInvalidUsername
	}

	if len(password) < 12 {
		return errors.New("invalid credentials")
	}

	return nil
}

func validateUsername(username string) error {
	username = strings.TrimSpace(username)

	if len(username) < 6 || len(username) > 50 {
		return ErrInvalidUsername
	}

	if !usernameRegex.MatchString(username) {
		return ErrInvalidUsername
	}

	return nil
}

func validatePassword(password string) error {
	if len(password) < 12 {
		return ErrInvalidPassword
	}

	var (
		hasUpper   = false
		hasLower   = false
		hasNumber  = false
		hasSpecial = false
	)

	for _, char := range password {
		switch {
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsNumber(char):
			hasNumber = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			hasSpecial = true
		}
	}

	if !hasUpper || !hasLower || !hasNumber || !hasSpecial {
		return ErrInvalidPassword
	}

	return nil
}

func sanitizeUsername(username string) string {
	return strings.ToLower(strings.TrimSpace(username))
}
