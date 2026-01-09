package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
)

func main() {
	dsn := "postgres://pet_admin:Onyx57.@192.168.12.102:5432/shelter_db?sslmode=disable"
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Update 'realadmin' to 'super_admin'
	// Note: Adjust the email if you are using a different one
	res1, err := db.ExecContext(ctx, "UPDATE users SET role = 'super_admin' WHERE email = 'realadmin@example.com'")
	if err != nil {
		log.Fatalf("Failed to update super_admin: %v", err)
	}
	rows1, _ := res1.RowsAffected()
	fmt.Printf("Updated %d user(s) to super_admin (realadmin@example.com)\n", rows1)

	// Update 'admin@idohr.org' (Director) to 'admin'
	// Note: Assuming 'admin@idohr.org' exists. If not, this does nothing safe.
	res2, err := db.ExecContext(ctx, "UPDATE users SET role = 'admin' WHERE email = 'admin@idohr.org'")
	if err != nil {
		log.Fatalf("Failed to update admin: %v", err)
	}
	rows2, _ := res2.RowsAffected()
	fmt.Printf("Updated %d user(s) to admin (Director)\n", rows2)

	// Optional: Print all users and roles to verify
	rows, err := db.QueryContext(ctx, "SELECT email, role FROM users")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	fmt.Println("\n--- Current User Roles ---")
	for rows.Next() {
		var email, role string
		if err := rows.Scan(&email, &role); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%-25s | %s\n", email, role)
	}
	fmt.Println("--------------------------")
}
