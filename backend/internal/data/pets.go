package data

import (
	"context"
	"database/sql"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"
)

// Define the Pet struct (matching your DB table)
type SpotlightPet struct {
	ID           string          `json:"id"`
	Name         string          `json:"name"`
	Descriptions json.RawMessage `json:"descriptions"`
	Photos       json.RawMessage `json:"photos"`
}

// Define the Model wrapper
type PetModel struct {
	DB *sql.DB
}

func (m PetModel) GetSpotlight() ([]*SpotlightPet, error) {
	// Guard against no DB connection (CSV fallback mode)
	if m.DB == nil {
		fileName := "Master Pet List - Cats Master List.csv"
		file, err := os.Open(fileName)
		if os.IsNotExist(err) {
			file, err = os.Open("backend/" + fileName)
		}
		if err != nil {
			fmt.Println("GetSpotlight: CSV not found, returning empty list")
			return []*SpotlightPet{}, nil
		}
		defer file.Close()

		reader := csv.NewReader(file)
		reader.FieldsPerRecord = -1
		records, err := reader.ReadAll()
		if err != nil {
			fmt.Println("GetSpotlight: Error reading CSV:", err)
			return []*SpotlightPet{}, nil
		}

		pets := []*SpotlightPet{}

		for i, row := range records {
			if i < 3 {
				continue
			} // Skip header lines
			if len(row) < 3 {
				continue
			}

			// CSV Structure: ..., Is Spotlight Featured, Spotlight Description
			lenRow := len(row)
			if lenRow < 2 {
				continue
			}

			// Check "Is Spotlight Featured"
			isFeatured := strings.TrimSpace(row[lenRow-2])
			description := ""
			lastCol := strings.TrimSpace(row[lenRow-1])

			if strings.EqualFold(isFeatured, "TRUE") {
				description = lastCol
			} else if strings.EqualFold(lastCol, "TRUE") {
				isFeatured = "TRUE"
			} else {
				continue
			}

			if strings.EqualFold(isFeatured, "TRUE") {
				id := row[0]
				name := row[2]

				descMap := map[string]string{"spotlight": description}
				descBytes, _ := json.Marshal(descMap)

				pet := &SpotlightPet{
					ID:           id,
					Name:         name,
					Descriptions: json.RawMessage(descBytes),
					Photos:       json.RawMessage(`[]`),
				}
				pets = append(pets, pet)
				if len(pets) >= 4 {
					break
				}
			}
		}
		return pets, nil
	}

	query := `
        SELECT id, name, COALESCE(descriptions, '{}'), COALESCE(photos, '[]')
        FROM pets
        WHERE profile_settings->>'is_spotlight_featured' = 'true'
           OR profile_settings->>'isSpotlightFeatured' = 'true'
        LIMIT 4`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		fmt.Println("Query Error:", err)
		return nil, err
	}
	defer rows.Close()

	pets := []*SpotlightPet{}

	for rows.Next() {
		var pet SpotlightPet
		err := rows.Scan(
			&pet.ID,
			&pet.Name,
			&pet.Descriptions,
			&pet.Photos,
		)
		if err != nil {
			fmt.Println("Scan Error:", err)
			return nil, err
		}
		pets = append(pets, &pet)
	}

	return pets, nil
}

// Helper to parse date from CSV "M/D/YYYY"
func parseDate(dateStr string) *time.Time {
	dateStr = strings.TrimSpace(dateStr)
	if dateStr == "" || strings.Contains(dateStr, "Due") || dateStr == "----" {
		return nil
	}
	layout := "1/2/2006"
	t, err := time.Parse(layout, dateStr)
	if err != nil {
		return nil
	}
	return &t
}

