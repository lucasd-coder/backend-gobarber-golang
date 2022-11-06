package model_test

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/lucasd-coder/backend-gobarber-golang/internal/model"
	"github.com/stretchr/testify/require"
)

func TestModel_NewAppointment(t *testing.T) {
	userID := uuid.NewString()
	providerID := uuid.NewString()
	date := time.Now()

	appointment := model.NewAppointment(userID, providerID, date)

	require.Equal(t, appointment.ProviderID, providerID)
	require.Equal(t, appointment.Date, date)
	require.Equal(t, appointment.UserID, userID)
}
