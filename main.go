package main

import (
	"database/sql"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	_ "github.com/lib/pq"
)

// --- Structs (Same as before) ---
type BehaviorData struct {
	EnergyLevel     string   `json:"energyLevel"`
	IsHouseTrained  bool     `json:"isHouseTrained"`
	IsGoodWithKids  bool     `json:"isGoodWithKids"`
	IsGoodWithCats  bool     `json:"isGoodWithCats"`
	IsGoodWithDogs  bool     `json:"isGoodWithDogs"`
	PrefersAlone    bool     `json:"prefersToBeAlone"`
	SpecialNeeds    string   `json:"specialNeeds"`
	PersonalityTags []string `json:"personalityTags"`
	BondedWith      []string `json:"bondedWith"`
	IsBonded        bool     `json:"isBonded"`
}

type MedicalData struct {
	VaccinationsUpToDate bool     `json:"vaccinationsUpToDate"`
	SpayedOrNeutered     bool     `json:"spayedOrNeutered"`
	Microchipped         bool     `json:"microchipped"`
	MicrochipID          string   `json:"microchipID"`
	MicrochipCompany     string   `json:"microchipCompany"`
	HealthConcerns       []string `json:"healthConcerns"`
	CurrentMedications   []string `json:"currentMedications"`
	HealthSummary        string   `json:"healthSummary"`
}

type DescriptionData struct {
	Primary        string   `json:"primary"`
	Spotlight      string   `json:"spotlight"`
	Behavioral     string   `json:"behavioral"`
	Origin         string   `json:"origin"`
	FunFact        string   `json:"fun"`
	AdditionalInfo []string `json:"additionalInformation"`
}

type AdoptionData struct {
	IsAdopted       bool    `json:"isAdopted"`
	Date            *string `json:"date"`
	Price           float64 `json:"price"`
	NewAdoptedName  string  `json:"newAdoptedName"`
	AdoptedBy       string  `json:"adoptedBy"`
	SurveyCompleted bool    `json:"surveyCompleted"`
}

type FosterData struct {
	BeingFostered bool    `json:"beingFostered"`
	StartDate     *string `json:"startDate"`
	EndDate       *string `json:"endDate"`
	ParentName    string  `json:"parentName"`
}

