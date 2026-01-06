package main

import (
	"net/http"

	"github.com/rs/cors" // Make sure you have this
)

func (app *application) routes() http.Handler { // <--- lowercase application
	mux := http.NewServeMux()

	// Register Routes (Use the lowercase function names we just defined)
	mux.HandleFunc("GET /", app.home)
	mux.HandleFunc("GET /pets", app.getAllPets)
	mux.HandleFunc("GET /pets/spotlight", app.getSpotlightPets)
	mux.HandleFunc("GET /pets/available", app.getAvailablePets)
	mux.HandleFunc("GET /pets/adopted-count", app.getAdoptedPetsCount)

	// Add the new Volunteer Route
	mux.HandleFunc("POST /applications/volunteer", app.submitVolunteerApplication)
	mux.HandleFunc("POST /applications/adoption", app.submitAdoptionApplication)

	// Setup CORS
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173", "http://localhost:4173", "https://idohr.netlify.app"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type"},
		AllowCredentials: true,
	})

	return c.Handler(mux)
}
