package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/cconner57/adoption-os/backend/internal/data"
)

// subscribeHandler accepts a push subscription from the frontend
func (app *application) subscribeHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Endpoint string `json:"endpoint"`
		Keys     struct {
			P256dh string `json:"p256dh"`
			Auth   string `json:"auth"`
		} `json:"keys"`
		UserAgent string `json:"user_agent"`
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	sub := &data.NotificationSubscription{
		Endpoint:  input.Endpoint,
		P256dh:    input.Keys.P256dh,
		Auth:      input.Keys.Auth,
		UserAgent: r.UserAgent(),
	}

	err := app.models.Notifications.Insert(sub)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	// Send a welcome test notification only to the subscriber
	go app.notifier.Send(sub, "Notifications connected! You will now receive alerts.")

	// Respond with the public key
	json.NewEncoder(w).Encode(map[string]string{
		"vapid_public_key": os.Getenv("VAPID_PUBLIC_KEY"),
	})
}

// testNotificationHandler sends a test message to the specific endpoint request
func (app *application) testNotificationHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Endpoint string `json:"endpoint"`
		Keys     struct {
			P256dh string `json:"p256dh"`
			Auth   string `json:"auth"`
		} `json:"keys"`
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	sub := &data.NotificationSubscription{
		Endpoint: input.Endpoint,
		P256dh:   input.Keys.P256dh,
		Auth:     input.Keys.Auth,
	}

	go app.notifier.Send(sub, "Test Notification: 🔔 Ding Dong! It works!")

	app.writeJSON(w, http.StatusOK, envelope{"message": "Test notification sent"}, nil)
}

func (app *application) getPublicKeyHandler(w http.ResponseWriter, r *http.Request) {
	key := app.notifier.VapidKeys.PublicKey
	if key == "" {
		app.serverErrorResponse(w, r, fmt.Errorf("vapid public key not configured"))
		return
	}
	app.writeJSON(w, http.StatusOK, envelope{"public_key": key}, nil)
}

// sendBroadcastHandler (Admin) sends a message to all subscribers
func (app *application) sendBroadcastHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Message string `json:"message"`
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	go app.notifier.SendToAll(input.Message)

	app.writeJSON(w, http.StatusOK, envelope{"status": "queued"}, nil)
}
