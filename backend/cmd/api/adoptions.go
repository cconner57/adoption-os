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

func (app *application) submitAdoptionApplication(w http.ResponseWriter, r *http.Request) {
	fmt.Println("DEBUG: Hit submitAdoptionApplication handler")
	var input data.AdoptionApplication

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
		err = app.JSONResponse(w, http.StatusOK, map[string]string{"message": "Adoption application received"})
		if err != nil {
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	// Validate
	v := validator.New()
	data.ValidateAdoptionApplication(v, &input)
	if !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	// Log success
	fmt.Printf("Adoption Application Validated Successfully: %+v\n", input)

	// Send Email Notification (Async)
	go func() {
		defer func() {
			if err := recover(); err != nil {
				app.logger.Error("Panic in adoption email goroutine", "error", err)
			}
		}()

		subject := fmt.Sprintf("New Adoption Application: %s %s", input.FirstName, input.LastName)

		// Helpers
		safeStr := func(s *string) string {
			if s == nil {
				return ""
			}
			return *s
		}

		// Helper to read logo as Base64
		readLogoBase64 := func() string {
			cwd, _ := os.Getwd()
			app.logger.Info("Current working directory", "cwd", cwd)

			candidates := []string{
				"../frontend/public/images/idohr-logo.jpg",
				"../../frontend/public/images/idohr-logo.jpg",
				"./frontend/public/images/idohr-logo.jpg",
				"/Users/conner/Desktop/adoption-os/frontend/public/images/idohr-logo.jpg", // Fallback absolute
			}

			var data []byte
			var err error

			for _, path := range candidates {
				data, err = os.ReadFile(path)
				if err == nil {
					app.logger.Info("Found logo at", "path", path)
					return base64.StdEncoding.EncodeToString(data)
				}
			}

			app.logger.Error("Failed to read logo from any candidate path", "error", err)
			return ""
		}

		logoBase64 := readLogoBase64()
		logoSrc := ""
		if logoBase64 != "" {
			logoSrc = fmt.Sprintf("data:image/jpeg;base64,%s", logoBase64)
		}

		var sb strings.Builder
		sb.WriteString(fmt.Sprintf(`<!DOCTYPE html>
<html>
<head>
<style>
  body { font-family: Arial, sans-serif; color: #333; line-height: 1.6; }
  .container { max-width: 800px; margin: 0 auto; padding: 20px; border: 1px solid #e0e0e0; border-radius: 8px; }
  .header { text-align: center; margin-bottom: 30px; }
  .logo { max-width: 150px; height: auto; }
  h1 { color: #00a5ad; font-size: 24px; text-align: center; }
  h2 { color: #00a5ad; font-size: 18px; border-bottom: 2px solid #00a5ad; padding-bottom: 5px; margin-top: 25px; }
  .section { margin-bottom: 20px; }
  .field { margin-bottom: 10px; }
  .label { font-weight: bold; color: #555; }
  .table { width: 100%; border-collapse: collapse; margin-top: 10px; }
  .table th, .table td { border: 1px solid #ddd; padding: 8px; text-align: left; font-size: 14px; }
  .table th { background-color: #f2f2f2; }
  .footer { margin-top: 30px; font-size: 12px; color: #333; text-align: center; border-top: 1px solid #eee; padding-top: 10px; }
</style>
</head>
<body>
<div class="container">
  <div class="header">
    <img src="%s" alt="IDOHR Logo" class="logo">
  </div>
  <h1>New Adoption Application</h1>
`, logoSrc))

		// --- Personal Information ---
		sb.WriteString(`<h2>Personal Information</h2>`)
		fmt.Fprintf(&sb, `<div class="field"><span class="label">Name:</span> %s %s</div>`, input.FirstName, input.LastName)
		if input.SpouseFirstName != nil && *input.SpouseFirstName != "" {
			fmt.Fprintf(&sb, `<div class="field"><span class="label">Spouse/Partner:</span> %s %s</div>`, *input.SpouseFirstName, safeStr(input.SpouseLastName))
		}
		if input.Age != nil {
			fmt.Fprintf(&sb, `<div class="field"><span class="label">Age:</span> %v</div>`, input.Age)
		}
		fmt.Fprintf(&sb, `<div class="field"><span class="label">Address:</span> %s %s, %s, %s %s</div>`, safeStr(input.Address), safeStr(input.AddressLine2), safeStr(input.City), safeStr(input.State), safeStr(input.Zip))
		fmt.Fprintf(&sb, `<div class="field"><span class="label">Phone:</span> %s</div>`, safeStr(input.PhoneNumber))
		if input.CellPhoneNumber != nil && *input.CellPhoneNumber != "" {
			fmt.Fprintf(&sb, `<div class="field"><span class="label">Cell:</span> %s</div>`, *input.CellPhoneNumber)
		}
		fmt.Fprintf(&sb, `<div class="field"><span class="label">Email:</span> %s</div>`, safeStr(input.Email))

		// --- Household Information ---
		sb.WriteString(`<h2>Household & Residence</h2>`)
		if input.HomeOwnership != nil && *input.HomeOwnership == "Rent" {
			fmt.Fprintf(&sb, `<div class="field"><span class="label">Landlord Name:</span> %s</div>`, safeStr(input.LandlordName))
			fmt.Fprintf(&sb, `<div class="field"><span class="label">Landlord Phone:</span> %s</div>`, safeStr(input.LandlordPhoneNumber))
		}
		fmt.Fprintf(&sb, `<div class="field"><span class="label">Length of Residence:</span> %s</div>`, safeStr(input.YearsAtAddress))
		if input.PreviousAddress != nil && *input.PreviousAddress != "" {
			fmt.Fprintf(&sb, `<div class="field"><span class="label">Previous Address:</span> %s</div>`, *input.PreviousAddress)
		}
		fmt.Fprintf(&sb, `<div class="field"><span class="label">Expect to Move:</span> %s</div>`, safeStr(input.ExpectToMove))
		fmt.Fprintf(&sb, `<div class="field"><span class="label">Travel Plans:</span> %s</div>`, safeStr(input.TravelPlan))
		fmt.Fprintf(&sb, `<div class="field"><span class="label">Allergies:</span> %s</div>`, safeStr(input.Allergies))
		if input.PrimaryOwner != nil {
			fmt.Fprintf(&sb, `<div class="field"><span class="label">Primary Owner:</span> %v</div>`, input.PrimaryOwner)
		}
		fmt.Fprintf(&sb, `<div class="field"><span class="label">All Adults Agreed:</span> %s</div>`, safeStr(input.AdultMembersAgreed))

		// Household Members
		if len(input.RoommatesNames) > 0 {
			sb.WriteString(`<div class="field"><span class="label">Roommates/Other Adults:</span></div><ul>`)
			for _, name := range input.RoommatesNames {
				if name != "" {
					fmt.Fprintf(&sb, "<li>%s</li>", name)
				}
			}
			sb.WriteString(`</ul>`)
		}
		if len(input.ChildrenNamesAges) > 0 {
			sb.WriteString(`<div class="field"><span class="label">Children:</span></div><ul>`)
			for _, child := range input.ChildrenNamesAges {
				if child.Name != "" {
					fmt.Fprintf(&sb, "<li>%s (Age: %s)</li>", child.Name, child.Age)
				}
			}
			sb.WriteString(`</ul>`)
		}

		// --- Cat Preferences ---
		sb.WriteString(`<h2>Adoption Preferences</h2>`)
		fmt.Fprintf(&sb, `<div class="field"><span class="label">Reason for Adopting:</span> %s</div>`, safeStr(input.AdoptionReason))
		fmt.Fprintf(&sb, `<div class="field"><span class="label">Why Interested in Specific Cat:</span> %s</div>`, safeStr(input.WhyInterested))
		fmt.Fprintf(&sb, `<div class="field"><span class="label">Desired Breed/Type:</span> %s</div>`, safeStr(input.CatPreferenceBreed))
		fmt.Fprintf(&sb, `<div class="field"><span class="label">Physical Preferences:</span> %s</div>`, safeStr(input.CatPreferencePhysical))
		fmt.Fprintf(&sb, `<div class="field"><span class="label">Personality Preferences:</span> %s</div>`, safeStr(input.CatPreferencePersonality))
		fmt.Fprintf(&sb, `<div class="field"><span class="label">Qualities Not Wanted:</span> %s</div>`, safeStr(input.CatPreferenceNotWant))
		fmt.Fprintf(&sb, `<div class="field"><span class="label">Indoor/Outdoor:</span> %s</div>`, safeStr(input.CatIndoorOutdoor))
		fmt.Fprintf(&sb, `<div class="field"><span class="label">Cat Access Areas:</span> %s</div>`, safeStr(input.CatAccess))
		fmt.Fprintf(&sb, `<div class="field"><span class="label">Home Alone Hours:</span> %s</div>`, safeStr(input.CatHomeAloneHours))
		fmt.Fprintf(&sb, `<div class="field"><span class="label">Discipline Method:</span> %s</div>`, safeStr(input.CatDisciplineType))

		// --- Current Pets ---
		sb.WriteString(`<h2>Current Pets</h2>`)
		fmt.Fprintf(&sb, `<div class="field"><span class="label">Currently Have Pets:</span> %s</div>`, safeStr(input.CurrentlyHavePets))
		if len(input.CurrentPets) > 0 {
			sb.WriteString(`<table class="table"><tr><th>Name</th><th>Species/Breed</th><th>Age</th><th>Spayed/Neutered</th><th>Likes Dogs?</th></tr>`)
			for _, pet := range input.CurrentPets {
				fmt.Fprintf(&sb, `<tr><td>%s</td><td>%s</td><td>%s</td><td>%s</td><td>%s</td></tr>`,
					pet.Name, pet.SpeciesBreedSize, pet.Age, pet.SpayedNeutered, pet.LikesDogs)
			}
			sb.WriteString(`</table>`)
		}

		// --- Past Pets ---
		sb.WriteString(`<h2>Past Pets</h2>`)
		fmt.Fprintf(&sb, `<div class="field"><span class="label">Owned Pets Before:</span> %s</div>`, safeStr(input.OwnPetsBefore))
		if len(input.PastPets) > 0 {
			sb.WriteString(`<table class="table"><tr><th>Name</th><th>Species/Breed</th><th>Age</th><th>Spayed/Neutered</th><th>Last Situation/Reason</th></tr>`)
			for _, pet := range input.PastPets {
				fmt.Fprintf(&sb, `<tr><td>%s</td><td>%s</td><td>%s</td><td>%s</td><td>%s</td></tr>`,
					pet.Name, pet.SpeciesBreedSize, pet.Age, pet.SpayedNeutered, pet.PassedAwayReason)
			}
			sb.WriteString(`</table>`)
		}

		// --- Other ---
		sb.WriteString(`<h2>Additional Information</h2>`)
		fmt.Fprintf(&sb, `<div class="field"><span class="label">Veterinarian:</span> %s</div>`, safeStr(input.AlreadyHaveVeterinarian))
		fmt.Fprintf(&sb, `<div class="field"><span class="label">Declawed/Debarked:</span> %s</div>`, safeStr(input.OwnedDeclawedOrDebarked))
		fmt.Fprintf(&sb, `<div class="field"><span class="label">Moved with Pet:</span> %s</div>`, safeStr(input.MovedWithPet))
		fmt.Fprintf(&sb, `<div class="field"><span class="label">Surrender Plan:</span> %s</div>`, safeStr(input.SurrenderPlan))
		if len(input.SurrenderConditions) > 0 {
			fmt.Fprintf(&sb, `<div class="field"><span class="label">Checked Surrender Conditions so far:</span> %s</div>`, strings.Join(input.SurrenderConditions, ", "))
		}
		fmt.Fprintf(&sb, `<div class="field"><span class="label">Food Type/Brand:</span> %s</div>`, safeStr(input.FoodTypeBrand))
		fmt.Fprintf(&sb, `<div class="field"><span class="label">Afford Vet Care:</span> %s</div>`, safeStr(input.AffordVetCare))
		fmt.Fprintf(&sb, `<div class="field"><span class="label">Afford Emergency:</span> %s</div>`, safeStr(input.AffordEmergencyCost))

		// --- Signatures ---
		sb.WriteString(`<h2>Agreements & Signatures</h2>`)
		fmt.Fprintf(&sb, `<div class="field"><span class="label">Signature 1:</span> %s</div>`, safeStr(input.AgreementSignature1))
		fmt.Fprintf(&sb, `<div class="field"><span class="label">Signature 2:</span> %s</div>`, safeStr(input.AgreementSignature2))
		if input.AgreementSignature3 != nil {
			fmt.Fprintf(&sb, `<div class="field"><span class="label">Signature 3:</span> %s</div>`, safeStr(input.AgreementSignature3))
		}

		fmt.Fprintf(&sb, `
  <div class="footer">
    This application was submitted via the I Dream of Home Rescue Adoption Form.<br>
    %s
  </div>
</div>
</body>
</html>`, time.Now().Format("Jan 02, 2006"))

		body := sb.String()

		// Save to Database
		appRecord := &data.Application{
			Type:         "adoption",
			Status:       "pending",
			Data:         []byte("{}"), // We should marshal input to JSON, but 'input' is struct.
			OriginalHTML: body,
		}
		// Marshal input
		jsonData, err := json.Marshal(input)
		if err == nil {
			appRecord.Data = jsonData
		}

		err = app.models.Applications.Insert(appRecord)
		if err != nil {
			app.logger.Error("Failed to persist adoption application", "error", err)
			// We continue to send email even if DB fails? Or fail?
			// Better to log and try to email, as email is critical path historically.
		}

		// Send email

		// Configure Recipient
		recipient := app.config.smtp.sender
		if recipient == "" {
			recipient = "cats@idohr.org" // Fallback
		}

		// Attachments (Logo + Signature)
		attachments := make(map[string][]byte)

		// Logo is embedded now.

		// Attach Final Signature Image
		if input.SignatureData != nil && *input.SignatureData != "" {
			parts := strings.Split(*input.SignatureData, ",")
			if len(parts) == 2 {
				sigBytes, err := base64.StdEncoding.DecodeString(parts[1])
				if err == nil {
					attachments["signature.png"] = sigBytes
				} else {
					app.logger.Error("Failed to decode adoption signature", "error", err)
				}
			}
		}

		app.logger.Info("Attempting to send adoption email", "recipient", recipient)
		err = app.mailer.Send(recipient, subject, body, attachments)
		if err != nil {
			app.logger.Error("Failed to send adoption email notification", "error", err)
		} else {
			app.logger.Info("Adoption email notification sent successfully", "recipient", recipient)
		}
	}()

	// Send back success
	err = app.JSONResponse(w, http.StatusOK, map[string]string{"message": "Adoption application received and validated"})
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
