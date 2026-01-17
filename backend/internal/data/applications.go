package data

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/cconner57/adoption-os/backend/internal/validator"
)

type Application struct {
	ID           int64           `json:"id"`
	Type         string          `json:"type"`          // 'volunteer', 'adoption', 'surrender'
	Status       string          `json:"status"`        // 'pending', 'approved', 'denied', 'needs_info'
	Data         json.RawMessage `json:"data"`          // Full form data
	OriginalHTML string          `json:"original_html"` // Generated email HTML
	CreatedAt    time.Time       `json:"created_at"`
	UpdatedAt    time.Time       `json:"updated_at"`
	Version      int32           `json:"version"`
}

type ApplicationModel struct {
	DB *sql.DB
}

func (m ApplicationModel) Insert(app *Application) error {
	query := `
		INSERT INTO applications (type, status, data, original_html)
		VALUES ($1, $2, $3, $4)
		RETURNING id, created_at, version`

	args := []any{app.Type, app.Status, app.Data, app.OriginalHTML}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	// Insert Metric
	// We do this concurrently or as part of same logic. If metric fails, app should still succeed?
	// Logic: "In the database, we'll want to keep track of how many of each applications we get per year."
	metricType := "application_received"
	metricData := fmt.Sprintf(`{"type": "%s", "year": %d}`, app.Type, time.Now().Year())
	m.DB.ExecContext(ctx, "INSERT INTO metrics (event_type, event_data) VALUES ($1, $2)", metricType, metricData)

	return m.DB.QueryRowContext(ctx, query, args...).Scan(&app.ID, &app.CreatedAt, &app.Version)
}

func (m ApplicationModel) ArchiveOldPending(ctx context.Context) error {
	// Revised Strategy: "Sanitize" instead of Delete.
	// 1. Select IDs and Data of expiring applications to process sanitization in Go (easier than complex SQL JSONB manipulation for variable fields)
	// Or use SQL if possible. SQL is safer for atomicity.
	// But we have different rules for different types.
	// Let's select them, sanitize in memory, and update.
	// Limit to reasonable batch size.

	query := `
		SELECT id, type, data 
		FROM applications 
		WHERE status = 'pending' 
		AND created_at < NOW() - INTERVAL '7 days'
		LIMIT 100`

	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		return err
	}
	defer rows.Close()

	// Prepare updates
	type minimalData struct {
		FirstName     string `json:"firstName,omitempty"`
		LastName      string `json:"lastName,omitempty"`
		Email         string `json:"email,omitempty"`
		ApplicantName string `json:"applicantName,omitempty"` // Sometimes stored as NameFull
		// Surrender Specific
		AnimalName           string `json:"animalName,omitempty"`
		AnimalAge            string `json:"animalAge,omitempty"`
		AnimalWhySurrendered string `json:"animalWhySurrendered,omitempty"`
	}

	for rows.Next() {
		var id int64
		var appType string
		var rawData []byte

		if err := rows.Scan(&id, &appType, &rawData); err != nil {
			continue
		}

		// Parse full data
		var fullMap map[string]interface{}
		if err := json.Unmarshal(rawData, &fullMap); err != nil {
			continue // corrupted data, maybe just force delete? ignoring for now.
		}

		// Extract Minimal Data
		sanitized := minimalData{}

		// Helper to safely get string
		getString := func(key string) string {
			if v, ok := fullMap[key].(string); ok {
				return v
			}
			return ""
		}

		sanitized.FirstName = getString("firstName")
		sanitized.LastName = getString("lastName")
		sanitized.Email = getString("email")
		if sanitized.Email == "" {
			sanitized.Email = getString("Email") // Case sensitivity check
		}
		sanitized.ApplicantName = getString("nameFull") // fallback if First/Last empty?

		if appType == "surrender" {
			sanitized.AnimalName = getString("animalName")
			sanitized.AnimalAge = getString("animalAge")
			sanitized.AnimalWhySurrendered = getString("animalWhySurrendered")
		}

		newJSON, _ := json.Marshal(sanitized)

		// Update Row
		updateQuery := `
			UPDATE applications 
			SET status = 'autodeleted', original_html = NULL, data = $1, updated_at = NOW()
			WHERE id = $2`

		_, _ = m.DB.ExecContext(ctx, updateQuery, newJSON, id)
	}

	// Double check rows error
	if err = rows.Err(); err != nil {
		return err
	}

	return nil
}

