package database

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq" // PostgreSQL driver
)

var DB *sql.DB

func ConnectDatabase() {
	// dsn := fmt.Sprintf(
	// 	"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
	// 	os.Getenv("DB_HOST"),
	// 	os.Getenv("DB_USER"),
	// 	os.Getenv("DB_PASSWORD"),
	// 	os.Getenv("DB_NAME"),
	// 	os.Getenv("DB_PORT"),
	// )

	var err error
	DB, err = sql.Open("postgres", "postgres://postgres:14022014@localhost:5432/tender_bid_system?sslmode=disable")
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Ping to verify connection
	if err = DB.Ping(); err != nil {
		log.Fatalf("Failed to ping database: %v", err)
	}

	log.Println("Database connection established.")
}
