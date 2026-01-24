package main

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/cconner57/adoption-os/backend/internal/data"
)

func (app *application) listCampaignsHandler(w http.ResponseWriter, r *http.Request) {
	status := r.URL.Query().Get("status")

	campaigns, err := app.models.Marketing.GetCampaigns(status)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	app.writeJSON(w, http.StatusOK, envelope{"campaigns": campaigns}, nil)
}

func (app *application) getCampaignHandler(w http.ResponseWriter, r *http.Request) {
	id := app.readStringParam(r, "id")
	if id == "" {
		app.notFoundResponse(w, r)
		return
	}

	campaign, err := app.models.Marketing.GetCampaign(id)
	if err != nil {
		if errors.Is(err, data.ErrRecordNotFound) {
			app.notFoundResponse(w, r)
		} else {
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	app.writeJSON(w, http.StatusOK, envelope{"campaign": campaign}, nil)
}

func (app *application) createCampaignHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		ID          string  `json:"id"`
		Name        string  `json:"name"`
		Status      string  `json:"status"`
		StartDate   string  `json:"startDate"`
		EndDate     string  `json:"endDate"`
		Goal        string  `json:"goal"`
		Progress    int     `json:"progress"`
		Metric      string  `json:"metric"`
		Type        string  `json:"type"`
		Prize       string  `json:"prize"`
		TicketPrice float64 `json:"ticketPrice"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	// Basic Validation
	if input.Name == "" {
		app.badRequestResponse(w, r, errors.New("campaign name is required"))
		return
	}

	campaign := &data.Campaign{
		ID:          input.ID,
		Name:        input.Name,
		Status:      input.Status,
		StartDate:   input.StartDate,
		EndDate:     input.EndDate,
		Goal:        input.Goal,
		Progress:    input.Progress,
		Metric:      input.Metric,
		Type:        input.Type,
		Prize:       input.Prize,
		TicketPrice: input.TicketPrice,
	}

	// Generate ID if needed (though frontend often sends UUIDs or slugs)
	if campaign.ID == "" {
		// Simple fallback if no ID provided, but typically UUIDs are generated
		// For now, let's assume client sends it or we rely on DB default if SERIAL (but we used TEXT PRIMARY KEY)
		// Let's enforce ID from client or generate one.
		// Since we didn't add a UUID lib, we'll rely on client or basic string
		// But wait, the migration doesn't auto-generate IDs.
		// Let's assume the frontend sends an ID or we error.
		app.badRequestResponse(w, r, errors.New("id is required"))
		return
	}

	err = app.models.Marketing.CreateCampaign(campaign)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	app.writeJSON(w, http.StatusCreated, envelope{"campaign": campaign}, nil)
}

func (app *application) updateCampaignHandler(w http.ResponseWriter, r *http.Request) {
	id := app.readStringParam(r, "id")
	if id == "" {
		app.notFoundResponse(w, r)
		return
	}

	campaign, err := app.models.Marketing.GetCampaign(id)
	if err != nil {
		if errors.Is(err, data.ErrRecordNotFound) {
			app.notFoundResponse(w, r)
		} else {
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	var input struct {
		Name        *string  `json:"name"`
		Status      *string  `json:"status"`
		StartDate   *string  `json:"startDate"`
		EndDate     *string  `json:"endDate"`
		Goal        *string  `json:"goal"`
		Progress    *int     `json:"progress"`
		Metric      *string  `json:"metric"`
		Type        *string  `json:"type"`
		Prize       *string  `json:"prize"`
		TicketPrice *float64 `json:"ticketPrice"`
		WinnerID    *string  `json:"winnerId"`
		// Participants is complex, might need separate handling or full replacement
	}

	err = app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	if input.Name != nil {
		campaign.Name = *input.Name
	}
	if input.Status != nil {
		campaign.Status = *input.Status
	}
	if input.StartDate != nil {
		campaign.StartDate = *input.StartDate
	}
	if input.EndDate != nil {
		campaign.EndDate = *input.EndDate
	}
	if input.Goal != nil {
		campaign.Goal = *input.Goal
	}
	if input.Progress != nil {
		campaign.Progress = *input.Progress
	}
	if input.Metric != nil {
		campaign.Metric = *input.Metric
	}
	if input.Type != nil {
		campaign.Type = *input.Type
	}
	if input.Prize != nil {
		campaign.Prize = *input.Prize
	}
	if input.TicketPrice != nil {
		campaign.TicketPrice = *input.TicketPrice
	}
	if input.WinnerID != nil {
		campaign.WinnerID = *input.WinnerID
	}

	err = app.models.Marketing.UpdateCampaign(campaign)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	if campaign.Progress >= 100 {
		go app.notifier.SendToAll(fmt.Sprintf("Campaign '%s' has reached 100%% goal! 🎉", campaign.Name))
	}

	app.writeJSON(w, http.StatusOK, envelope{"campaign": campaign}, nil)
}
