package main

import (
	"net/http"

	"github.com/rs/cors"
)

func (app *application) routes() http.Handler {
	mux := http.NewServeMux()

	// Register Routes
	mux.HandleFunc("GET /", app.healthcheckHandler)
	mux.HandleFunc("GET /healthz", app.healthcheckHandler)
	mux.HandleFunc("GET /sitemap.xml", app.sitemapHandler)
	mux.HandleFunc("GET /pets", app.getAllPets)
	mux.HandleFunc("GET /pets/spotlight", app.getSpotlightPets)
	mux.HandleFunc("GET /pets/available", app.getAvailablePets)
	mux.HandleFunc("GET /pets/adopted-count", app.getAdoptedPetsCount)

	mux.HandleFunc("POST /applications/volunteer", app.submitVolunteerApplication)
	mux.HandleFunc("POST /applications/adoption", app.submitAdoptionApplication)
	mux.HandleFunc("POST /applications/surrender", app.submitSurrenderApplication)
	mux.HandleFunc("POST /metrics", app.submitMetric)

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173", "http://localhost:4173", "https://idohr.netlify.app"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type"},
		AllowCredentials: true,
	})

	return app.recoverPanic(app.rateLimit(c.Handler(mux)))
}
