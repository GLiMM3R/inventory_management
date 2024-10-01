package database

import (
	"inverntory_management/config"
	"inverntory_management/internal/database/schema"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func InitPostgres() {
	var err error

	// Database connection string
	dsn := config.AppConfig.DATABASE_URL

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			ParameterizedQueries:      false,       // Don't include params in the SQL log
			Colorful:                  false,       // Disable color
		},
	)

	// Initialize GORM with PostgreSQL
	DB, err = gorm.Open(postgres.Open(dsn),
		&gorm.Config{
			TranslateError: true, SkipDefaultTransaction: true,
			Logger:      newLogger,
			PrepareStmt: true})
	if err != nil {
		panic("Failed to connect database")
	}

	if err := DB.AutoMigrate(
		&schema.User{},
		&schema.Branch{},
		&schema.Category{},
		&schema.Inventory{},
		&schema.Product{},
		&schema.Variant{},
		&schema.Price{},
		&schema.PriceHistory{},
		&schema.Sale{},
		&schema.InventoryTransfer{},
	); err != nil {
		panic("Failed to migrate database")
	}
}
