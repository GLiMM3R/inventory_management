package database

import (
	"fmt"
	"inverntory_management/config"
	"inverntory_management/internal/feature/branch"
	"inverntory_management/internal/feature/user"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitPostgres() {
	var err error

	// Database connection string
	dsn := config.AppConfig.DATABASE_URL
	fmt.Println(dsn)
	// Initialize GORM with PostgreSQL
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{TranslateError: true})
	if err != nil {
		panic("Failed to connect database")
	}

	if err := DB.AutoMigrate(&user.User{}, &branch.Branch{}); err != nil {
		panic("Failed to migrate database")
	}
}
