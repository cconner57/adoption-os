package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"mime/multipart"
	"net/http"
)

// Define the envelope type
type envelope map[string]any

func (app *application) writeJSON(w http.ResponseWriter, status int, data envelope, headers http.Header) error {
	js, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		return err
	}

	js = append(js, '\n')

	for key, value := range headers {
		w.Header()[key] = value
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(js)

	return nil
}

func (app *application) readJSON(w http.ResponseWriter, r *http.Request, dst any) error {
	maxBytes := 1_048_576 // 1MB
	r.Body = http.MaxBytesReader(w, r.Body, int64(maxBytes))

	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()

	err := dec.Decode(dst)
	if err != nil {
		// If the error is regarding the size
		var maxBytesError *http.MaxBytesError
		if errors.As(err, &maxBytesError) {
			return fmt.Errorf("body must not be larger than %d bytes", maxBytes)
		}
		return err
	}

	return nil
}

func (app *application) badRequestResponse(w http.ResponseWriter, r *http.Request, err error) {
	app.errorResponse(w, r, http.StatusBadRequest, err.Error())
}

func (app *application) serverErrorResponse(w http.ResponseWriter, r *http.Request, err error) {
	app.logError(r, err)
	// message := "the server encountered a problem and could not process your request"
	app.errorResponse(w, r, http.StatusInternalServerError, err.Error()) // <--- EXPOSE ERROR
}

func (app *application) errorResponse(w http.ResponseWriter, r *http.Request, status int, message any) {
	app.JSONError(w, status, message)
}

func (app *application) JSONResponse(w http.ResponseWriter, status int, data any) error {
	return app.writeJSON(w, status, envelope{"status": "success", "data": data}, nil)
}

func (app *application) JSONError(w http.ResponseWriter, status int, message any) error {
	return app.writeJSON(w, status, envelope{"status": "error", "message": message}, nil)
}

func (app *application) logError(r *http.Request, err error) {
	var (
		method = r.Method
		url    = r.URL.String()
		reqID  = app.contextGetRequestID(r)
	)

	app.logger.Error(err.Error(), "method", method, "url", url, "request_id", reqID)
}

func (app *application) failedValidationResponse(w http.ResponseWriter, r *http.Request, errors map[string]string) {
	app.errorResponse(w, r, http.StatusUnprocessableEntity, errors)
}

func (app *application) ValidateImageFile(fileHeader *multipart.FileHeader) error {
	file, err := fileHeader.Open()
	if err != nil {
		return err
	}
	defer file.Close()

	// Read first 512 bytes
	buffer := make([]byte, 512)
	_, err = file.Read(buffer)
	if err != nil {
		return fmt.Errorf("failed to read file content: %w", err)
	}

	// Reset file pointer (important if we were to reuse this file object, though typically we don't here)
	// file.Seek(0, 0)

	contentType := http.DetectContentType(buffer)

	switch contentType {
	case "image/jpeg", "image/png", "image/webp":
		return nil
	default:
		return fmt.Errorf("invalid file type: %s. Only jpeg, png, and webp are allowed", contentType)
	}
}
