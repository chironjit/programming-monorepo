## Introduction

Demo webapp with login and protected routes that uses the gin framework. Implements

- Rate Limiting: 5 requests per minute per IP
- Input Validation: Strong password requirements, username sanitization
- Structured Logging: JSON format with security event tracking
- HTTPS Enforcement: Automatic redirect to HTTPS
- Security Headers: Added common security headers
- Concurrency Safety: Added mutex protection for user map
- Enhanced Cookie Security: Secure, SameSite, proper expiration
- Better Error Handling: More specific error messages and logging

Does not use databases to maintain simplicity and focus on the web aspects.

Use `go run .` to run webapp

Temporarily disable HTTPS enforcement by commenting out the `requireHTTPS` middleware in the main function for local testing.
