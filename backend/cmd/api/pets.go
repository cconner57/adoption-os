package main

import (
	"fmt"
	"net/http"
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
	app.JSONResponse(w, http.StatusOK, map[string]string{"message": "GetAllPets coming soon"})
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
