package database

import (
	"log"
	"tender_bid_system/model"
)

// SeedDatabase seeds the database with initial data
func SeedDatabase() {
	// Example: Seeding users
	users := []model.User{
		{Username: "admin", Email: "admin@example.com", Password: "hashedpassword1"},
		{Username: "contractor", Email: "contractor@example.com", Password: "hashedpassword2"},
	}

	for _, user := range users {
		if err := DB.Create(&user).Error; err != nil {
			log.Printf("Failed to seed user: %v", err)
		}
	}

	// Example: Seeding tenders
	tenders := []model.Tender{
		{Title: "Road Construction", Description: "Construction of a new highway", Status: "Open"},
		{Title: "Bridge Repair", Description: "Repairing the Golden Bridge", Status: "Open"},
	}

	for _, tender := range tenders {
		if err := DB.Create(&tender).Error; err != nil {
			log.Printf("Failed to seed tender: %v", err)
		}
	}

	log.Println("Database seeding completed.")
}
