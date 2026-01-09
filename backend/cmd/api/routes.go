package main

import (
	"net/http"

	"github.com/rs/cors"
)

func (app *application) routes() http.Handler {
	mux := http.NewServeMux()

	// Register Routes
	// Public Routes
	mux.HandleFunc("GET /", app.healthcheckHandler)
	mux.HandleFunc("GET /healthz", app.healthcheckHandler)
	mux.HandleFunc("GET /sitemap.xml", app.sitemapHandler)
	mux.HandleFunc("GET /pets", app.getAllPets)
	mux.HandleFunc("GET /pets/spotlight", app.getSpotlightPets)
	mux.HandleFunc("GET /pets/available", app.getAvailablePets)
	mux.HandleFunc("GET /pets/adopted-count", app.getAdoptedPetsCount)

	// Protected Routes (Applications & Metrics)
	// We create a protected mux or just wrap handlers inline. Inline is easier for mixed usage here.
	mux.Handle("POST /applications/volunteer", app.requireAuthentication(http.HandlerFunc(app.submitVolunteerApplication)))
	mux.Handle("POST /applications/adoption", app.requireAuthentication(http.HandlerFunc(app.submitAdoptionApplication)))
	mux.Handle("POST /applications/surrender", app.requireAuthentication(http.HandlerFunc(app.submitSurrenderApplication)))
	mux.Handle("POST /metrics", app.requireAuthentication(http.HandlerFunc(app.submitMetric)))

	// User Authentication
	mux.HandleFunc("POST /api/users", app.registerUserHandler)
	mux.HandleFunc("POST /api/login", app.loginUserHandler)
	mux.Handle("GET /users/me", app.requireLogin(http.HandlerFunc(app.profileUserHandler)))

	// Static Files (Uploads)
	fileServer := http.FileServer(http.Dir("./uploads"))
	mux.Handle("GET /uploads/", http.StripPrefix("/uploads", app.cacheControl(fileServer)))

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173", "http://localhost:4173", "https://idohr.netlify.app", "http://192.168.12.102:5173"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-API-KEY"},
		AllowCredentials: true,
	})

	return app.removeIdentity(app.securityHeaders(app.enableRequestID(app.recoverPanic(app.rateLimit(app.bodyLimit(c.Handler(mux)))))))
}
