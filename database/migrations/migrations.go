package migrations

import (
	"go-rest-api/models"

	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) {
	db.AutoMigrate(models.User{})
	db.AutoMigrate(models.Order{})
}
