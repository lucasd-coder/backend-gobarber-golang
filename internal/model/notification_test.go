package model_test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/lucasd-coder/backend-gobarber-golang/internal/model"
	"github.com/stretchr/testify/require"
)

func TestModel_NewNotification(t *testing.T) {
	recipientID := uuid.New()
	content := "Hello world"

	notification := model.NewNotification(recipientID, content)

	require.Equal(t, notification.Content, content)
	require.NotEmpty(t, notification.ID)
	require.Equal(t, notification.RecipientID, recipientID)
}
