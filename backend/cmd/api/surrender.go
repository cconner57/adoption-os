package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/cconner57/adoption-os/backend/internal/data"
	"github.com/cconner57/adoption-os/backend/internal/validator"
)

func (app *application) submitSurrenderApplication(w http.ResponseWriter, r *http.Request) {
	var input data.SurrenderApplication

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	// Honeypot Check
	if input.FaxNumber != "" {
		app.logger.Warn("Bot detected: honeypot populated", "field", "fax_number", "ip", r.RemoteAddr)
		// Fake success
		err = app.JSONResponse(w, http.StatusOK, map[string]string{"message": "Surrender application submitted successfully"})
		if err != nil {
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	v := validator.New()
	data.ValidateSurrenderApplication(v, &input)

	if !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	// Construct Email Body
	var body strings.Builder
	body.WriteString(fmt.Sprintf("New Surrender Application Received\n\n"))
	body.WriteString(fmt.Sprintf("Applicant: %s %s\n", input.FirstName, input.LastName))
	body.WriteString(fmt.Sprintf("Email: %s\n", input.Email))
	body.WriteString(fmt.Sprintf("Phone: %s\n", input.PhoneNumber))
	body.WriteString(fmt.Sprintf("Address: %s, %s, %s %s\n\n", input.StreetAddress, input.City, input.State, input.ZipCode))

	body.WriteString("Animal Details:\n")
	body.WriteString(fmt.Sprintf("Name: %s\n", input.AnimalName))
	body.WriteString(fmt.Sprintf("Age: %s\n", input.AnimalAge))
	body.WriteString(fmt.Sprintf("Sex: %s\n", input.AnimalSex))
	body.WriteString(fmt.Sprintf("Ownership Duration: %s\n", input.AnimalOwnershipDuration))
	body.WriteString(fmt.Sprintf("Location Found: %s\n", input.AnimalLocationFound))
	body.WriteString(fmt.Sprintf("Reason for Surrender: %s\n\n", input.AnimalWhySurrendered))

	body.WriteString("Behaviors & History:\n")
	body.WriteString(fmt.Sprintf("Attacked People: %s (%s)\n", input.AnimalEverAttackedPeople, input.AnimalEverAttackedPeopleExplanation))
	body.WriteString(fmt.Sprintf("Attacked Other Cats: %s (%s)\n", input.AnimalEverAttackedOtherCats, input.AnimalEverAttackedOtherCatsExplanation))
	body.WriteString(fmt.Sprintf("Attacked Other Dogs: %s (%s)\n", input.AnimalEverAttackedOtherDogs, input.AnimalEverAttackedOtherDogsExplanation))
	body.WriteString(fmt.Sprintf("Housetrained: %s\n", input.AnimalHouseTrained))
	body.WriteString(fmt.Sprintf("Vet/Groomer Behavior: %s\n", input.AnimalVetOrGroomerBehavior))

	body.WriteString(fmt.Sprintf("\nFull Request Data:\n%+v\n", input))

	sender := app.config.smtp.sender
	err = app.mailer.Send(sender, "New Surrender Application - "+input.AnimalName, body.String(), nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	app.JSONResponse(w, http.StatusCreated, map[string]string{"message": "Surrender application submitted successfully"})
}
