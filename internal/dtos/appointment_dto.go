package dtos

import "time"

type AppointmentDTO struct {
	Date       time.Time `json:"date" binding:"required"`
	ProviderID string    `json:"provider_id" binding:"required"`
}
