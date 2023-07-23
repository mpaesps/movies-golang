package db

import (
	"movies-crud/pkg/models"

	"github.com/jinzhu/gorm"
)

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&models.Movie{})
}
