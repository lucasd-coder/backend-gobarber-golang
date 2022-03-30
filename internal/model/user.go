package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID        uuid.UUID `json:"id" gorm:"primaryKey;type:uuid"`
	Name      string    `json:"name" binding:"required" gorm:"type:varchar(256)"`
	Email     string    `json:"email" binding:"required,email" gorm:"type:varchar(256);uniqueIndex"`
	Password  string    `json:"password" binding:"min=6,max=200" gorm:"type:varchar(256)"`
	Avatar    string    `json:"avatar" gorm:"type:varchar(256)"`
	CreatedAt time.Time `json:"-" json:"created_at"`
	UpdatedAt time.Time `json:"-" json:"updated_at"`
}

func (user *User) BeforeCreate(tx *gorm.DB) error {
	uuid := uuid.NewString()
	tx.Statement.SetColumn("ID", uuid)
	return nil
}

func NewUser(name string, email string, password string) *User {
	return &User{
		Name:     name,
		Email:    email,
		Password: password,
	}
}
