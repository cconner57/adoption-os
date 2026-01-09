package main

import (
	"context"
	"net/http"
	"time"
)

// Change 'Home' to 'home' (lowercase is standard for internal handlers)
func (app *application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	// Create a context with a 1-second timeout for the DB ping
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	// Check DB status
	err := app.db.PingContext(ctx)
	if err != nil {
		data := envelope{
			"status": "error",
			"db":     "down",
		}
		// Return 503 Service Unavailable
		app.writeJSON(w, http.StatusServiceUnavailable, data, nil)
		return
	}

	data := envelope{
		"status": "ok",
		"uptime": time.Since(time.Now()).String(), // Placeholder, realistically start time should be stored in app struct
		"system_info": map[string]string{
			"environment": app.config.env,
			"version":     "1.0.0",
		},
	}

	err = app.writeJSON(w, http.StatusOK, data, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
