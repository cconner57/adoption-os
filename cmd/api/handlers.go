package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"idohr-be/internal/models"
)

// Application struct holds dependencies (like the DB)
// so we don't have to pass them around manually everywhere.
type Application struct {
	DB *sql.DB
}

func (app *Application) Home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Pet Adoption API is Running! 🐶"))
}

func (app *Application) GetAllPets(w http.ResponseWriter, r *http.Request) {
	// 1. SELECT: We must explicitly select every column we want to scan.
	// We use COALESCE to ensure we get empty JSON objects "{}" instead of NULLs,
	// which prevents crashes when unmarshalling.
	query := `
		SELECT 
			id, name, species, sex, 
			COALESCE(slug, ''), 
			COALESCE(litter_name, ''),
			created_at, updated_at,
			COALESCE(physical, '{}'),
			COALESCE(behavior, '{}'),
			COALESCE(medical, '{}'),
			COALESCE(descriptions, '{}'),
			COALESCE(details, '{}'),
			COALESCE(adoption, '{}'),
			COALESCE(foster, '{}'),
			COALESCE(returned, '{}'),
			COALESCE(sponsored, '{}'),
			COALESCE(photos, '[]'),
			COALESCE(profile_settings, '{}')
		FROM pets 
		ORDER BY created_at DESC
		LIMIT 50`

	rows, err := app.DB.Query(query)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	defer rows.Close()

	var pets []models.Pet
	for rows.Next() {
		var p models.Pet

		// temporary variables to hold raw data
		var litterName, slug string
		var physical, behavior, medical, descriptions, details, adoption, foster, returned, sponsored, photos, profileSettings []byte

		// 2. SCAN: The order here MUST match the order in the SELECT query above.
		err := rows.Scan(
			&p.ID, &p.Name, &p.Species, &p.Sex,
			&slug,
			&litterName,
			&p.CreatedAt, &p.UpdatedAt,
			&physical, &behavior, &medical, &descriptions, &details,
			&adoption, &foster, &returned, &sponsored, &photos, &profileSettings,
		)
		if err != nil {
			log.Println("Error scanning row:", err)
			continue
		}

		// 3. MAP: Handle the standard fields
		if slug != "" {
			p.Slug = &slug
		}
		if litterName != "" {
			p.LitterName = &litterName
		}

		// 4. UNMARSHAL: Convert the raw JSON bytes into our Go Structs
		// We ignore errors here because we coalesced to '{}' in SQL, so it shouldn't fail.
		_ = json.Unmarshal(physical, &p.Physical)
		_ = json.Unmarshal(behavior, &p.Behavior)
		_ = json.Unmarshal(medical, &p.Medical)
		_ = json.Unmarshal(descriptions, &p.Descriptions)
		_ = json.Unmarshal(details, &p.Details)
		_ = json.Unmarshal(adoption, &p.Adoption)
		_ = json.Unmarshal(foster, &p.Foster)
		_ = json.Unmarshal(returned, &p.Returned)
		_ = json.Unmarshal(sponsored, &p.Sponsored)
		_ = json.Unmarshal(photos, &p.Photos)
		_ = json.Unmarshal(profileSettings, &p.ProfileSettings)

		pets = append(pets, p)
	}

	if pets == nil {
		pets = []models.Pet{}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(pets)
}

func (app *Application) GetAvailablePets(w http.ResponseWriter, r *http.Request) {
	// 1. UPDATE QUERY:
	// - Filter using the JSON operator ->> to look inside the 'details' column
	// - Select ALL columns (including JSONB) so the frontend gets the full pet object
	query := `
		SELECT 
			id, name, species, sex, 
			COALESCE(slug, ''), 
			COALESCE(litter_name, ''),
			created_at, updated_at,
			COALESCE(physical, '{}'),
			COALESCE(behavior, '{}'),
			COALESCE(medical, '{}'),
			COALESCE(descriptions, '{}'),
			COALESCE(details, '{}'),
			COALESCE(adoption, '{}'),
			COALESCE(foster, '{}'),
			COALESCE(returned, '{}'),
			COALESCE(sponsored, '{}'),
			COALESCE(photos, '[]'),
			COALESCE(profile_settings, '{}')
		FROM pets 
		WHERE details->>'status' = 'available'
		ORDER BY created_at DESC
		LIMIT 50`

	rows, err := app.DB.Query(query)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	defer rows.Close()

	var pets []models.Pet
	for rows.Next() {
		var p models.Pet

		// Temporary variables for raw data
		var litterName, slug string
		var physical, behavior, medical, descriptions, details, adoption, foster, returned, sponsored, photos, profileSettings []byte

		// 2. SCAN: Must match the SELECT order above
		err := rows.Scan(
			&p.ID, &p.Name, &p.Species, &p.Sex,
			&slug,
			&litterName,
			&p.CreatedAt, &p.UpdatedAt,
			&physical, &behavior, &medical, &descriptions, &details,
			&adoption, &foster, &returned, &sponsored, &photos, &profileSettings,
		)
		if err != nil {
			log.Println("Error scanning row:", err)
			continue
		}

		// 3. MAP & UNMARSHAL
		if slug != "" {
			p.Slug = &slug
		}
		if litterName != "" {
			p.LitterName = &litterName
		}

		_ = json.Unmarshal(physical, &p.Physical)
		_ = json.Unmarshal(behavior, &p.Behavior)
		_ = json.Unmarshal(medical, &p.Medical)
		_ = json.Unmarshal(descriptions, &p.Descriptions)
		_ = json.Unmarshal(details, &p.Details)
		_ = json.Unmarshal(adoption, &p.Adoption)
		_ = json.Unmarshal(foster, &p.Foster)
		_ = json.Unmarshal(returned, &p.Returned)
		_ = json.Unmarshal(sponsored, &p.Sponsored)
		_ = json.Unmarshal(photos, &p.Photos)
		_ = json.Unmarshal(profileSettings, &p.ProfileSettings)

		pets = append(pets, p)
	}

	if pets == nil {
		pets = []models.Pet{}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(pets)
}

func (app *Application) GetSpotlightPets(w http.ResponseWriter, r *http.Request) {
	// 1. UPDATE QUERY:
	// - Select ALL columns so the card component works on the frontend
	// - Filter by profile_settings (isSpotlightFeatured) AND details (status)
	query := `
		SELECT 
			id, name, species, sex, 
			COALESCE(slug, ''), 
			COALESCE(litter_name, ''),
			created_at, updated_at,
			COALESCE(physical, '{}'),
			COALESCE(behavior, '{}'),
			COALESCE(medical, '{}'),
			COALESCE(descriptions, '{}'),
			COALESCE(details, '{}'),
			COALESCE(adoption, '{}'),
			COALESCE(foster, '{}'),
			COALESCE(returned, '{}'),
			COALESCE(sponsored, '{}'),
			COALESCE(photos, '[]'),
			COALESCE(profile_settings, '{}')
		FROM pets 
		WHERE profile_settings->>'isSpotlightFeatured' = 'true'
		AND details->>'status' = 'available'
		LIMIT 10`

	rows, err := app.DB.Query(query)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	defer rows.Close()

	var pets []models.Pet
	for rows.Next() {
		var p models.Pet

		// Temporary variables
		var litterName, slug string
		var physical, behavior, medical, descriptions, details, adoption, foster, returned, sponsored, photos, profileSettings []byte

		// 2. SCAN: Must match the SELECT order
		err := rows.Scan(
			&p.ID, &p.Name, &p.Species, &p.Sex,
			&slug,
			&litterName,
			&p.CreatedAt, &p.UpdatedAt,
			&physical, &behavior, &medical, &descriptions, &details,
			&adoption, &foster, &returned, &sponsored, &photos, &profileSettings,
		)
		if err != nil {
			log.Println("Error scanning row:", err)
			continue
		}

		// 3. MAP & UNMARSHAL
		if slug != "" {
			p.Slug = &slug
		}
		if litterName != "" {
			p.LitterName = &litterName
		}

		_ = json.Unmarshal(physical, &p.Physical)
		_ = json.Unmarshal(behavior, &p.Behavior)
		_ = json.Unmarshal(medical, &p.Medical)
		_ = json.Unmarshal(descriptions, &p.Descriptions)
		_ = json.Unmarshal(details, &p.Details)
		_ = json.Unmarshal(adoption, &p.Adoption)
		_ = json.Unmarshal(foster, &p.Foster)
		_ = json.Unmarshal(returned, &p.Returned)
		_ = json.Unmarshal(sponsored, &p.Sponsored)
		_ = json.Unmarshal(photos, &p.Photos)
		_ = json.Unmarshal(profileSettings, &p.ProfileSettings)

		pets = append(pets, p)
	}

	if pets == nil {
		pets = []models.Pet{}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(pets)
}
