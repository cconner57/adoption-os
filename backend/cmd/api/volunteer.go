package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/cconner57/adoption-os/backend/internal/data"
	"github.com/cconner57/adoption-os/backend/internal/validator"
)

func (app *application) submitVolunteerApplication(w http.ResponseWriter, r *http.Request) {
	var input data.VolunteerApplication

	// Read JSON from frontend
	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	// Honeypot Check
	if input.FaxNumber != "" {
		app.logger.Warn("Bot detected: honeypot populated", "field", "fax_number", "ip", r.RemoteAddr)
		// Fake success
		err = app.JSONResponse(w, http.StatusOK, map[string]string{"message": "Application received and validated"})
		if err != nil {
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	// Validate
	v := validator.New()
	data.ValidateVolunteerApplication(v, &input)
	if !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	// Log success (Task requirement: console log validation success)
	fmt.Printf("Volunteer Application Validated Successfully: %+v\n", input)

	// Send Email Notification
	// Runs synchronously to ensure we only return success if email sends (or is simulated)
	subject := fmt.Sprintf("New Volunteer Application: %s %s", input.FirstName, input.LastName)

	// Helper to format bool
	yesNo := func(b bool) string {
		if b {
			return "Yes"
		}
		return "No"
	}

	// Helper to format comma separated list
	formatList := func(list []string) string {
		return strings.Join(list, ", ")
	}

	// Helper to format date
	formatDate := func(dateStr string) string {
		if dateStr == "" {
			return ""
		}
		t, err := time.Parse("2006-01-02", dateStr)
		if err != nil {
			return dateStr // Return original if parse fails
		}
		return t.Format("Jan 02, 2006")
	}

	var sb strings.Builder
	sb.WriteString(fmt.Sprintf(`<!DOCTYPE html>
<html>
<head>
<style>
  body { font-family: Arial, sans-serif; color: #333; line-height: 1.6; }
  .container { max-width: 600px; margin: 0 auto; padding: 20px; border: 1px solid #e0e0e0; border-radius: 8px; }
  .header { text-align: center; margin-bottom: 30px; }
  .logo { max-width: 150px; height: auto; }
  h1 { color: #00a5ad; font-size: 24px; text-align: center; }
  h2 { color: #00a5ad; font-size: 18px; border-bottom: 2px solid #00a5ad; padding-bottom: 5px; margin-top: 25px; }
  .field { margin-bottom: 10px; }
  .label { font-weight: bold; color: #555; }
  .footer { margin-top: 30px; font-size: 12px; color: #333; text-align: center; border-top: 1px solid #eee; padding-top: 10px; }
</style>
</head>
<body>
<div class="container">
  <div class="header">
    <img src="cid:logo" alt="IDOHR Logo" class="logo">
  </div>
  <h1>New Volunteer Application</h1>
`))

	sb.WriteString(`<h2>Personal Information</h2>`)
	fmt.Fprintf(&sb, `<div class="field"><span class="label">Name:</span> %s %s</div>`, input.FirstName, input.LastName)
	fmt.Fprintf(&sb, `<div class="field"><span class="label">Address:</span> %s, %s, %s</div>`, input.Address, input.City, input.Zip)
	fmt.Fprintf(&sb, `<div class="field"><span class="label">Phone:</span> %s</div>`, input.PhoneNumber)
	fmt.Fprintf(&sb, `<div class="field"><span class="label">Birthday:</span> %s</div>`, formatDate(input.Birthday))
	fmt.Fprintf(&sb, `<div class="field"><span class="label">Age:</span> %d</div>`, safeInt(input.Age))
	fmt.Fprintf(&sb, `<div class="field"><span class="label">Allergies:</span> %s</div>`, yesNo(input.Allergies))

	sb.WriteString(`<h2>Emergency Contact</h2>`)
	fmt.Fprintf(&sb, `<div class="field"><span class="label">Name:</span> %s</div>`, input.EmergencyContactName)
	fmt.Fprintf(&sb, `<div class="field"><span class="label">Phone:</span> %s</div>`, input.EmergencyContactPhone)

	sb.WriteString(`<h2>Experience & Interests</h2>`)
	sb.WriteString(`<div class="field"><span class="label">Volunteer Experience:</span></div>`)
	fmt.Fprintf(&sb, `<div>%s</div><br>`, input.VolunteerExperience)
	sb.WriteString(`<div class="field"><span class="label">Interest Reason:</span></div>`)
	fmt.Fprintf(&sb, `<div>%s</div><br>`, input.InterestReason)
	fmt.Fprintf(&sb, `<div class="field"><span class="label">Position Preferences:</span> %s</div>`, formatList(input.PositionPreferences))
	fmt.Fprintf(&sb, `<div class="field"><span class="label">Availability:</span> %s</div>`, formatList(input.Availability))

	sb.WriteString(`<h2>Agreement</h2>`)
	fmt.Fprintf(&sb, `<div class="field"><span class="label">Signed By:</span> %s</div>`, input.NameFull)
	fmt.Fprintf(&sb, `<div class="field"><span class="label">Date:</span> %s</div>`, formatDate(input.SignatureDate))

	// Conditionally add Parent/Guardian info
	if input.ParentName != "" {
		fmt.Fprintf(&sb, `<br><div class="field"><span class="label">Parent/Guardian Name:</span> %s</div>`, input.ParentName)
	}
	if input.ParentSignatureDate != "" {
		fmt.Fprintf(&sb, `<div class="field"><span class="label">Parent/Guardian Date:</span> %s</div>`, formatDate(input.ParentSignatureDate))
	}

	fmt.Fprintf(&sb, `
  <div class="footer">
    This application was submitted via the I Dream of Home Rescue Volunteer Form.<br>
    %s
  </div>
</div>
</body>
</html>`, time.Now().Format("Jan 02, 2006"))

	body := sb.String()

	// Save to Database
	appRecord := &data.Application{
		Type:         "volunteer",
		Status:       "pending",
		Data:         []byte("{}"),
		OriginalHTML: &body,
	}

	jsonData, err := json.Marshal(input)
	if err == nil {
		appRecord.Data = jsonData
	}

	err = app.models.Applications.Insert(appRecord)
	if err != nil {
		app.logger.Error("Failed to persist volunteer application", "error", err)
	}

	// Send to the configured sender address (acting as Admin)
	// Or you could read a specific recipient from config.
	recipient := app.config.smtp.sender
	if recipient == "" {
		recipient = "cats@idohr.org" // Fallback
	}

	// Prepare attachments
	attachments := make(map[string][]byte)

	// Note: Logo is now embedded via CID
	// We re-use readLogoBase64 logic but return bytes?
	getCandidateLogo := func() []byte {
		cwd, _ := os.Getwd()
		app.logger.Info("Current working directory", "cwd", cwd)
		candidates := []string{
			"../frontend/public/images/idohr-logo.jpg",
			"../../frontend/public/images/idohr-logo.jpg",
			"./frontend/public/images/idohr-logo.jpg",
			"/Users/conner/Desktop/adoption-os/frontend/public/images/idohr-logo.jpg",
		}
		for _, path := range candidates {
			data, err := os.ReadFile(path)
			if err == nil {
				app.logger.Info("Found logo at", "path", path)
				return data
			}
		}
		return nil
	}

	if logoBytes := getCandidateLogo(); logoBytes != nil {
		attachments["logo.jpg"] = logoBytes
	}

	if input.SignatureData != nil && *input.SignatureData != "" {
		// signatureData is likely "data:image/png;base64,....."
		// Split by comma to get the actual base64 part
		parts := strings.Split(*input.SignatureData, ",")
		if len(parts) == 2 {
			sigBytes, err := base64.StdEncoding.DecodeString(parts[1])
			if err == nil {
				attachments["signature.png"] = sigBytes
			} else {
				app.logger.Error("Failed to decode signature", "error", err)
			}
		}
	}

	// Check if we have credentials; if not, simulate sending
	if app.config.smtp.password == "" || app.config.smtp.username == "" {
		app.logger.Info("Development Mode: Simulating sending email", "recipient", recipient, "subject", subject)
	} else {
		// Attempt to send email
		err := app.mailer.Send(recipient, subject, body, attachments)
		if err != nil {
			app.logger.Error("Failed to send email notification", "error", err)

			// In development, don't block the application if email fails (common with bad/missing creds)
			if app.config.env == "development" {
				app.logger.Warn("Development Mode: Ignoring email usage error, proceeding with success.")
			} else {
				app.serverErrorResponse(w, r, fmt.Errorf("failed to send email notification: %w", err))
				return
			}
		} else {
			app.logger.Info("Email notification sent successfully", "recipient", recipient)
		}
	}

	// TODO: Insert into DB (Reserved for next task)
	// err = app.models.Volunteers.Insert(&input)
	// if err != nil {
	// 	app.serverErrorResponse(w, r, err)
	// 	return
	// }

	// Send back a success response
	err = app.JSONResponse(w, http.StatusOK, map[string]string{"message": "Application received and validated"})
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
