package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Config struct holds all configuration variables
type Config struct {
	DBHost     string
	DBUser     string
	DBPassword string
	DBName     string
	DBPort     string
	AppPort    string
	JWTSecret  string
}

// LoadConfig loads configuration from .env file or environment variables
func LoadConfig() *Config {
	// Load environment variables from .env file
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, relying on system environment variables.")
	}

	// Return the configuration
	return &Config{
		DBHost:     getEnv("DB_HOST", "localhost"),
		DBUser:     getEnv("DB_USER", "postgres"),
		DBPassword: getEnv("DB_PASSWORD", "14022014"),
		DBName:     getEnv("DB_NAME", "tender_bid_system"),
		DBPort:     getEnv("DB_PORT", "5432"),
		AppPort:    getEnv("APP_PORT", "8080"),
		JWTSecret:  getEnv("JWT_SECRET", "supersecretkey"),
	}
}

// getEnv fetches environment variable or uses default value if not set
func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
