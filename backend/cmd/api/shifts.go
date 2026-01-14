package main

import (
	"errors"
	"net/http"

	"github.com/cconner57/adoption-os/backend/internal/data"
	"github.com/cconner57/adoption-os/backend/internal/validator"
)

func (app *application) createShiftHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		VolunteerID int64  `json:"volunteerId"`
		Date        string `json:"date"`
		StartTime   string `json:"startTime"`
		EndTime     string `json:"endTime"`
		Role        string `json:"role"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	shift := &data.Shift{
		VolunteerID: input.VolunteerID,
		Date:        input.Date,
		StartTime:   input.StartTime,
		EndTime:     input.EndTime,
		Role:        input.Role,
		Status:      "scheduled", // Default status
	}

	v := validator.New()
	if input.VolunteerID <= 0 {
		v.AddError("volunteerId", "must be a valid ID")
	}
	if input.Date == "" {
		v.AddError("date", "must be provided")
	}
	// Basic validation, detailed validation can be added later

	if !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	err = app.models.Shifts.Insert(shift)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	// Recalculate stats in bg or sync? Sync is fine for now.
	app.recalculateVolunteerStats(shift.VolunteerID)

	app.JSONResponse(w, http.StatusCreated, envelope{"shift": shift})
}

func (app *application) listVolunteerShiftsHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	shifts, err := app.models.Shifts.GetForVolunteer(id)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	app.JSONResponse(w, http.StatusOK, envelope{"shifts": shifts})
}

func (app *application) updateShiftHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	// Currently we don't have a Get specific shift method, we could add one or just GetForVolunteer and filter,
	// but better to add a Get(id) in model or just do a direct update.
	// For concurrency check we need version.
	// Given the simplicity, let's just do a direct update or we need to fetch it first.
	// Let's rely on the client sending the full object including version if possible,
	// or we fetch it first using a new method Get(id).
	// For now, I'll implement a Get in the model to make this robust.
	// WAIT: I defined Update(shift *Shift) which requires version.
	// So I need to fetch the shift first.

	// Let's add Get(id) to ShiftModel in data/shifts.go quickly?
	// Or I can add it here.
	// Actually, I'll add Get(id) to ShiftModel quickly first in a separate tool call to be clean.
	// But to save time let's check if I can just assume existence.
	// No, Update requires checking version.
	// I'll assume I'll add Get method.

	shift, err := app.models.Shifts.Get(id)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	var input struct {
		Date        *string `json:"date"`
		StartTime   *string `json:"startTime"`
		EndTime     *string `json:"endTime"`
		Role        *string `json:"role"`
		Status      *string `json:"status"`
		Notes       *string `json:"notes"`
		ID          *int64  `json:"id"`          // Ignored
		VolunteerID *int64  `json:"volunteerId"` // Ignored
	}

	err = app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	if input.Date != nil {
		shift.Date = *input.Date
	}
	if input.StartTime != nil {
		shift.StartTime = *input.StartTime
	}
	if input.EndTime != nil {
		shift.EndTime = *input.EndTime
	}
	if input.Role != nil {
		shift.Role = *input.Role
	}
	if input.Status != nil {
		shift.Status = *input.Status
	}
	if input.Notes != nil {
		shift.Notes = *input.Notes
	}

	v := validator.New()
	if shift.Date == "" {
		v.AddError("date", "must be provided")
	}

	if !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	err = app.models.Shifts.Update(shift)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrEditConflict):
			app.editConflictResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	app.recalculateVolunteerStats(shift.VolunteerID)

	app.JSONResponse(w, http.StatusOK, envelope{"shift": shift})
}

func (app *application) deleteShiftHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	// Get shift first to preserve VolunteerID for recalc
	shift, err := app.models.Shifts.Get(id)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	err = app.models.Shifts.Delete(id)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	app.recalculateVolunteerStats(shift.VolunteerID)

	app.JSONResponse(w, http.StatusOK, envelope{"message": "shift deleted successfully"})
}
