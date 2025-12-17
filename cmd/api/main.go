package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	// Import your new models package!
	// REPLACE "github.com/YOUR_USER/pet-adoption-backend" with your actual module name
	// You can find the module name in your go.mod file.
	"idohr-be/internal/models"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/rs/cors"
)

func main() {
	_ = godotenv.Load()

	dbHost := os.Getenv("DB_HOST")
	if dbHost == "" {
		dbHost = "localhost"
	}
	dbPass := os.Getenv("DB_PASSWORD")

	connStr := fmt.Sprintf("postgres://pet_admin:%s@%s:5432/shelter_db?sslmode=disable", dbPass, dbHost)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal("Could not connect to DB: ", err)
	}
	fmt.Printf("Connected to Database on %s!\n", dbHost)

	mux := http.NewServeMux()

	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Pet Adoption API is Running! 🐶"))
	})

	mux.HandleFunc("GET /pets", func(w http.ResponseWriter, r *http.Request) {
		getAllPets(w, r, db)
	})

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type"},
		AllowCredentials: true,
	})

	port := "8080"
	fmt.Println("Server starting on port " + port + "...")
	log.Fatal(http.ListenAndServe(":"+port, c.Handler(mux)))
}

func getAllPets(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	// We are now fetching more data using the shared structs
	rows, err := db.Query("SELECT id, name, species, breed, status FROM pets LIMIT 50")
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	defer rows.Close()

	var pets []models.Pet
	for rows.Next() {
		var p models.Pet
		// We only scan the fields we asked for in the SQL query
		// The other fields in the struct will just stay empty/default for this specific endpoint
		if err := rows.Scan(&p.ID, &p.Name, &p.Species, &p.Breed, &p.Status); err != nil {
			log.Println("Error scanning row:", err)
			continue
		}
		pets = append(pets, p)
	}

	if pets == nil {
		pets = []models.Pet{}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(pets)
}
