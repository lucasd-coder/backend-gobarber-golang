package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Appointment struct {
	ID         uuid.UUID `json:"id" gorm:"primaryKey;type:uuid"`
	User       User      `json:"-" json:"user_id" binding:"required" gorm:"OnUpdate:CASCADE,OnDelete:CASCADE;foreignkey:UserID"`
	UserID     string    `json:"-"`
	Provider   User      `json:"-" json:"provider_id" binding:"required" gorm:"OnUpdate:CASCADE,OnDelete:CASCADE;foreignkey:ProviderID"`
	ProviderID string    `json:"-"`
	Date       time.Time `json:"date"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

func (appointment *Appointment) BeforeCreate(tx *gorm.DB) error {
	uuid := uuid.NewString()
	tx.Statement.SetColumn("ID", uuid)
	return nil
}

func NewAppointment(userId, providerId string, date time.Time) *Appointment {
	return &Appointment{
		UserID:     userId,
		ProviderID: providerId,
		Date:       date,
	}
}
