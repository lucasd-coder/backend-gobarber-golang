package model

import (
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Notification struct {
	ID          primitive.ObjectID `json:"id,omitempty" bson:"_id"`
	Content     string             `json:"content" bson:"content,omitempty"`
	RecipientID uuid.UUID          `json:"recipient_id" bson:"recipient_id,omitempty"`
	Read        bool               `json:"read" bson:"read" default:"false"`
	CreatedAt   time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt   time.Time          `json:"updated_at" bson:"updated_at"`
}

func NewNotification(recipient_id uuid.UUID, content string) *Notification {
	return &Notification{
		ID:          primitive.NewObjectID(),
		RecipientID: recipient_id,
		Content:     content,
		CreatedAt:   time.Now(),
	}
}
