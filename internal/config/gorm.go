package config

import (
	"log"

	"github.com/andruwizz/fiber-rest-minimal/internal/model"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func NewDatabase() *gorm.DB {
	// Initialize and return a new GORM database connection
	db, err := gorm.Open(sqlite.Open("database.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect database: %v", err)
	}

	_, err = db.DB()
	if err != nil {
		log.Fatalf("Failed to get database instance: %v", err)
	}

	db.AutoMigrate(&model.Todo{})

	return db
}
