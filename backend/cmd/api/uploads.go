package main

import (
	"fmt"
	"net/http"
	"path/filepath"
	"time"

	"github.com/cconner57/adoption-os/backend/internal/data"
)

const (
	MaxUploadSize = 10 << 20 // 10 MB
)

func (app *application) uploadPetPhotoHandler(w http.ResponseWriter, r *http.Request) {
	// Parse ID from URL
	id := r.PathValue("id")
	if id == "" {
		app.badRequestResponse(w, r, fmt.Errorf("missing pet id"))
		return
	}

	// Block uploads in local development
	if app.config.env != "production" {
		app.badRequestResponse(w, r, fmt.Errorf("uploads are disabled in %s environment", app.config.env))
		return
	}

	// Limit upload size
	r.Body = http.MaxBytesReader(w, r.Body, MaxUploadSize)
	if err := r.ParseMultipartForm(MaxUploadSize); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	file, _, err := r.FormFile("photo")
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}
	defer file.Close()

	// Detect content type
	buffer := make([]byte, 512)
	_, err = file.Read(buffer)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
	file.Seek(0, 0)
	contentType := http.DetectContentType(buffer)

	// Validate Content Type
	if contentType != "image/jpeg" && contentType != "image/png" {
		app.badRequestResponse(w, r, fmt.Errorf("unsupported content type: %s", contentType))
		return
	}

	// Prepare directories
	// Structure: [ASSETS_DIR]/pets/[id]/photos/
	petDir := filepath.Join(app.config.assetsDir, "pets", id, "photos")

	// Create timestamp-based filename
	timestamp := time.Now().Unix()
	baseName := fmt.Sprintf("%s_%d.jpg", id, timestamp) // Input name for processor

	// Process and Save
	result, err := data.SaveAndProcessImage(file, baseName, petDir)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	// Construct public URLs (relative to the asset root)
	// The frontend/proxy maps /pet-photos/ -> [ASSETS_DIR]
	// We return the relative path from ASSETS_DIR so the frontend can prepend /pet-photos/

	thumbUrl := fmt.Sprintf("pets/%s/photos/%s", id, result.ThumbnailPath)
	largeUrl := fmt.Sprintf("pets/%s/photos/%s", id, result.LargePath)

	response := envelope{
		"url":          largeUrl,
		"thumbnailUrl": thumbUrl,
	}

	app.writeJSON(w, http.StatusOK, response, nil)
}
