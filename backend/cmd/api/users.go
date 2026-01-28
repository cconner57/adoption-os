package main

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/cconner57/adoption-os/backend/internal/data"
	"github.com/cconner57/adoption-os/backend/internal/password"
	"github.com/google/uuid"
)

func (app *application) registerUserHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
		Token    string `json:"token"`
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

	// Determine Role
	// Default to VOLUNTEER_1 (Level 20)
	role := "VOLUNTEER_1"

	// If token is provided, validate it against DB
	if input.Token != "" {
		inviteRole, err := app.verifyAndConsumeInvite(input.Token, input.Email)
		if err != nil {
			app.JSONError(w, http.StatusBadRequest, err.Error())
			return
		}
		role = inviteRole
	}

	user := &data.User{
		Name:         input.Name,
		Email:        input.Email,
		PasswordHash: hash,
		Activated:    true, // Auto-activate for now
		Role:         role,
	}

	err = app.models.Users.Insert(user)
	if err != nil {
		// Checks for duplicate email could happen here
		app.serverErrorResponse(w, r, err)
		return
	}

	app.writeJSON(w, http.StatusCreated, envelope{"user": user}, nil)
}

func (app *application) verifyAndConsumeInvite(token, email string) (string, error) {
	invite, err := app.models.Invitations.Get(token)
	if err != nil {
		if errors.Is(err, data.ErrRecordNotFound) {
			return "", errors.New("invalid invite token")
		}
		return "", err
	}

	// Check expiry
	if time.Now().After(invite.ExpiresAt) {
		return "", errors.New("invite token has expired")
	}

	// Verify email matches (security check)
	if invite.Email != email {
		return "", errors.New("email does not match invite")
	}

	// Delete invite after use
	// We handle the delete error by logging it but not failing the registration
	err = app.models.Invitations.Delete(token)
	if err != nil {
		app.logger.Error("Failed to delete used invite", "token", token, "error", err)
	}

	return invite.Role, nil
}

func (app *application) validateInviteHandler(w http.ResponseWriter, r *http.Request) {
	token := r.PathValue("token")

	if token == "" {
		app.badRequestResponse(w, r, errors.New("missing token"))
		return
	}

	invite, err := app.models.Invitations.Get(token)
	if err != nil {
		if errors.Is(err, data.ErrRecordNotFound) {
			app.JSONError(w, http.StatusNotFound, "Invalid invite token")
			return
		}
		app.serverErrorResponse(w, r, err)
		return
	}

	if time.Now().After(invite.ExpiresAt) {
		app.JSONError(w, http.StatusBadRequest, "Invite token has expired")
		return
	}

	app.writeJSON(w, http.StatusOK, envelope{"invite": invite}, nil)
}

func (app *application) inviteUserHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Email string `json:"email"`
		Role  string `json:"role"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	if input.Email == "" || input.Role == "" {
		app.badRequestResponse(w, r, errors.New("email and role are required"))
		return
	}

	// Generate Token
	token := uuid.New().String()

	invite := &data.Invitation{
		Token:     token,
		Email:     input.Email,
		Role:      input.Role,
		ExpiresAt: time.Now().Add(48 * time.Hour),
	}

	err = app.models.Invitations.Insert(invite)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	// Send Email (Mock)
	app.logger.Info("INVITE LINK GENERATED", "email", input.Email, "link", fmt.Sprintf("https://myapp.com/register?token=%s", token))

	app.writeJSON(w, http.StatusCreated, envelope{"invitation": invite}, nil)
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
		app.logger.Info("Login failed: User lookup error", "email", input.Email, "error", err.Error())
		if errors.Is(err, data.ErrRecordNotFound) {
			app.invalidCredentialsResponse(w, r)
			return
		}
		app.serverErrorResponse(w, r, err)
		return
	}

	app.logger.Info("Login: User found", "id", user.ID, "activated", user.Activated)

	// 2. Verify password
	match, err := password.CheckPassword(input.Password, user.PasswordHash)
	if err != nil {
		app.logger.Error("Login: Password check error", "error", err)
		app.serverErrorResponse(w, r, err)
		return
	}

	if !match {
		app.logger.Info("Login failed: Password mismatch", "email", input.Email)
		app.invalidCredentialsResponse(w, r)
		return
	}

	app.logger.Info("Login successful", "email", input.Email)

	// 3. Create Session
	// 24 hour expiry
	ttl := 24 * time.Hour
	token, err := app.models.Sessions.Insert(user.ID, ttl, r.RemoteAddr, r.UserAgent())
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	// 4. Set Cookie
	cookieMode := http.SameSiteLaxMode
	isSecure := false

	if app.config.env == "production" {
		cookieMode = http.SameSiteNoneMode
		isSecure = true
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "session_token",
		Value:    token,
		Path:     "/",
		Expires:  time.Now().Add(ttl),
		HttpOnly: true,
		Secure:   isSecure,
		SameSite: cookieMode,
	})

	// 5. Return 200 OK + Token/User
	// We return the token so the frontend can use it in Authorization header if cookies fail
	err = app.writeJSON(w, http.StatusOK, envelope{
		"message": "authentication successful",
		"token":   token,
		"user":    user,
	}, nil)
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

func (app *application) updateUserHandler(w http.ResponseWriter, r *http.Request) {
	userId := app.contextGetUser(r)

	user, err := app.models.Users.Get(userId)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	var input struct {
		Name     *string `json:"name"`
		Email    *string `json:"email"`
		Password *string `json:"password"`
	}

	err = app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	if input.Name != nil {
		user.Name = *input.Name
	}

	if input.Email != nil {
		user.Email = *input.Email
	}

	if input.Password != nil && *input.Password != "" {
		hash, err := password.HashPassword(*input.Password)
		if err != nil {
			app.serverErrorResponse(w, r, err)
			return
		}
		user.PasswordHash = hash
	}

	err = app.models.Users.Update(user)
	if err != nil {
		if errors.Is(err, data.ErrEditConflict) {
			app.editConflictResponse(w, r)
		} else {
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	app.writeJSON(w, http.StatusOK, envelope{"user": user}, nil)
}

func (app *application) logoutUserHandler(w http.ResponseWriter, r *http.Request) {
	// Clear the session cookie
	cookieMode := http.SameSiteLaxMode
	isSecure := false

	if app.config.env == "production" {
		cookieMode = http.SameSiteNoneMode
		isSecure = true
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "session_token",
		Value:    "",
		Path:     "/",
		Expires:  time.Unix(0, 0),
		MaxAge:   -1,
		HttpOnly: true,
		Secure:   isSecure,
		SameSite: cookieMode,
	})

	app.writeJSON(w, http.StatusOK, envelope{"message": "logout successful"}, nil)
}
