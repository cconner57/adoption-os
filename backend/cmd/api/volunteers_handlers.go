package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/cconner57/adoption-os/backend/internal/data"
	"github.com/cconner57/adoption-os/backend/internal/validator"
)

func (app *application) createVolunteerHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		FirstName             string   `json:"firstName"`
		LastName              string   `json:"lastName"`
		Email                 string   `json:"email"`
		Phone                 string   `json:"phone"`
		Role                  string   `json:"role"`
		Status                string   `json:"status"`
		Bio                   string   `json:"bio"`
		Skills                []string `json:"skills"`
		PositionPreferences   []string `json:"positionPreferences"`
		Availability          []string `json:"availability"`
		Allergies             bool     `json:"allergies"`
		JoinDate              string   `json:"joinDate"`
		Birthday              string   `json:"birthday"`
		EmergencyContactName  string   `json:"emergencyContactName"`
		EmergencyContactPhone string   `json:"emergencyContactPhone"`
		InterestReason        string   `json:"interestReason"`
		VolunteerExperience   string   `json:"volunteerExperience"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	// Validate (Basic)
	v := validator.New()
	if input.FirstName == "" {
		v.AddError("firstName", "must be provided")
	}
	if input.LastName == "" {
		v.AddError("lastName", "must be provided")
	}
	if input.Email == "" {
		v.AddError("email", "must be provided")
	}

	if !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	// Initialize new volunteer
	volunteer := &data.Volunteer{
		FirstName:             input.FirstName,
		LastName:              input.LastName,
		Email:                 input.Email,
		Phone:                 input.Phone,
		Role:                  input.Role,
		Status:                input.Status,
		Bio:                   input.Bio,
		Skills:                input.Skills,
		PositionPreferences:   input.PositionPreferences,
		Availability:          input.Availability,
		Allergies:             input.Allergies,
		ReliabilityScore:      100, // Default
		Badges:                []string{"New Recruit"},
		Birthday:              input.Birthday,
		EmergencyContactName:  input.EmergencyContactName,
		EmergencyContactPhone: input.EmergencyContactPhone,
		InterestReason:        input.InterestReason,
		VolunteerExperience:   input.VolunteerExperience,
	}

	// Default Join Date if logic omitted
	if input.JoinDate != "" {
		volunteer.JoinDate = input.JoinDate
	} else {
		volunteer.JoinDate = time.Now().Format("2006-01-02")
	}

	err = app.models.Volunteers.InsertGetId(volunteer)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	headers := make(http.Header)
	headers.Set("Location", fmt.Sprintf("/v1/volunteers/%d", volunteer.ID))

	err = app.JSONResponse(w, http.StatusCreated, map[string]any{"volunteer": volunteer})
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) listVolunteersHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Name   string
		Role   string
		Status string
		data.Filters
	}

	v := validator.New()
	qs := r.URL.Query()

	input.Name = app.readString(qs, "name", "")
	input.Role = app.readString(qs, "role", "")
	input.Status = app.readString(qs, "status", "")

	input.Filters.Page = app.readInt(qs, "page", 1, v)
	input.Filters.PageSize = app.readInt(qs, "page_size", 20, v)
	input.Filters.Sort = app.readString(qs, "sort", "id")
	input.Filters.SortSafelist = []string{"id", "first_name", "last_name", "created_at"}

	if data.ValidateFilters(v, input.Filters); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	// Call Model (simplified GetAll for now)
	volunteers, metadata, err := app.models.Volunteers.GetAll(input.Name, "", "", "", input.Filters)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	err = app.JSONResponse(w, http.StatusOK, map[string]any{
		"volunteers": volunteers,
		"metadata":   metadata,
	})
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) getVolunteerHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	volunteer, err := app.models.Volunteers.Get(id)
	if err != nil {
		switch {
		case err == data.ErrRecordNotFound:
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	err = app.JSONResponse(w, http.StatusOK, map[string]any{"volunteer": volunteer})
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) updateVolunteerHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	volunteer, err := app.models.Volunteers.Get(id)
	if err != nil {
		switch {
		case err == data.ErrRecordNotFound:
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	var input struct {
		ID                    *int64   `json:"id"`        // Ignored
		CreatedAt             *string  `json:"createdAt"` // Ignored
		UpdatedAt             *string  `json:"updatedAt"` // Ignored
		Version               *int     `json:"version"`   // Ignored
		FirstName             *string  `json:"firstName"`
		LastName              *string  `json:"lastName"`
		Email                 *string  `json:"email"`
		Phone                 *string  `json:"phone"`
		Address               *string  `json:"address"`
		City                  *string  `json:"city"`
		Zip                   *string  `json:"zip"`
		Role                  *string  `json:"role"`
		Status                *string  `json:"status"`
		Bio                   *string  `json:"bio"`
		PhotoURL              *string  `json:"photoUrl"`
		ReliabilityScore      *int     `json:"reliabilityScore"`
		TotalHours            *int     `json:"totalHours"`
		Streak                *int     `json:"streak"`
		JoinDate              *string  `json:"joinDate"`
		Birthday              *string  `json:"birthday"`
		EmergencyContactName  *string  `json:"emergencyContactName"`
		EmergencyContactPhone *string  `json:"emergencyContactPhone"`
		InterestReason        *string  `json:"interestReason"`
		VolunteerExperience   *string  `json:"volunteerExperience"`
		Allergies             *bool    `json:"allergies"`
		Skills                []string `json:"skills"`
		PositionPreferences   []string `json:"positionPreferences"`
		Availability          []string `json:"availability"`
		Badges                []string `json:"badges"`
	}

	err = app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	if input.FirstName != nil {
		volunteer.FirstName = *input.FirstName
	}
	if input.LastName != nil {
		volunteer.LastName = *input.LastName
	}
	if input.Email != nil {
		volunteer.Email = *input.Email
	}
	if input.Phone != nil {
		volunteer.Phone = *input.Phone
	}
	if input.Address != nil {
		volunteer.Address = *input.Address
	}
	if input.City != nil {
		volunteer.City = *input.City
	}
	if input.Zip != nil {
		volunteer.Zip = *input.Zip
	}
	if input.Role != nil {
		volunteer.Role = *input.Role
	}
	if input.Status != nil {
		volunteer.Status = *input.Status
	}
	if input.Bio != nil {
		volunteer.Bio = *input.Bio
	}
	if input.PhotoURL != nil {
		volunteer.PhotoURL = *input.PhotoURL
	}
	if input.ReliabilityScore != nil {
		volunteer.ReliabilityScore = *input.ReliabilityScore
	}
	if input.TotalHours != nil {
		volunteer.TotalHours = *input.TotalHours
	}
	if input.Streak != nil {
		volunteer.Streak = *input.Streak
	}
	if input.JoinDate != nil {
		volunteer.JoinDate = *input.JoinDate
	}
	if input.Birthday != nil {
		volunteer.Birthday = *input.Birthday
	}
	if input.EmergencyContactName != nil {
		volunteer.EmergencyContactName = *input.EmergencyContactName
	}
	if input.EmergencyContactPhone != nil {
		volunteer.EmergencyContactPhone = *input.EmergencyContactPhone
	}
	if input.InterestReason != nil {
		volunteer.InterestReason = *input.InterestReason
	}
	if input.VolunteerExperience != nil {
		volunteer.VolunteerExperience = *input.VolunteerExperience
	}
	if input.Allergies != nil {
		volunteer.Allergies = *input.Allergies
	}
	if input.Skills != nil {
		volunteer.Skills = input.Skills
	}
	if input.PositionPreferences != nil {
		volunteer.PositionPreferences = input.PositionPreferences
	}
	if input.Availability != nil {
		volunteer.Availability = input.Availability
	}
	if input.Badges != nil {
		volunteer.Badges = input.Badges
	}

	v := validator.New()
	if !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	err = app.models.Volunteers.Update(volunteer)
	if err != nil {
		switch {
		case err == data.ErrEditConflict:
			app.editConflictResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	err = app.JSONResponse(w, http.StatusOK, map[string]any{"volunteer": volunteer})
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
