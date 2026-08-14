package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	dsn := "postgres://root:password@localhost:5432/jobshare?sslmode=disable"
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("Failed to connect to db: %v", err)
	}
	defer db.Close()

	email := "superadmin@jobshare.com"
	password := "AdminSecret123!"
	role := "admin"

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatalf("Failed to hash password: %v", err)
	}

	query := `
		INSERT INTO users (email, password_hash, role) 
		VALUES ($1, $2, $3) 
		ON CONFLICT (email) DO NOTHING
		RETURNING id;
	`

	var id string

	err = db.QueryRow(query, email, string(hashedPassword), role).Scan(&id)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("Admin account already exists!")
			return
		}
		log.Fatalf("Failed to insert admin: %v", err)
	}

	fmt.Printf("Superadmin created successfully! ID: %d | Email: %s\n", id, email)
}
