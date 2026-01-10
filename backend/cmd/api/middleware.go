package main

import (
	"crypto/subtle"
	"fmt"
	"net/http"
	"os"
	"runtime/debug"

	"github.com/google/uuid"
)

// TODO: PanicRecovery
// Goal: Prevent the entire server from crashing if a handler encounters a panic.
// Logic: Recover from panic, log the error using jsonlog, and send a 500 Internal Server Error.
// Priority: High (Add before Production)
func (app *application) recoverPanic(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Create a deferred function (which will always be run in the event
		// of a panic as Go unwinds the stack).
		defer func() {
			// Use the builtin recover function to check if there has been a
			// panic or not.
			if err := recover(); err != nil {
				// Set a "Connection: close" header on the response.
				w.Header().Set("Connection", "close")
				// Call the app.serverErrorResponse helper method to return a 500
				// Internal Server Error response.
				app.logger.Error("panic recovery", "error", err, "trace", string(debug.Stack()))
				app.serverErrorResponse(w, r, fmt.Errorf("%s", err))
			}
		}()

		next.ServeHTTP(w, r)
	})
}

// TODO: BodyLimit
// Goal: Protect the server from memory exhaustion attacks (e.g., uploading 1GB video).
// Logic: Wrap the request body in a MaxBytesReader (e.g., limit to 5MB).
// Priority: Medium (Add before Image Upload feature)
func (app *Application) enableBodyLimit(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// r.Body = http.MaxBytesReader(w, r.Body, 5 * 1024 * 1024)
		next.ServeHTTP(w, r)
	})
}

// Request ID Middleware
func (app *application) enableRequestID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := r.Header.Get("X-Request-Id")
		if id == "" {
			id = uuid.New().String()
		}

		w.Header().Set("X-Request-Id", id)
		next.ServeHTTP(w, app.contextSetRequestID(r, id))
	})
}

// Cache Control Middleware (1 Year)
func (app *application) cacheControl(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Cache-Control", "public, max-age=31536000, immutable")
		next.ServeHTTP(w, r)
	})
}

// Authentication Middleware (Constant Time Comparison)
func (app *application) requireAuthentication(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Get key from header
		apiKey := r.Header.Get("X-API-Key")

		// Get key from env (assume already loaded)
		expectedKey := os.Getenv("API_KEY")

		if expectedKey == "" {
			// If no key is set in env, maybe we should default strictly,
			// but for this demo let's log error and deny everything safely.
			app.serverErrorResponse(w, r, fmt.Errorf("server missing API_KEY configuration"))
			return
		}

		// Constant Time Comparison
		// Convert to bytes
		if subtle.ConstantTimeCompare([]byte(apiKey), []byte(expectedKey)) != 1 {
			app.JSONError(w, http.StatusUnauthorized, "Invalid or missing authentication token")
			return
		}

		next.ServeHTTP(w, r)
	})
}

// Max Body Size Middleware (10MB Global Limit)
func (app *application) bodyLimit(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 10MB limit for general safety (file uploads, etc)
		// Specific handlers like readJSON enforce tighter 1MB limits for JSON
		const maxBytes = 10 * 1024 * 1024
		r.Body = http.MaxBytesReader(w, r.Body, maxBytes)
		next.ServeHTTP(w, r)
	})
}

// Remove ID Headers (Fingerprint Removal)
func (app *application) removeIdentity(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Del("Server")
		w.Header().Del("X-Powered-By")
		next.ServeHTTP(w, r)
	})
}

// Security Headers Middleware
func (app *application) securityHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("X-Content-Type-Options", "nosniff")
		w.Header().Set("X-Frame-Options", "DENY")
		w.Header().Set("X-XSS-Protection", "1; mode=block")
		w.Header().Set("Referrer-Policy", "strict-origin-when-cross-origin")
		w.Header().Set("Content-Security-Policy", "default-src 'none'; frame-ancestors 'none'; sandbox")

		next.ServeHTTP(w, r)
	})
}

// Session Auth Middleware
func (app *application) requireLogin(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var token string

		// 1. Check for Bearer Token
		authHeader := r.Header.Get("Authorization")
		if len(authHeader) > 7 && authHeader[:7] == "Bearer " {
			token = authHeader[7:]
		}

		// 2. Fallback to Cookie
		if token == "" {
			cookie, err := r.Cookie("session_token")
			if err == nil {
				token = cookie.Value
			}
		}

		if token == "" {
			// No token provided via header or cookie
			app.JSONError(w, http.StatusUnauthorized, "Authentication required")
			return
		}

		// 3. Validate Session
		session, err := app.models.Sessions.Get(token)
		if err != nil {
			// Database error
			app.serverErrorResponse(w, r, err)
			return
		}

		if session == nil {
			// Invalid or expired token
			// Clear the cookie just in case it was a cookie
			http.SetCookie(w, &http.Cookie{
				Name:     "session_token",
				Value:    "",
				Path:     "/",
				MaxAge:   -1,
				HttpOnly: true,
			})
			app.JSONError(w, http.StatusUnauthorized, "Invalid or expired session")
			return
		}

		// Valid session -> Add user to context
		r = app.contextSetUser(r, session.UserID)

		next.ServeHTTP(w, r)
	})
}
