package main

import (
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/cconner57/adoption-os/backend/internal/data"
)

func (app *application) generateContractHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		ApplicationID int64 `json:"application_id"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
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

	// 2. Identify Pet
	petName, err := app.identifyPetFromApp(application)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	// 3. Fetch Pet
	pet, err := app.models.Pets.GetByName(petName)
	if err != nil {
		if errors.Is(err, data.ErrRecordNotFound) {
			app.badRequestResponse(w, r, fmt.Errorf("pet '%s' not found in database", petName))
		} else {
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	// 4. Check Medical Status
	isFullyVetted := app.isPetFullyVetted(pet)
	contractType := "FOSTER_TO_ADOPT"
	if isFullyVetted {
		contractType = "STANDARD"
	}

	// 5. Generate Token
	tokenBytes := make([]byte, 16)
	_, _ = rand.Read(tokenBytes)
	token := hex.EncodeToString(tokenBytes)

	// 6. Save Contract
	contract := &data.Contract{
		Token:         token,
		ApplicationID: application.ID,
		Type:          contractType,
		ExpiresAt:     time.Now().Add(24 * time.Hour * 7),
	}

	err = app.models.Contracts.Insert(contract)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	contractURL := fmt.Sprintf("https://adoption-os.com/contract/%s", token)

	app.writeJSON(w, http.StatusCreated, envelope{
		"contract_url": contractURL,
		"type":         contractType,
		"token":        token,
	}, nil)
}

func (app *application) identifyPetFromApp(application *data.Application) (string, error) {
	var appData struct {
		PetName           string `json:"petName"`
		CatPreferenceName string `json:"catPreferenceName"`
		AnimalName        string `json:"animalName"`
	}
	if err := json.Unmarshal(application.Data, &appData); err != nil {
		return "", errors.New("failed to parse application data")
	}

	petName := appData.PetName
	if petName == "" {
		petName = appData.CatPreferenceName
	}
	if petName == "" {
		petName = appData.AnimalName
	}

	if petName == "" {
		return "", errors.New("could not identify pet name from application")
	}

	if idx := strings.Index(petName, "("); idx != -1 {
		petName = strings.TrimSpace(petName[:idx])
	}
	return petName, nil
}

func (app *application) isPetFullyVetted(pet *data.Pet) bool {
	type Medical struct {
		SpayedOrNeutered     bool   `json:"spayedOrNeutered"`
		SpayedOrNeuteredDate string `json:"spayedOrNeuteredDate"`
		VaccinationsUpToDate bool   `json:"vaccinationsUpToDate"`
		Microchip            struct {
			Microchipped bool `json:"microchipped"`
		} `json:"microchip"`
	}

	var medical Medical
	if len(pet.Medical) > 0 {
		if err := json.Unmarshal(pet.Medical, &medical); err != nil {
			app.logger.Error("failed to unmarshal pet medical", "error", err)
			return false
		}
	}

	return medical.Microchip.Microchipped &&
		medical.VaccinationsUpToDate &&
		medical.SpayedOrNeutered
}
