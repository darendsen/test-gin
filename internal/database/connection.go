package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/darendsen/test-gin/internal/models"
)

func Connect() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// Migrate the schema
	db.AutoMigrate(&models.User{})

	return db, nil
}