func (m PetModel) GetAdoptedCount(year int) (int, error) {
	// 1. Try DB if available
	if m.DB != nil {
		if err := m.DB.Ping(); err == nil {
			ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
			defer cancel()

			yearPrefix := fmt.Sprintf("%%%d%%", year)

			query := `
				SELECT count(*)
				FROM pets
				WHERE LOWER(status) = 'adopted'
				AND adoption->>'date' LIKE $1
			`
			var count int

			err := m.DB.QueryRowContext(ctx, query, yearPrefix).Scan(&count)
			if err == nil {
				return count, nil
			}
		}
	}

	// 2. Fallback to CSV
	fileName := "Master Pet List - Cats Master List.csv"
	file, err := os.Open(fileName)
	if os.IsNotExist(err) {
		file, err = os.Open("backend/" + fileName)
	}
	if err != nil {
		return 0, fmt.Errorf("could not open CSV fallback: %v", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.FieldsPerRecord = -1
	records, err := reader.ReadAll()
	if err != nil {
		return 0, err
	}

	count := 0
	targetYear := year

	for i, row := range records {
		if i == 0 || len(row) < 39 {
			continue
		}
		// status is index 37, date is 39
		status := strings.ToLower(strings.TrimSpace(row[37]))
		if status != "adopted" {
			continue
		}

		datePtr := parseDate(row[39])
		if datePtr != nil {
			if datePtr.Year() == targetYear {
				count++
			}
		}
	}

	return count, nil
}

type SitemapPet struct {
	ID        string
	UpdatedAt time.Time
}

func (m PetModel) GetAllAvailablePets() ([]*SitemapPet, error) {
	if m.DB == nil {
		return []*SitemapPet{}, nil
	}

	query := `
		SELECT id, updated_at
		FROM pets
		WHERE status = 'available'
	`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var pets []*SitemapPet
	for rows.Next() {
		var p SitemapPet
		err := rows.Scan(&p.ID, &p.UpdatedAt)
		if err != nil {
			return nil, err
		}
		pets = append(pets, &p)
	}
	return pets, nil
}

// Full Pet Struct for Grid/List Views
type Pet struct {
	ID        string    `json:"id"`
	Slug      *string   `json:"slug,omitempty"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`

	Name       string  `json:"name"`
	Species    string  `json:"species"`
	Sex        string  `json:"sex"`
	LitterName *string `json:"litterName,omitempty"`

	// JSONB Fields
	Physical        json.RawMessage `json:"physical"`
	Behavior        json.RawMessage `json:"behavior"`
	Medical         json.RawMessage `json:"medical"`
	Descriptions    json.RawMessage `json:"descriptions"`
	Details         json.RawMessage `json:"details"`
	Adoption        json.RawMessage `json:"adoption"`
	Foster          json.RawMessage `json:"foster"`
	Returned        json.RawMessage `json:"returned"`
	Sponsored       json.RawMessage `json:"sponsored"`
	Photos          json.RawMessage `json:"photos"`
	ProfileSettings json.RawMessage `json:"profileSettings"`
}

func (m PetModel) GetAll(status string, search string) ([]*Pet, error) {
	if m.DB == nil {
		return []*Pet{}, nil
	}

	query := `
		SELECT 
			id, 
			name, 
			COALESCE(sex, 'unknown'),
			COALESCE(slug, ''),
			COALESCE(litter_name, ''),
			created_at, 
			updated_at,
			COALESCE(physical, '{}'),
			COALESCE(behavior, '{}'),
			COALESCE(medical, '{}'),
			COALESCE(descriptions, '{}'),
			COALESCE(details, '{}'::jsonb) || jsonb_build_object('status', status),
			COALESCE(adoption, '{}'),
			COALESCE(foster, '{}'),
			COALESCE(returned, '{}'),
			COALESCE(sponsored, '{}'),
			COALESCE(photos, '[]'),
			COALESCE(profile_settings, '{}')
		FROM pets
		WHERE 1=1
	`
	cleanArgs := []interface{}{}
	cleanArgCount := 1

	if status != "" && status != "all" {
		query += fmt.Sprintf(" AND LOWER(status) = LOWER($%d)", cleanArgCount)
		cleanArgs = append(cleanArgs, status)
		cleanArgCount++
	}

	if search != "" {
		searchT := "%" + strings.ToLower(search) + "%"
		query += fmt.Sprintf(" AND (LOWER(name) LIKE $%d OR LOWER(physical->>'breed') LIKE $%d)", cleanArgCount, cleanArgCount+1)
		cleanArgs = append(cleanArgs, searchT, searchT)
		cleanArgCount += 2
	}

	query += " ORDER BY name ASC"

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	rows, err := m.DB.QueryContext(ctx, query, cleanArgs...)
	if err != nil {
		fmt.Println("GetAll Query Error:", err)
		return nil, err
	}
	defer rows.Close()

	pets := []*Pet{}
	for rows.Next() {
		var p Pet
		var slug, litterName string

		err := rows.Scan(
			&p.ID,
			&p.Name,
			&p.Sex,
			&slug,
			&litterName,
			&p.CreatedAt,
			&p.UpdatedAt,
			&p.Physical,
			&p.Behavior,
			&p.Medical,
			&p.Descriptions,
			&p.Details,
			&p.Adoption,
			&p.Foster,
			&p.Returned,
			&p.Sponsored,
			&p.Photos,
			&p.ProfileSettings,
		)
		if err != nil {
			fmt.Println("GetAll Scan Error:", err)
			return nil, err
		}

		p.Species = "cat"
		if slug != "" {
			p.Slug = &slug
		}
		if litterName != "" {
			p.LitterName = &litterName
		}

		pets = append(pets, &p)
	}

	return pets, nil
}

func (m PetModel) Update(p *Pet) error {
	if m.DB == nil {
		return fmt.Errorf("database connection not available")
	}

	query := `
		UPDATE pets
		SET 
			name = $1,
			sex = $2,
			updated_at = NOW(),
			physical = $3,
			behavior = $4,
			medical = $5,
			descriptions = $6,
			details = $7,
			adoption = $8,
			foster = $9,
			returned = $10,
			sponsored = $11,
			photos = $12,
			profile_settings = $13,
			status = $14,
			litter_name = $16,
			species = $17
		WHERE id = $15
	`

	var detailsMap map[string]interface{}
	status := "available"
	if len(p.Details) > 0 {
		if err := json.Unmarshal(p.Details, &detailsMap); err == nil {
			if s, ok := detailsMap["status"].(string); ok {
				status = s
			}
		}
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	args := []interface{}{
		p.Name,
		p.Sex,
		p.Physical,
		p.Behavior,
		p.Medical,
		p.Descriptions,
		p.Details,
		p.Adoption,
		p.Foster,
		p.Returned,
		p.Sponsored,
		p.Photos,
		p.ProfileSettings,
		status,       // $14
		p.ID,         // $15
		p.LitterName, // $16
		p.Species,    // $17
	}

	result, err := m.DB.ExecContext(ctx, query, args...)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return ErrRecordNotFound
	}

	return nil
}

func (m PetModel) Insert(p *Pet) error {
	if m.DB == nil {
		return fmt.Errorf("database connection not available")
	}

	query := `
		INSERT INTO pets (
			name, sex, physical, behavior, medical, descriptions, 
			details, adoption, foster, returned, sponsored, photos, profile_settings, 
			status, litter_name, species, created_at, updated_at
		)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, NOW(), NOW())
		RETURNING id, created_at, updated_at
	`

	var detailsMap map[string]interface{}
	status := "available"
	if len(p.Details) > 0 {
		if err := json.Unmarshal(p.Details, &detailsMap); err == nil {
			if s, ok := detailsMap["status"].(string); ok {
				status = s
			}
		}
	}

	args := []interface{}{
		p.Name,
		p.Sex,
		p.Physical,
		p.Behavior,
		p.Medical,
		p.Descriptions,
		p.Details,
		p.Adoption,
		p.Foster,
		p.Returned,
		p.Sponsored,
		p.Photos,
		p.ProfileSettings,
		status,
		p.LitterName,
		p.Species,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := m.DB.QueryRowContext(ctx, query, args...).Scan(&p.ID, &p.CreatedAt, &p.UpdatedAt)
	if err != nil {
		return err
	}

	return nil
}

type Adoption struct {
	AdoptedBy          string  `json:"adoptedBy"`
	Date               *string `json:"date"`
	NewAdoptedName     string  `json:"newAdoptedName"`
	Fee                *int    `json:"fee"`
	SurveyCompleted    bool    `json:"surveyCompleted"`
	AdopterContactInfo struct {
		Name  string `json:"name"`
		Email string `json:"email"`
		Phone string `json:"phone"`
	} `json:"adopterContactInfo"`
}

// SeedAdoptionDates Upsert Logic
func (m PetModel) SeedAdoptionDates() error {
	fmt.Println("Starting Adoption Date & DOB Seeding...")

	file, err := os.Open("Master Pet List - Cats Master List.csv")
	if err != nil {
		return err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.FieldsPerRecord = -1
	records, err := reader.ReadAll()
	if err != nil {
		return err
	}

	nameIdx := 2
	dobIdx := 3
	breedIdx := 5
	sexIdx := 10
	intakeIdx := 35
	statusIdx := 37
	dateIdx := 39
	spayNeuterDateIdx := 29
	microchipIdIdx := 12

	microchipCompanyIdx := 13

	countUpdated := 0
	countCreated := 0

	// Helper to parse "M/D/YYYY" to "YYYY-MM-DD"
	parseToISO := func(s string) *string {
		s = strings.TrimSpace(s)
		if s == "" || s == "----" || strings.Contains(s, "Due") {
			return nil
		}
		t, err := time.Parse("1/2/2006", s)
		if err != nil {
			return nil
		}
		formatted := t.Format("2006-01-02")
		return &formatted
	}

	for i, row := range records {
		if i < 3 {
			continue
		}
		if len(row) <= dateIdx || len(row) <= sexIdx || len(row) <= breedIdx {
			continue
		}

		name := strings.TrimSpace(row[nameIdx])
		if name == "" {
			continue
		}

		// 1. Prepare Adoption Date
		adoptionDatePtr := parseToISO(row[dateIdx])

		// 2. Prepare DOB
		dobPtr := parseToISO(row[dobIdx])

		adoption := Adoption{Date: adoptionDatePtr}
		adoptionJSON, _ := json.Marshal(adoption)

		// Prepare partial physical update for DOB
		// We use a map to handle explicit null if needed, or just the value
		physicalUpdateMap := map[string]interface{}{}
		if dobPtr != nil {
			physicalUpdateMap["dateOfBirth"] = *dobPtr
		} else {
			physicalUpdateMap["dateOfBirth"] = nil
		}
		physicalUpdateJSON, _ := json.Marshal(physicalUpdateMap)

		// 2.5 Prepare Medical Update (Microchip)
		medicalUpdateMap := map[string]interface{}{}
		microchipId := strings.TrimSpace(row[microchipIdIdx])
		microchipCompany := strings.TrimSpace(row[microchipCompanyIdx])

		microchipMap := map[string]interface{}{}

		if microchipId != "" {
			microchipMap["microchipID"] = microchipId
			medicalUpdateMap["microchipped"] = true
		} else {
			microchipMap["microchipID"] = nil
			// medicalUpdateMap["microchipped"] = false
		}

		if microchipCompany != "" {
			microchipMap["microchipCompany"] = microchipCompany
		} else {
			microchipMap["microchipCompany"] = nil
		}

		medicalUpdateMap["microchip"] = microchipMap

		// 2.6 Prepare Medical Update (Spay/Neuter)
		spayDatePtr := parseToISO(row[spayNeuterDateIdx])
		if spayDatePtr != nil {
			medicalUpdateMap["spayedOrNeutered"] = true
			medicalUpdateMap["spayedOrNeuteredDate"] = *spayDatePtr
		} else {
			medicalUpdateMap["spayedOrNeutered"] = false
			medicalUpdateMap["spayedOrNeuteredDate"] = nil
		}

		medicalUpdateJSON, _ := json.Marshal(medicalUpdateMap)

		// 3. Try UPDATE first
		// We use || to merge the new DOB into the existing physical JSONB object
		// adoption column is replaced fully as per previous logic (or could be merged too, but struct implies full)
		updateQuery := `
			UPDATE pets
			SET adoption = $1,
			    physical = physical || $2,
			    medical = medical || $3
			WHERE LOWER(name) = LOWER($4)
		`
		res, err := m.DB.Exec(updateQuery, adoptionJSON, physicalUpdateJSON, medicalUpdateJSON, name)
		if err != nil {
			fmt.Printf("Error updating %s: %v\n", name, err)
			continue
		}

		rows, _ := res.RowsAffected()
		if rows > 0 {
			countUpdated++
			continue
		}

		// 4. If not found, INSERT new pet
		sex := strings.ToLower(strings.TrimSpace(row[sexIdx]))
		if sex == "" {
			sex = "unknown"
		}

		breed := strings.TrimSpace(row[breedIdx])
		status := strings.ToLower(strings.TrimSpace(row[statusIdx]))
		if status == "" {
			status = "available"
		}

		intake := ""
		if len(row) > intakeIdx {
			intake = strings.TrimSpace(row[intakeIdx])
		}

		// Construct full physical object for new pet
		physicalMap := map[string]interface{}{
			"breed":       breed,
			"dateOfBirth": nil,
		}
		if dobPtr != nil {
			physicalMap["dateOfBirth"] = *dobPtr
		}
		physicalJSON, _ := json.Marshal(physicalMap)

		detailsJSON, _ := json.Marshal(map[string]string{"intakeDate": intake})

		medicalJSON, _ := json.Marshal(medicalUpdateMap) // Reuse the map we built above

		emptyJSON := json.RawMessage("{}")
		emptyArray := json.RawMessage("[]")

		insertQuery := `
			INSERT INTO pets (
				name, species, sex, status, adoption,
				physical, details,
				behavior, medical, descriptions, foster, returned, sponsored, photos, profile_settings,
				created_at, updated_at
			) VALUES (
				$1, 'cat', $2, $3, $4,
				$5, $6,
				$7, $8, $7, $7, $7, $7, $9, $7,
				NOW(), NOW()
			)
		`

		_, err = m.DB.Exec(insertQuery,
			name, sex, status, adoptionJSON,
			physicalJSON, detailsJSON,
			emptyJSON, medicalJSON, emptyArray,
		)

		if err != nil {
			fmt.Printf("Error inserting %s: %v\n", name, err)
		} else {
			countCreated++
			fmt.Printf("Created new pet: %s\n", name)
		}
	}

	fmt.Printf("Seeding completed. Updated: %d, Created: %d\n", countUpdated, countCreated)
	return nil
}