func (m ApplicationModel) Get(id int64) (*Application, error) {
	query := `
		SELECT id, type, status, data, original_html, created_at, updated_at, version
		FROM applications
		WHERE id = $1`

	var app Application
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := m.DB.QueryRowContext(ctx, query, id).Scan(
		&app.ID,
		&app.Type,
		&app.Status,
		&app.Data,
		&app.OriginalHTML,
		&app.CreatedAt,
		&app.UpdatedAt,
		&app.Version,
	)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}

	return &app, nil
}

func (m ApplicationModel) GetAll(typeFilter string, statusFilter string, yearFilter int, filters Filters) ([]*Application, Metadata, error) {
	// Construct query with dynamic filters
	whereClause := "WHERE 1=1"
	args := []any{}
	argCount := 1

	if typeFilter != "" && typeFilter != "all" {
		whereClause += fmt.Sprintf(" AND type = $%d", argCount)
		args = append(args, typeFilter)
		argCount++
	}

	if statusFilter != "" && statusFilter != "all" {
		whereClause += fmt.Sprintf(" AND status = $%d", argCount)
		args = append(args, statusFilter)
		argCount++
	}

	if yearFilter > 0 {
		whereClause += fmt.Sprintf(" AND EXTRACT(YEAR FROM created_at) = $%d", argCount)
		args = append(args, yearFilter)
		argCount++
	}

	totalRecordsQuery := fmt.Sprintf("SELECT count(*) FROM applications %s", whereClause)
	var totalRecords int
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := m.DB.QueryRowContext(ctx, totalRecordsQuery, args...).Scan(&totalRecords)
	if err != nil {
		return nil, Metadata{}, err
	}

	query := fmt.Sprintf(`
		SELECT id, type, status, data, created_at, updated_at, version
		FROM applications
		%s
		ORDER BY %s %s, id ASC
		LIMIT $%d OFFSET $%d`, whereClause, filters.sortColumn(), filters.sortDirection(), argCount, argCount+1)

	args = append(args, filters.limit(), filters.offset())

	rows, err := m.DB.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, Metadata{}, err
	}
	defer rows.Close()

	applications := []*Application{}
	for rows.Next() {
		var app Application
		// Note: We deliberately skip original_html in list view to save bandwidth
		err := rows.Scan(
			&app.ID,
			&app.Type,
			&app.Status,
			&app.Data,
			&app.CreatedAt,
			&app.UpdatedAt,
			&app.Version,
		)
		if err != nil {
			return nil, Metadata{}, err
		}
		applications = append(applications, &app)
	}

	if err = rows.Err(); err != nil {
		return nil, Metadata{}, err
	}

	metadata := calculateMetadata(totalRecords, filters.Page, filters.PageSize)
	return applications, metadata, nil
}

func (m ApplicationModel) Update(app *Application) error {
	query := `
		UPDATE applications 
		SET status = $1, updated_at = NOW(), version = version + 1
		WHERE id = $2 AND version = $3
		RETURNING version`

	args := []any{
		app.Status,
		app.ID,
		app.Version,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := m.DB.QueryRowContext(ctx, query, args...).Scan(&app.Version)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return ErrEditConflict
		default:
			return err
		}
	}

	return nil
}

func ValidateApplication(v *validator.Validator, app *Application) {
	v.Check(app.Type != "", "type", "must be provided")
	v.Check(app.Status != "", "status", "must be provided")
}
