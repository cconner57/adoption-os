package main

import (
	"errors"
	"net/http"
	"time"

	"github.com/cconner57/adoption-os/backend/internal/data"
	"github.com/cconner57/adoption-os/backend/internal/password"
)

func (app *application) loginUserHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	// 1. Lookup user by email
	user, err := app.models.Users.GetByEmail(input.Email)
	if err != nil {
		if errors.Is(err, data.ErrRecordNotFound) {
			app.invalidCredentialsResponse(w, r)
			return
		}
		app.serverErrorResponse(w, r, err)
		return
	}

	// 2. Verify password
	match, err := password.CheckPassword(input.Password, user.PasswordHash)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	if !match {
		app.invalidCredentialsResponse(w, r)
		return
	}

	// 3. Create Session
	// 24 hour expiry
	ttl := 24 * time.Hour
	token, err := app.models.Sessions.Insert(user.ID, ttl, r.RemoteAddr, r.UserAgent())
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	// 4. Set Cookie
	http.SetCookie(w, &http.Cookie{
		Name:     "session_token",
		Value:    token,
		Path:     "/",
		Expires:  time.Now().Add(ttl),
		HttpOnly: true,
		Secure:   true, // Requirement: "Secure"
		SameSite: http.SameSiteStrictMode,
	})

	// 5. Return 200 OK
	// We can return user info if needed, but user just asked for 200 OK or token.
	// Since cookie is set, just OK is fine.
	err = app.writeJSON(w, http.StatusOK, envelope{"message": "authentication successful"}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
