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
