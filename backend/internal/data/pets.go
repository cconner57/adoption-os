package data // <--- Must be 'data'

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
		// Basic parsing looking for "Is Spotlight Featured" in the last column (index 53 roughly, or just check last)
		// Header is row 2 (index 2 in 0-based slice if we have empty lines at top)
		// Let's assume we iterate and look for 'TRUE' at the end

		for i, row := range records {
			if i < 3 {
				continue
			} // Skip header lines
			if len(row) < 3 {
				continue
			}

			// CSV Structure: ..., Is Spotlight Featured, Spotlight Description
			// "Is Spotlight Featured" should be the second to last column now
			// "Spotlight Description" should be the last column

			// Determine indices based on row length to be dynamic or check strict count
			// Let's assume the last two columns we just touched are at the end.
			lenRow := len(row)
			if lenRow < 2 {
				continue
			}

			// Check "Is Spotlight Featured"
			isFeatured := strings.TrimSpace(row[lenRow-2])

			// Fallback: if user didn't update all rows to have the new column,
			// the "TRUE" might still be at the very end for other rows (though we only marked 4).
			// But for the target 4, we added a column.
			// Let's safe check: if row ends with TRUE, it means no description.
			// If row ends with description, the one before is TRUE.

			description := ""
			lastCol := strings.TrimSpace(row[lenRow-1])

			// Scenario A: Row has Description. [..., "TRUE", "Desc"]
			// Scenario B: Row has no Description (not targeted). [..., "FALSE"] or [..., ""]

			if strings.EqualFold(isFeatured, "TRUE") {
				description = lastCol
			} else if strings.EqualFold(lastCol, "TRUE") {
				// Old format or missing description column for this row (unlikely if we just added it to header but CSV reader handles ragged lines?)
				// Actually CSV reader with FieldsPerRecord=-1 allows variable length.
				isFeatured = "TRUE"
				// Description remains empty
			} else {
				// Not featured
				continue
			}

			if strings.EqualFold(isFeatured, "TRUE") {
				// Construct a basic SpotlightPet
				id := row[0]
				name := row[2]

				// Create JSON for description
				// escaped description to be safe? It's just simple text.
				// We need to marshal it to be safe JSON string value.
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

	// This query looks for the boolean column we updated
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
		// Ping to check connection before query (optional, but safe)
		if err := m.DB.Ping(); err == nil {
			query := `
				SELECT count(*)
				FROM pets
				WHERE status = 'adopted'
				AND (adoption->>'date') LIKE $1
			`
			yearPrefix := fmt.Sprintf("%d%%", year)
			var count int

			ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
			defer cancel()

			err := m.DB.QueryRowContext(ctx, query, yearPrefix).Scan(&count)
			if err == nil {
				return count, nil
			}
			fmt.Println("DB Query failed, falling back to CSV:", err)
		}
	}

	// 2. Fallback to CSV
	// Assume CSV is in the root backend folder or current working dir
	// We are running from backend/ so it should be "Master Pet List - Cats Master List.csv"
	// Or try absolute path if needed, but relative is best.
	fileName := "Master Pet List - Cats Master List.csv"

	// If running from root, path might be backend/...
	// Let's try opening both locations
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
		if i == 0 || len(row) < 39 { // Ensure enough columns
			continue
		}
		// status is index 36, adopted? check user diff again.
		// User diff: "adopted" is at index 36.
		// Date is at index 38. "1/3/2026"

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
		// Just return empty for CSV mode (simplification)
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
	Species    string  `json:"species"` // Hardcoded to 'cat' for now as DB col missing
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
			-- Inject 'status' from the column into the details JSON
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
			litter_name = $16
		WHERE id = $15
	`

	// Extract status from Details JSON if present, otherwise keep existing or default?
	// The frontend sends status inside details.
	// We need to parse p.Details to get the status for the column.
	var detailsMap map[string]interface{}
	status := "available" // Default fall back? Or should we fetch existing?
	// Better to try to parse.
	if len(p.Details) > 0 {
		if err := json.Unmarshal(p.Details, &detailsMap); err == nil {
			if s, ok := detailsMap["status"].(string); ok {
				status = s
			}
		}
	}

	// We should also remove 'status' from details JSON before saving to avoid duplication/confusion?
	// Or kept it in sync?
	// The GetAll logic injects it FROM the column.
	// So if we save it to JSON, it gets overwritten by column on read.
	// So it doesn't matter much if it's in JSON, but cleaner if we rely on column.
	// Let's just save the column.

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	// Handle standard fields arguments + JSONB arguments + ID
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
	}

	_, err := m.DB.ExecContext(ctx, query, args...)
	if err != nil {
		return err
	}

	return nil
}
