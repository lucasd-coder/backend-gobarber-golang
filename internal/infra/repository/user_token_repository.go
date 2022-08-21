package repository

import (
	"github.com/lucasd-coder/backend-gobarber-golang/internal/model"
	"github.com/lucasd-coder/backend-gobarber-golang/pkg/logger"

	"gorm.io/gorm"
)

type UserTokenRepository struct {
	Connection *gorm.DB
}

func NewUserTokenRepository(connectionDb *gorm.DB) *UserTokenRepository {
	return &UserTokenRepository{
		Connection: connectionDb,
	}
}

func (db *UserTokenRepository) Generate(userToken *model.UserToken) *model.UserToken {
	if db.Connection.Find(&userToken, "user_id = ?", userToken.UserID).Updates(&userToken).RowsAffected == 0 {
		err := db.Connection.Create(&userToken).Error
		if err != nil {
			logger.Log.Error(err.Error())
			return &model.UserToken{}
		}
	}

	return userToken
}

func (db *UserTokenRepository) FindByToken(token string) *model.UserToken {
	var userToken model.UserToken
	db.Connection.Model(model.UserToken{}).Where("token = ?", token).Find(&userToken)
	return &userToken
}
