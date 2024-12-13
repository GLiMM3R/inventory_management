package migration

import (
	"inverntory_management/config"
	"inverntory_management/internal/database"
	"inverntory_management/internal/database/schema"

	"gorm.io/gorm"
)

type Migration struct {
	cfg config.Config
	db  *gorm.DB
}

func New(cfg config.Config) *Migration {
	return &Migration{
		cfg: cfg,
		db:  database.InitPostgres(cfg),
	}
}

func (m *Migration) Run() {
	// Perform database schema migration based on the configuration
	if err := m.db.AutoMigrate(
		&schema.User{},
		&schema.Category{},
		&schema.Product{},
		&schema.ProductVariant{},
		&schema.Attribute{},
		&schema.Sale{},
		&schema.Media{},
	); err != nil {
		panic("Failed to migrate database")
	}
}
