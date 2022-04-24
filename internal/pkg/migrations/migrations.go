package migrations

import (
	"backend-gobarber-golang/internal/model"

	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) {
	db.AutoMigrate(model.User{}, model.UserToken{}, model.Appointment{})
}
