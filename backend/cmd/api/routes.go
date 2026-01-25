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
	// Pet Management
	mux.Handle("PUT /pets/{id}", app.requireLogin(http.HandlerFunc(app.updatePet)))
	mux.Handle("POST /pets", app.requireLogin(http.HandlerFunc(app.createPet)))
	mux.Handle("POST /applications/volunteer", http.HandlerFunc(app.submitVolunteerApplication))
	mux.Handle("POST /applications/adoption", http.HandlerFunc(app.submitAdoptionApplication))
	mux.Handle("POST /applications/surrender", http.HandlerFunc(app.submitSurrenderApplication))
	mux.Handle("POST /metrics", app.requireAuthentication(http.HandlerFunc(app.submitMetric)))

	// Application Management
	mux.Handle("GET /v1/applications", app.requireLogin(http.HandlerFunc(app.listApplicationsHandler)))
	mux.Handle("PUT /v1/applications/{id}", app.requireLogin(http.HandlerFunc(app.updateApplicationStatusHandler)))
	mux.Handle("GET /v1/applications/{id}/original", app.requireLogin(http.HandlerFunc(app.getApplicationOriginalHandler)))
	mux.Handle("POST /v1/applications/{id}/resend-email", app.requireLogin(http.HandlerFunc(app.resendApplicationEmailHandler)))

	// Volunteer Management
	mux.Handle("POST /v1/volunteers", app.requireLogin(http.HandlerFunc(app.createVolunteerHandler)))
	mux.Handle("GET /v1/volunteers", app.requireLogin(http.HandlerFunc(app.listVolunteersHandler)))
	mux.Handle("GET /v1/volunteers/{id}", app.requireLogin(http.HandlerFunc(app.getVolunteerHandler)))
	mux.Handle("PUT /v1/volunteers/{id}", app.requireLogin(http.HandlerFunc(app.updateVolunteerHandler)))

	// Shift Management
	mux.Handle("POST /v1/shifts", app.requireLogin(http.HandlerFunc(app.createShiftHandler)))
	mux.Handle("GET /v1/shifts", app.requireLogin(http.HandlerFunc(app.listShiftsHandler))) // Added
	mux.Handle("GET /v1/volunteers/{id}/shifts", app.requireLogin(http.HandlerFunc(app.listVolunteerShiftsHandler)))
	mux.Handle("PUT /v1/shifts/{id}", app.requireLogin(http.HandlerFunc(app.updateShiftHandler)))
	mux.Handle("DELETE /v1/shifts/{id}", app.requireLogin(http.HandlerFunc(app.deleteShiftHandler)))
	mux.Handle("GET /v1/shifts/meta/roles", app.requireLogin(http.HandlerFunc(app.getShiftRoleStatsHandler)))

	// Marketing Management
	mux.Handle("GET /v1/marketing/campaigns", app.requireLogin(http.HandlerFunc(app.listCampaignsHandler)))
	mux.Handle("GET /v1/marketing/campaigns/{id}", app.requireLogin(http.HandlerFunc(app.getCampaignHandler)))
	mux.Handle("POST /v1/marketing/campaigns", app.requireLogin(http.HandlerFunc(app.createCampaignHandler)))
	mux.Handle("PUT /v1/marketing/campaigns/{id}", app.requireLogin(http.HandlerFunc(app.updateCampaignHandler)))

	// User Authentication
	mux.HandleFunc("POST /api/users", app.registerUserHandler)
	mux.HandleFunc("POST /api/login", app.loginUserHandler)
	mux.HandleFunc("POST /api/users/logout", app.logoutUserHandler)
	mux.Handle("GET /api/users/me", app.requireLogin(http.HandlerFunc(app.profileUserHandler)))
	mux.Handle("PUT /api/users", app.requireLogin(http.HandlerFunc(app.updateUserHandler)))

	// Static Files (Uploads)
	// fileServer := http.FileServer(http.Dir("./uploads"))
	// mux.Handle("GET /uploads/", http.StripPrefix("/uploads", app.cacheControl(fileServer)))

	// Serve Assets (Images)
	// We map /pet-photos/ -> app.config.assetsDir
	assetsServer := http.FileServer(http.Dir(app.config.assetsDir))
	mux.Handle("GET /pet-photos/", http.StripPrefix("/pet-photos/", app.cacheControl(assetsServer)))

	// Upload Route
	// Upload Route
	mux.Handle("POST /pets/{id}/photos", app.requireLogin(http.HandlerFunc(app.uploadPetPhotoHandler)))

	// Notifications
	mux.HandleFunc("POST /v1/notifications/subscribe", app.subscribeHandler)
	mux.HandleFunc("POST /v1/notifications/test", app.testNotificationHandler)                                   // New Test Route
	mux.Handle("POST /v1/notifications/broadcast", app.requireLogin(http.HandlerFunc(app.sendBroadcastHandler))) // Admin only

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173", "http://127.0.0.1:5173", "https://idohr.app", "https://www.idohr.app"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-API-KEY"},
		AllowCredentials: true,
	})

	return app.removeIdentity(app.securityHeaders(app.enableRequestID(app.recoverPanic(app.rateLimit(app.bodyLimit(c.Handler(mux)))))))
}
