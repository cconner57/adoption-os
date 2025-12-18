package main

import (
	"github.com/rs/cors"
	"net/http"
)

func (app *Application) routes() http.Handler {
	mux := http.NewServeMux()

	// Register Routes
	mux.HandleFunc("GET /", app.Home)
	mux.HandleFunc("GET /pets", app.GetAllPets)
	mux.HandleFunc("GET /pets/spotlight", app.GetSpotlightPets)
	mux.HandleFunc("GET /pets/available", app.GetAvailablePets)

	// Setup CORS
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type"},
		AllowCredentials: true,
	})

	// Wrap the mux with CORS and return it
	return c.Handler(mux)
}
