package database

import (
	"fmt"
	"log"

	"github.com/Waleed978/multi_tenent_backend/config"
	"github.com/Waleed978/multi_tenent_backend/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// DB is the global database connection instance.
var DB *gorm.DB

// InitDB initializes the database connection and performs auto-migration.
func InitDB(cfg *config.Config) {
	var err error
	dsn := cfg.DatabaseURL
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	fmt.Println("Database connection established successfully!")

	// Auto-migrate the Student model. GORM will create the table if it doesn't exist.
	// It will also add/update columns based on the struct definition.
	err = DB.AutoMigrate(&models.Student{})
	if err != nil {
		log.Fatalf("Failed to auto-migrate database: %v", err)
	}
	fmt.Println("Database migration completed successfully!")
}
