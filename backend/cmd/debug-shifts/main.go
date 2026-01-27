package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	// Hardcoded DSN from previous step output
	dsn := "postgres://pet_admin:Onyx57.@192.168.12.102:5432/shelter_db?sslmode=disable"
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT id, volunteer_id, date, start_time, role, status, version FROM shifts ORDER BY id DESC LIMIT 5")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	fmt.Println("--- Recent Shifts ---")
	for rows.Next() {
		var id, volID int64
		var date, start, role, status string
		var version int
		if err := rows.Scan(&id, &volID, &date, &start, &role, &status, &version); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("ID: %d, VolID: %d, Date: %s, Start: %s, Role: %s, Status: %s, Version: %d\n", id, volID, date, start, role, status, version)
	}
	fmt.Println("---------------------")
}
