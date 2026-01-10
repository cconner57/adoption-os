package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/cconner57/adoption-os/backend/internal/data"
)

func (app *application) getSpotlightPets(w http.ResponseWriter, r *http.Request) {
	pet, err := app.models.Pets.GetSpotlight()
	if err != nil {
		fmt.Println("Handler GetSpotlight Error:", err)
		app.serverErrorResponse(w, r, err)
		return
	}

	app.JSONResponse(w, http.StatusOK, pet)
}

func (app *application) getAllPets(w http.ResponseWriter, r *http.Request) {
	status := r.URL.Query().Get("status")
	search := r.URL.Query().Get("search")

	// If no status provided, default to 'available' is handled by frontend?
	// Or should backend default?
	// User said "backend route where it will default to return all available".
	// So if status is empty, imply 'available'.
	if status == "" {
		status = "available"
	}

	pets, err := app.models.Pets.GetAll(status, search)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	app.JSONResponse(w, http.StatusOK, pets)
}

func (app *application) getAvailablePets(w http.ResponseWriter, r *http.Request) {
	app.JSONResponse(w, http.StatusOK, map[string]string{"message": "GetAvailablePets coming soon"})
}

func (app *application) getAdoptedPetsCount(w http.ResponseWriter, r *http.Request) {
	// Parse year from query param, default to current year if missing
	yearStr := r.URL.Query().Get("year")
	year := 2026 // Default
	if yearStr != "" {
		fmt.Sscanf(yearStr, "%d", &year)
	}

	count, err := app.models.Pets.GetAdoptedCount(year)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	app.JSONResponse(w, http.StatusOK, map[string]int{"count": count})
}

func (app *application) updatePet(w http.ResponseWriter, r *http.Request) {
	// 1. Get ID from URL path
	// Using r.URL.Path? We don't have r.PathValue in this Go version maybe?
	// But we use stdlib mux.
	// Actually we are using a custom router setup? No, standard ServeMux.
	// If Go < 1.22, we might not have path params easily.
	// But since we are calling it /api/pets/{id}, we need to extract it request URL?
	// Or did we use a router library?
	// routes.go: `mux.Handle("PUT /api/pets/{id}", ...)` -> This is Go 1.22+ syntax.
	// So `r.PathValue("id")` should work.

	id := r.PathValue("id")
	if id == "" {
		app.badRequestResponse(w, r, fmt.Errorf("missing pet id"))
		return
	}

	// 2. Parse Body
	var input struct {
		Name            string          `json:"name"`
		Sex             string          `json:"sex"`
		Physical        json.RawMessage `json:"physical"`
		Behavior        json.RawMessage `json:"behavior"`
		Medical         json.RawMessage `json:"medical"`
		Descriptions    json.RawMessage `json:"descriptions"`
		Details         json.RawMessage `json:"details"`
		Adoption        json.RawMessage `json:"adoption"`
		Foster          json.RawMessage `json:"foster"`
		Returned        json.RawMessage `json:"returned"`
		Sponsored       json.RawMessage `json:"sponsored"`
		Photos          json.RawMessage `json:"photos"`
		ProfileSettings json.RawMessage `json:"profileSettings"`
		LitterName      *string         `json:"litterName"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	// 3. Construct Pet Model
	pet := &data.Pet{
		ID:              id,
		Name:            input.Name,
		Sex:             input.Sex,
		Physical:        input.Physical,
		Behavior:        input.Behavior,
		Medical:         input.Medical,
		Descriptions:    input.Descriptions,
		Details:         input.Details,
		Adoption:        input.Adoption,
		Foster:          input.Foster,
		Returned:        input.Returned,
		Sponsored:       input.Sponsored,
		Photos:          input.Photos,
		ProfileSettings: input.ProfileSettings,
		LitterName:      input.LitterName,
	}

	if input.LitterName != nil {
		fmt.Printf("DEBUG: Updating pet %s with LitterName: %s\n", id, *input.LitterName)
	} else {
		fmt.Printf("DEBUG: Updating pet %s with LitterName: nil\n", id)
	}

	// 4. Update via Model
	err = app.models.Pets.Update(pet)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	// 5. Return success
	app.JSONResponse(w, http.StatusOK, map[string]string{"status": "success", "message": "Pet updated successfully"})
}
