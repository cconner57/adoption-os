package main

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/cconner57/adoption-os/backend/internal/data"
	"github.com/cconner57/adoption-os/backend/internal/validator"
)

// Shared helper to get logo bytes with fallback
func (app *application) getLogoBytes() []byte {
	candidates := []string{
		"../frontend/public/images/idohr-logo.jpg",
		"../../frontend/public/images/idohr-logo.jpg",
		"./frontend/public/images/idohr-logo.jpg",
		"/Users/conner/Desktop/adoption-os/frontend/public/images/idohr-logo.jpg",
	}
	for _, path := range candidates {
		data, err := os.ReadFile(path)
		if err == nil {
			return data
		}
	}
	return nil
}

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

	// Automation: If approving a volunteer application, create the volunteer profile and send welcome email
	if input.Status == "approved" && application.Type == "volunteer" && application.Status != "approved" {
		// Send Welcome Email
		go app.sendVolunteerWelcomeEmail(application)

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

	// Automation: If approving an adoption application, update Pet status and send email
	if input.Status == "approved" && application.Type == "adoption" && application.Status != "approved" {
		// 1. Send "Congratulations" Email (Async)
		go app.sendAdoptionApprovedEmail(application)

		var adoptionData struct {
			CatPreferenceName string `json:"catPreferenceName"`
		}
		if err := json.Unmarshal(application.Data, &adoptionData); err == nil && adoptionData.CatPreferenceName != "" {
			// Find Pet
			pet, err := app.models.Pets.GetByName(adoptionData.CatPreferenceName)
			if err == nil && pet != nil {
				// Update Status
				var detailsMap map[string]interface{}
				_ = json.Unmarshal(pet.Details, &detailsMap)
				if detailsMap == nil {
					detailsMap = make(map[string]interface{})
				}
				detailsMap["status"] = "adopted"
				newDetails, _ := json.Marshal(detailsMap)
				pet.Details = json.RawMessage(newDetails)

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
	if application.OriginalHTML != nil {
		w.Write([]byte(*application.OriginalHTML))
	} else {
		w.Write([]byte("<!-- No original content available -->"))
	}
}

func (app *application) resendApplicationEmailHandler(w http.ResponseWriter, r *http.Request) {
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

	if application.OriginalHTML == nil {
		app.badRequestResponse(w, r, errors.New("cannot resend email: original content has been deleted"))
		return
	}

	// Determine Subject and Attachments based on Type
	var subject string
	attachments := make(map[string][]byte)

	// 1. Attach Logo (Standard for all)
	if logoBytes := app.getLogoBytes(); logoBytes != nil {
		attachments["logo.jpg"] = logoBytes
	}

	// 2. Type Specific Logic (Subject + Signature)
	switch application.Type {
	case "adoption":
		var d data.AdoptionApplication
		if err := json.Unmarshal(application.Data, &d); err == nil {
			// Helper for safe strings
			safeStr := func(s *string) string {
				if s == nil {
					return ""
				}
				return *s
			}
			subject = fmt.Sprintf("New Adoption Application for %s: %s %s", safeStr(d.PetName), d.FirstName, d.LastName)

			// Signature
			if d.SignatureData != nil && *d.SignatureData != "" {
				parts := strings.Split(*d.SignatureData, ",")
				if len(parts) == 2 {
					if sigBytes, err := base64.StdEncoding.DecodeString(parts[1]); err == nil {
						attachments["signature.png"] = sigBytes
					}
				}
			}
		} else {
			subject = "New Adoption Application (Resent)"
		}

	case "volunteer":
		var d data.VolunteerApplication
		if err := json.Unmarshal(application.Data, &d); err == nil {
			subject = fmt.Sprintf("New Volunteer Application: %s %s", d.FirstName, d.LastName)
			// Signature
			if d.SignatureData != nil && *d.SignatureData != "" {
				parts := strings.Split(*d.SignatureData, ",")
				if len(parts) == 2 {
					if sigBytes, err := base64.StdEncoding.DecodeString(parts[1]); err == nil {
						attachments["signature.png"] = sigBytes
					}
				}
			}
		} else {
			subject = "New Volunteer Application (Resent)"
		}

	case "surrender":
		var d data.SurrenderApplication
		if err := json.Unmarshal(application.Data, &d); err == nil {
			subject = fmt.Sprintf("New Surrender Application - %s", d.AnimalName)
			// No signature image for surrender usually? Checking surrender.go... it doesn't seem to attach one.
		} else {
			subject = "New Surrender Application (Resent)"
		}

	default:
		subject = fmt.Sprintf("Application Details (Resent): %s", application.Type)
	}

	// Send Email
	recipient := app.config.smtp.sender
	if recipient == "" {
		recipient = "cats@idohr.org"
	}

	app.logger.Info("Resending application email", "id", id, "recipient", recipient)

	err = app.mailer.Send(recipient, subject, *application.OriginalHTML, attachments)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	app.JSONResponse(w, http.StatusOK, map[string]string{"message": "Email resent successfully"})
}

func (app *application) sendAdoptionApprovedEmail(application *data.Application) {
	// Parse Data
	var d data.AdoptionApplication
	if err := json.Unmarshal(application.Data, &d); err != nil {
		app.logger.Error("Failed to unmarshal adoption data for approved email", "error", err)
		return
	}

	// Helper for safe strings
	safeStr := func(s *string) string {
		if s == nil {
			return ""
		}
		return *s
	}

	attachments := make(map[string][]byte)
	if logoBytes := app.getLogoBytes(); logoBytes != nil {
		attachments["logo.jpg"] = logoBytes
	}

	subject := fmt.Sprintf("Congratulations! Adoption Approved for %s", safeStr(d.PetName))
	recipient := safeStr(d.Email)

	// Safety Check
	if recipient == "" {
		app.logger.Warn("Cannot send adoption approved email: missing recipient email", "appId", application.ID)
		return
	}

	body := fmt.Sprintf(`<!DOCTYPE html>
<html>
<head>
<style>
  body { font-family: Arial, sans-serif; line-height: 1.6; color: #333; }
  .container { max-width: 600px; margin: 0 auto; padding: 20px; border: 1px solid #e0e0e0; border-radius: 8px; }
  .header { text-align: center; margin-bottom: 30px; }
  .logo { max-width: 150px; height: auto; margin-bottom: 20px; }
  h1 { color: #00a5ad; }
  .content { font-size: 16px; }
  .step-box { background-color: #f0f9fa; border-left: 5px solid #00a5ad; padding: 15px; margin: 20px 0; }
  .step-title { font-weight: bold; color: #005f63; margin-bottom: 10px; display: block; }
  .contact-info { font-weight: bold; }
</style>
</head>
<body>
<div class="container">
  <div class="header">
    <img src="cid:logo.jpg" alt="IDOHR Logo" class="logo">
    <h1>Application Approved!</h1>
  </div>
  
  <div class="content">
    <p>Dear %s,</p>
    <p>We are thrilled to let you know that your adoption application for <strong>%s</strong> has been approved!</p>
    
    <p>We believe you will provide a wonderful home. There is just one next step to finalize the adoption process.</p>

    <div class="step-box">
      <span class="step-title">STEP 2 – SUBMIT “HOME TOUR” VIDEO (1 to 1.5 Minutes)</span>
      <p>Please submit a short video of the outside and inside of your home.</p>
      <p>Send an email to Lori (IDOHR Director) at: <a href="mailto:cats@idohr.org">cats@idohr.org</a></p>
      <p>OR</p>
      <p>Send a text message to Lori: <span class="contact-info">(909-261-4185)</span></p>
    </div>

    <p>Once we receive and review your video, we will send over the final adoption agreement.</p>
    <p>Best Regards,<br>I Dream of Home Rescue Team</p>
  </div>
</div>
</body>
</html>`, d.FirstName, safeStr(d.PetName))

	app.logger.Info("Sending adoption approved email", "recipient", recipient)
	err := app.mailer.Send(recipient, subject, body, attachments)
	if err != nil {
		app.logger.Error("Failed to send adoption approved email", "error", err)
	}
}

func (app *application) sendVolunteerWelcomeEmail(application *data.Application) {
	// Parse Data
	var d data.VolunteerApplication
	if err := json.Unmarshal(application.Data, &d); err != nil {
		app.logger.Error("Failed to unmarshal volunteer data for welcome email", "error", err)
		return
	}

	attachments := make(map[string][]byte)
	if logoBytes := app.getLogoBytes(); logoBytes != nil {
		attachments["logo.jpg"] = logoBytes
	}

	subject := fmt.Sprintf("Welcome to the IDOHR Team, %s!", d.FirstName)
	// Fallback to config sender if volunteer email is somehow empty/spam? No, we trust form.
	recipient := d.Email

	// Safety Check
	if recipient == "" {
		app.logger.Warn("Cannot send volunteer welcome email: missing recipient email", "appId", application.ID)
		return
	}

	body := fmt.Sprintf(`<!DOCTYPE html>
<html>
<head>
<style>
  body { font-family: Arial, sans-serif; line-height: 1.6; color: #333; }
  .container { max-width: 600px; margin: 0 auto; padding: 20px; border: 1px solid #e0e0e0; border-radius: 8px; }
  .header { text-align: center; margin-bottom: 30px; }
  .logo { max-width: 150px; height: auto; margin-bottom: 20px; }
  h1 { color: #00a5ad; }
  .content { font-size: 16px; }
</style>
</head>
<body>
<div class="container">
  <div class="header">
    <img src="cid:logo.jpg" alt="IDOHR Logo" class="logo">
    <h1>Welcome to the Team!</h1>
  </div>
  
  <div class="content">
    <p>Dear %s,</p>
    <p>Congratulations! Your volunteer application has been approved.</p>
    
    <p>We are so excited to have you join I Dream of Home Rescue. Your support means the world to our cats and kittens.</p>

    <p>A coordinator will be reaching out to you shortly with details about orientation and your first shift.</p>

    <p>Thank you for making a difference!</p>

    <p>Warmly,<br>I Dream of Home Rescue Team</p>
  </div>
</div>
</body>
</html>`, d.FirstName)

	app.logger.Info("Sending volunteer welcome email", "recipient", recipient)
	err := app.mailer.Send(recipient, subject, body, attachments)
	if err != nil {
		app.logger.Error("Failed to send volunteer welcome email", "error", err)
	}
}
