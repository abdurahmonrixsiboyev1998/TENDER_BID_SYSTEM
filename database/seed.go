package database

import (
	"log"
)

func SeedDatabase() {
	// Seed users
	users := []struct {
		Username string
		Email    string
		Password string
	}{
		{Username: "admin", Email: "admin@example.com", Password: "hashedpassword1"},
		{Username: "contractor", Email: "contractor@example.com", Password: "hashedpassword2"},
	}

	for _, user := range users {
		_, err := DB.Exec(
			"INSERT INTO users (username, email, password) VALUES ($1, $2, $3)",
			user.Username, user.Email, user.Password,
		)
		if err != nil {
			log.Printf("Failed to seed user: %v", err)
		}
	}

	tenders := []struct {
		Title       string
		Description string
		Status      string
	}{
		{Title: "Road Construction", Description: "Construction of a new highway", Status: "Open"},
		{Title: "Bridge Repair", Description: "Repairing the Golden Bridge", Status: "Open"},
	}

	for _, tender := range tenders {
		_, err := DB.Exec(
			"INSERT INTO tenders (title, description, status) VALUES ($1, $2, $3)",
			tender.Title, tender.Description, tender.Status,
		)
		if err != nil {
			log.Printf("Failed to seed tender: %v", err)
		}
	}

	log.Println("Database seeding completed.")
}
