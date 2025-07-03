package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

// Config struct holds application-wide configurations.
type Config struct {
	DatabaseURL string
}

// LoadConfig loads configuration from environment variables or .env file.
func LoadConfig() (*Config, error) {
	// Load .env file. This will not overwrite existing environment variables.
	err := godotenv.Load()
	if err != nil {
		fmt.Println("No .env file found, loading from environment variables.")
	}

	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		return nil, fmt.Errorf("DATABASE_URL environment variable not set")
	}

	return &Config{
		DatabaseURL: dbURL,
	}, nil
}
