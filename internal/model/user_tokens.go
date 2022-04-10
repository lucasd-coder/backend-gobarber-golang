package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserToken struct {
	ID        uuid.UUID `json:"id" gorm:"primaryKey;type:uuid"`
	Token     uuid.UUID `json:"token" gorm:"type:uuid"`
	UserID    string    `json:"-"`
	User      User      `json:"-" json:"user_id" binding:"required" gorm:"OnUpdate:CASCADE,OnDelete:CASCADE;foreignkey:UserID"`
	CreatedAt time.Time `json:"-" json:"created_at"`
	UpdatedAt time.Time `json:"-" json:"updated_at"`
}

func (userToken *UserToken) BeforeCreate(tx *gorm.DB) error {
	uuid := uuid.NewString()
	tx.Statement.SetColumn("ID", uuid)
	tx.Statement.SetColumn("Token", uuid)
	return nil
}

func NewUserToken(userId string) *UserToken {
	return &UserToken{
		UserID: userId,
	}
}
