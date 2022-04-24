package model

import (
	"time"

	"github.com/google/uuid"
)

type Notification struct {
	ID          uuid.UUID `json:"id,omitempty" bson:"_id"`
	Content     string    `json:"content" bson:"content,omitempty"`
	RecipientID uuid.UUID `json:"recipient_id" bson:"recipient_id,omitempty"`
	Read        bool      `json:"read" bson:"read" default:"false"`
	CreatedAt   time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" bson:"updated_at"`
}
