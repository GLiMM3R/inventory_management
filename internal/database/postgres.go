package database

import (
	"inverntory_management/config"
	"inverntory_management/internal/database/schema"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitPostgres() {
	var err error

	// Database connection string
	dsn := config.AppConfig.DATABASE_URL

	// Initialize GORM with PostgreSQL
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{TranslateError: true})
	if err != nil {
		panic("Failed to connect database")
	}

	if err := DB.AutoMigrate(&schema.User{}, &schema.Branch{}, &schema.Inventory{}); err != nil {
		panic("Failed to migrate database")
	}
}
