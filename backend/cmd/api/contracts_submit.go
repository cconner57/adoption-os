package main

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/cconner57/adoption-os/backend/internal/data"
)

func (app *application) submitContractHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		ApplicationID int64  `json:"application_id"`
		Signature     string `json:"signature"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	if input.Signature == "" {
		app.badRequestResponse(w, r, errors.New("signature is required"))
		return
	}

	// 1. Fetch Application
	application, err := app.models.Applications.Get(input.ApplicationID)
	if err != nil {
		if errors.Is(err, data.ErrRecordNotFound) {
			app.notFoundResponse(w, r)
		} else {
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	// 2. Fetch Latest Contract for App
	contract, err := app.models.Contracts.GetByApplicationID(input.ApplicationID)
	if err != nil {
		if errors.Is(err, data.ErrRecordNotFound) {
			app.badRequestResponse(w, r, errors.New("no contract found for this application"))
		} else {
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	// 3. Identify Pet & Check Vetting
	petName, err := app.identifyPetFromApp(application)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	pet, err := app.models.Pets.GetByName(petName)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	isFullyVetted := app.isPetFullyVetted(pet)

	// 4. Update Statuses Logic
	newAppStatus := "adopted"
	newPetStatus := "adopted"

	if !isFullyVetted {
		// Foster-to-Adopt logic
		newAppStatus = "adoption_pending"
		newPetStatus = "adoption-pending" // Ensure hidden from available
		// Note: "adoption-pending" is already in PetStatuses
	}

	// 5. Update Contract Signature
	err = app.models.Contracts.UpdateSignature(contract.Token, input.Signature)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	// 6. Update Application Status
	application.Status = newAppStatus
	err = app.models.Applications.Update(application)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	// 7. Update Pet Status
	// We need to parse pet details to update status in JSON
	// Actually, PetModel.Update handles updating status if passed in struct field or map?
	// PetModel.Update logic:
	// "if val, ok := detailsMap["status"].(string); ok { status = s }"
	// So we need to update the status in the details map and call Update.
	// Wait, PetModel.Update takes *Pet.
	// But Pet struct has `Details json.RawMessage`.
	// Let's modify the map.

	var detailsMap map[string]interface{}
	if len(pet.Details) > 0 {
		if err := json.Unmarshal(pet.Details, &detailsMap); err != nil {
			detailsMap = make(map[string]interface{})
		}
	} else {
		detailsMap = make(map[string]interface{})
	}
	detailsMap["status"] = newPetStatus
	newDetails, _ := json.Marshal(detailsMap)
	pet.Details = json.RawMessage(newDetails)

	// Also update adoption date if adopted?
	if newPetStatus == "adopted" {
		var adoptionMap map[string]interface{}
		if len(pet.Adoption) > 0 {
			if err := json.Unmarshal(pet.Adoption, &adoptionMap); err != nil {
				adoptionMap = make(map[string]interface{})
			}
		} else {
			adoptionMap = make(map[string]interface{})
		}
		if _, ok := adoptionMap["date"]; !ok {
			adoptionMap["date"] = time.Now().Format("1/2/2006")
		}
		newAdoption, _ := json.Marshal(adoptionMap)
		pet.Adoption = json.RawMessage(newAdoption)
	}

	err = app.models.Pets.Update(pet)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	app.writeJSON(w, http.StatusOK, envelope{"message": "Contract submitted successfully", "status": newAppStatus}, nil)
}
