package main

import (
	"errors"
	"net/http"
	"time"

	"github.com/cconner57/adoption-os/backend/internal/data"
	"github.com/cconner57/adoption-os/backend/internal/password"
)

func (app *application) registerUserHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	// Validate input (basic)
	if input.Name == "" || input.Email == "" || input.Password == "" {
		app.badRequestResponse(w, r, errors.New("name, email, and password are required"))
		return
	}

	// Hash password
	hash, err := password.HashPassword(input.Password)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	user := &data.User{
		Name:         input.Name,
		Email:        input.Email,
		PasswordHash: hash,
		Activated:    true, // Auto-activate for now
	}

	err = app.models.Users.Insert(user)
	if err != nil {
		// Checks for duplicate email could happen here
		app.serverErrorResponse(w, r, err)
		return
	}

	app.writeJSON(w, http.StatusCreated, envelope{"user": user}, nil)
}

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

func (app *application) profileUserHandler(w http.ResponseWriter, r *http.Request) {
	id := app.contextGetUser(r)

	user, err := app.models.Users.Get(id)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	app.JSONResponse(w, http.StatusOK, user)
}
