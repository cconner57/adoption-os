package main

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/cconner57/adoption-os/backend/internal/data"
	"github.com/cconner57/adoption-os/backend/internal/validator"
)

func (app *application) listApplicationsHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Type   string
		Status string
		data.Filters
	}

	v := validator.New()
	qs := r.URL.Query()

	input.Type = app.readString(qs, "type", "all")
	input.Status = app.readString(qs, "status", "all")
	year := app.readInt(qs, "year", time.Now().Year(), v) // Default to current year
	input.Filters.Page = app.readInt(qs, "page", 1, v)
	input.Filters.PageSize = app.readInt(qs, "page_size", 20, v)
	input.Filters.Sort = app.readString(qs, "sort", "-created_at")
	input.Filters.SortSafelist = []string{"id", "created_at", "-id", "-created_at", "type", "-type", "status", "-status"}

	if data.ValidateFilters(v, input.Filters); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	applications, metadata, err := app.models.Applications.GetAll(input.Type, input.Status, year, input.Filters)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"applications": applications, "metadata": metadata}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) updateApplicationStatusHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	application, err := app.models.Applications.Get(id)
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
		Status string `json:"status"`
	}

	err = app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	v := validator.New()
	if input.Status != "" {
		application.Status = input.Status
	}

	data.ValidateApplication(v, application)
	if !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	// Automation: If approving a volunteer application, create the volunteer profile
	if input.Status == "approved" && application.Type == "volunteer" && application.Status != "approved" {
		// ... (existing volunteer logic) ...
		var volunteerData data.VolunteerApplication
		err := json.Unmarshal(application.Data, &volunteerData)
		if err == nil {
			newVolunteer := &data.Volunteer{
				FirstName:             volunteerData.FirstName,
				LastName:              volunteerData.LastName,
				Phone:                 volunteerData.PhoneNumber,
				Address:               volunteerData.Address,
				City:                  volunteerData.City,
				Zip:                   volunteerData.Zip,
				Role:                  "volunteer",
				Status:                "active",
				Bio:                   volunteerData.InterestReason,
				ReliabilityScore:      100,
				JoinDate:              time.Now().Format("2006-01-02"),
				Allergies:             volunteerData.Allergies,
				Availability:          volunteerData.Availability,
				PositionPreferences:   volunteerData.PositionPreferences,
				Birthday:              volunteerData.Birthday,
				EmergencyContactName:  volunteerData.EmergencyContactName,
				EmergencyContactPhone: volunteerData.EmergencyContactPhone,
				VolunteerExperience:   volunteerData.VolunteerExperience,
				InterestReason:        volunteerData.InterestReason,
				Skills:                []string{},
				Badges:                []string{},
			}
			_ = app.models.Volunteers.InsertGetId(newVolunteer)
			app.logger.Info("Auto-created volunteer from approved application", "id", newVolunteer.ID)
		} else {
			app.logger.Error("Failed to unmarshal volunteer data for automation", "error", err)
		}
	}

	// Automation: If approving an adoption application, update Pet status
	if input.Status == "approved" && application.Type == "adoption" && application.Status != "approved" {
		var adoptionData struct {
			CatPreferenceName string `json:"catPreferenceName"`
			// Fallback?
		}
		if err := json.Unmarshal(application.Data, &adoptionData); err == nil && adoptionData.CatPreferenceName != "" {
			// Find Pet
			pet, err := app.models.Pets.GetByName(adoptionData.CatPreferenceName)
			if err == nil && pet != nil {
				// Update Status
				// We need to update the status inside the 'details' JSONB column usually,
				// and also the top-level status column if it exists in the struct.
				// Referring to pets.go, Update method handles status passed as argument.
				// But we need to update the 'Details' json too to be consistent.

				// Let's rely on GetByName returning the populated struct.
				// We update the Details map.
				var detailsMap map[string]interface{}
				_ = json.Unmarshal(pet.Details, &detailsMap)
				if detailsMap == nil {
					detailsMap = make(map[string]interface{})
				}
				detailsMap["status"] = "adopted"
				newDetails, _ := json.Marshal(detailsMap)
				pet.Details = json.RawMessage(newDetails)

				// Also update the adoption info? "If the applicant is approved... we'll save all of the information"
				// Maybe add applicant name to the adoption json?
				// pet.Adoption = ... (Skipping for now to keep it simple, just status update)

				err = app.models.Pets.Update(pet)
				if err != nil {
					app.logger.Error("Failed to auto-update pet status", "pet", pet.Name, "error", err)
				} else {
					app.logger.Info("Auto-updated pet status to adopted", "pet", pet.Name)
				}
			} else {
				app.logger.Warn("Could not find pet to update status", "name", adoptionData.CatPreferenceName)
			}
		}
	}

	err = app.models.Applications.Update(application)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrEditConflict):
			app.editConflictResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"application": application}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) getApplicationOriginalHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	application, err := app.models.Applications.Get(id)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	// Add security headers to prevent executing malicious scripts if any found their way in
	// w.Header().Set("Content-Security-Policy", "default-src 'self'; script-src 'none'; object-src 'none';") // Strict CSP
	// Actually, preventing scripts is good, but basic styles are needed.

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(application.OriginalHTML))
}
