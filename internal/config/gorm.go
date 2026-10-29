package config

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func NewDatabase() *gorm.DB {
	// Initialize and return a new GORM database connection
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect database: %v", err)
	}

	_, err = db.DB()
	if err != nil {
		log.Fatalf("Failed to get database instance: %v", err)
	}

	return db
}
