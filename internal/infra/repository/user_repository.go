package repository

import (
	"github.com/lucasd-coder/backend-gobarber-golang/internal/model"
	"github.com/lucasd-coder/backend-gobarber-golang/pkg/logger"

	"gorm.io/gorm"
)

type UserRepository struct {
	Connection *gorm.DB
}

func NewUserRepository(connectionDb *gorm.DB) *UserRepository {
	return &UserRepository{
		Connection: connectionDb,
	}
}

func (db *UserRepository) Save(user *model.User) {
	err := db.Connection.Create(&user).Error
	if err != nil {
		logger.Log.Error(err.Error())
		return
	}
}

func (db *UserRepository) Update(user *model.User) {
	err := db.Connection.Save(&user).Error
	if err != nil {
		logger.Log.Error(err.Error())
		return
	}
}

func (db *UserRepository) FindByEmail(email string) *model.User {
	var user model.User
	db.Connection.Model(model.User{}).Where("email = ?", email).Find(&user)
	return &user
}

func (db *UserRepository) FindById(id string) *model.User {
	var user model.User
	db.Connection.Model(model.User{}).Where("id = ?", id).Find(&user)
	return &user
}
