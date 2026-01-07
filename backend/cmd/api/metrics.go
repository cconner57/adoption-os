package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/cconner57/adoption-os/backend/internal/data"
)

func (app *application) submitMetric(w http.ResponseWriter, r *http.Request) {
	var input struct {
		EventType string          `json:"eventType"`
		EventData json.RawMessage `json:"eventData"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	if input.EventType == "" {
		app.failedValidationResponse(w, r, map[string]string{"eventType": "must be provided"})
		return
	}

	metric := &data.Metric{
		EventType: input.EventType,
		EventData: input.EventData,
	}

	err = app.models.Metrics.Insert(metric)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	app.writeJSON(w, http.StatusCreated, envelope{"message": "metric recorded"}, nil)
	fmt.Printf("Metric Recorded: %s\n", input.EventType)
}
