package main

import (
	"fmt"
	"net/http"
	"runtime/debug"
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
				app.logger.Printf("%s\n%s", err, debug.Stack())
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
