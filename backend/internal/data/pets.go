package data // <--- Must be 'data'

import (
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

	rows, err := m.DB.Query(query)
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
			err := m.DB.QueryRow(query, yearPrefix).Scan(&count)
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

	rows, err := m.DB.Query(query)
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
