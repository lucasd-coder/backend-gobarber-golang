package dtos

import (
	"time"
)

type ResponseProfileDTO struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Avatar    string    `json:"avatar"`
	CreatedAt time.Time `json:"-" json:"created_at"`
	UpdatedAt time.Time `json:"-" json:"updated_at"`
}

func NewResponseProfileDTO(id string, name string, email string, avatar string, createdAt time.Time, updatedAt time.Time) *ResponseProfileDTO {
	return &ResponseProfileDTO{
		ID:        id,
		Name:      name,
		Email:     email,
		Avatar:    avatar,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}
}