func main() {
	// --- CONFIGURATION ---
	// UPDATE THIS if running remotely (use Pi IP) or locally (use 'localhost')
	connStr := "postgres://pet_admin:secure_password_here@localhost:5432/shelter_db?sslmode=disable"

	fmt.Println("Connecting to database...")
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err = db.Ping(); err != nil {
		log.Fatal("Could not connect to DB. Ensure Docker is running.", err)
	}
	fmt.Println("Connected!")

	// 1. ENSURE TABLES EXIST
	createTables(db)

	// 2. READ CSV
	// Note: Make sure the file name matches exactly what you uploaded
	fileName := "Master Pet List - Cats Master List.csv"
	fmt.Printf("Reading CSV file: %s ...\n", fileName)

	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal("Could not open CSV file. Check filename and location:", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	// Allow variable number of fields if rows have trailing empty commas
	reader.FieldsPerRecord = -1
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	// 3. INSERT DATA
	fmt.Println("Seeding data...")
	count := 0

	// Default Species since this is the "Cats" list
	defaultSpecies := "cat"

	for i, row := range records {
		// SKIP HEADER ROWS: Row 0 is "VACCINATION DATES", Row 1 is "ID, NT-Date..."
		if i < 2 {
			continue
		}

		// Helpers
		get := func(idx int) string {
			if idx >= len(row) {
				return ""
			}
			return strings.TrimSpace(row[idx])
		}

		getBool := func(idx int) bool {
			val := strings.ToLower(get(idx))
			return val == "true" || val == "yes" || val == "1"
		}

		getDate := func(idx int) sql.NullTime {
			val := get(idx)
			// Handle cases like "----" or empty
			if val == "" || strings.Contains(val, "-") && len(val) < 6 {
				return sql.NullTime{}
			}
			// Supported formats: M/D/YYYY (Excel default) or YYYY-MM-DD
			formats := []string{"1/2/2006", "2006-01-02", "1-2-06"}

			for _, f := range formats {
				t, err := time.Parse(f, val)
				if err == nil {
					return sql.NullTime{Time: t, Valid: true}
				}
			}
			return sql.NullTime{}
		}

		// Parse Currency (Remove '$' and ',')
		parseCurrency := func(idx int) float64 {
			val := get(idx)
			val = strings.ReplaceAll(val, "$", "")
			val = strings.ReplaceAll(val, ",", "")
			if val == "" || val == "----" {
				return 0.0
			}
			f, _ := strconv.ParseFloat(val, 64)
			return f
		}

		// --- MAPPING NEW COLUMNS ---
		// Based on "Master Pet List - Cats Master List.csv"

		name := get(2)
		if name == "" {
			continue
		} // Skip empty rows

		behavior := BehaviorData{
			EnergyLevel:     get(17),
			IsHouseTrained:  getBool(21),
			IsGoodWithKids:  getBool(20),
			IsGoodWithCats:  getBool(18),
			IsGoodWithDogs:  getBool(19),
			PrefersAlone:    getBool(24),
			SpecialNeeds:    get(25),
			IsBonded:        getBool(15),
			BondedWith:      splitByComma(get(16)),
			PersonalityTags: splitByComma(get(23)),
		}

		medical := MedicalData{
			VaccinationsUpToDate: getBool(32),
			SpayedOrNeutered:     getBool(31),
			Microchipped:         getBool(12),
			MicrochipID:          get(13),
			MicrochipCompany:     get(14),
			HealthConcerns:       splitByComma(get(42)),
			CurrentMedications:   splitByComma(get(43)),
			HealthSummary:        get(9),
		}

		// Description columns are missing in this specific CSV, so leaving empty
		description := DescriptionData{}

		adoption := AdoptionData{
			IsAdopted:       getBool(39),
			Date:            nullableString(get(38)),
			Price:           parseCurrency(37),
			NewAdoptedName:  get(41),
			AdoptedBy:       get(40),
			SurveyCompleted: getBool(44),
		}

		// Foster data columns are missing in this CSV, defaulting to nulls
		foster := FosterData{}

		// Marshal JSON
		bJSON, _ := json.Marshal(behavior)
		mJSON, _ := json.Marshal(medical)
		dJSON, _ := json.Marshal(description)
		aJSON, _ := json.Marshal(adoption)
		fJSON, _ := json.Marshal(foster)

		// SQL Insert
		sqlStatement := `
			INSERT INTO pets (
				name, species, breed, litter_name, sex, status, 
				size, color, coat_length, date_of_birth,
				intake_date, shelter_location, 
				behavior_data, medical_data, description_data, adoption_data, foster_data,
				is_returned, is_sponsored
			)
			VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19)
			RETURNING id`

		var id string
		err = db.QueryRow(sqlStatement,
			name,
			defaultSpecies, // Hardcoded 'cat'
			get(5),         // Breed
			get(4),         // Litter
			get(10),        // Sex
			get(36),        // Status
			get(11),        // Size
			get(7),         // Color
			get(6),         // Coat Length
			getDate(3),     // DOB
			getDate(34),    // Intake Date
			get(35),        // Shelter Location
			bJSON, mJSON, dJSON, aJSON, fJSON,
			getBool(49), // Is Returned
			getBool(45), // Is Sponsored
		).Scan(&id)

		if err != nil {
			log.Printf("Error inserting %s: %v", name, err)
		} else {
			count++
		}
	}
	fmt.Printf("Done! Successfully seeded %d pets.\n", count)
}

func createTables(db *sql.DB) {
	fmt.Println("Ensuring tables exist...")
	_, err := db.Exec(`CREATE EXTENSION IF NOT EXISTS "uuid-ossp";`)
	if err != nil {
		log.Fatal("Error enabling UUID extension:", err)
	}

	petsTable := `
	CREATE TABLE IF NOT EXISTS pets (
		id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
		slug TEXT,
		created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
		name TEXT NOT NULL,
		species TEXT NOT NULL, 
		sex TEXT NOT NULL,     
		litter_name TEXT,
		breed TEXT,            
		age_group TEXT,        
		size TEXT,             
		color TEXT,
		coat_length TEXT,
		date_of_birth DATE,
		current_weight NUMERIC(5,2), 
		status TEXT NOT NULL,  
		intake_date DATE,
		shelter_location TEXT,
		behavior_data JSONB,   
		medical_data JSONB,    
		description_data JSONB, 
		adoption_data JSONB,   
		foster_data JSONB,     
		is_spotlight BOOLEAN DEFAULT FALSE,
		is_sponsored BOOLEAN DEFAULT FALSE,
		is_returned BOOLEAN DEFAULT FALSE
	);`
	_, err = db.Exec(petsTable)
	if err != nil {
		log.Fatal("Error creating pets table:", err)
	}

	logsTable := `
	CREATE TABLE IF NOT EXISTS health_logs (
		id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
		pet_id UUID REFERENCES pets(id) ON DELETE CASCADE,
		date TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
		recorded_by TEXT,
		weight NUMERIC(5,2),
		temperature NUMERIC(4,1),
		eating_status TEXT,   
		drinking_status TEXT,
		activity_log TEXT,
		urination TEXT,       
		defecation TEXT,
		notes TEXT
	);`
	_, err = db.Exec(logsTable)
	if err != nil {
		log.Fatal("Error creating health_logs table:", err)
	}
}

func splitByComma(s string) []string {
	if s == "" {
		return []string{}
	}
	parts := strings.Split(s, ",")
	for i := range parts {
		parts[i] = strings.TrimSpace(parts[i])
	}
	return parts
}

func nullableString(s string) *string {
	if s == "" {
		return nil
	}
	return &s
}
