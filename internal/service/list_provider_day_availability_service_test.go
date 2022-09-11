package service_test

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/lucasd-coder/backend-gobarber-golang/internal/dtos"
	"github.com/lucasd-coder/backend-gobarber-golang/internal/model"
	"github.com/lucasd-coder/backend-gobarber-golang/internal/service"
	"github.com/lucasd-coder/backend-gobarber-golang/internal/test"
	"github.com/stretchr/testify/assert"
)

func TestProviderDayAvailabilityService_InvalidID(t *testing.T) {
	mockAppointmentRepository := new(test.MockAppointmentsRepository)

	var providerId string = " "

	dto := dtos.FindAllInDayFromProviderDTO{ProviderID: providerId, Day: 11, Month: 3, Year: 2025}

	testeService := service.ListProviderDayAvailabilityService{mockAppointmentRepository}

	_, err := testeService.Execute(&dto)

	assert.NotNil(t, err)
	assert.Equal(t, "provider_id invalid.", err.Error())
}

func TestProviderDayAvailabilityService_FindAllInDayFromProviderListAppointments(t *testing.T) {
	mockAppointmentRepository := new(test.MockAppointmentsRepository)

	var userId string = uuid.NewString()
	var id string = uuid.NewString()

	timeStamp := time.Date(2025, time.April,
		11, 14, 34, 0, 0, time.UTC)

	var providerId string = "0399e631-e2f0-4df5-b1d0-ca6d567a318c"

	appointments := []*model.Appointment{
		{
			ID:         uuid.MustParse(id),
			UserID:     userId,
			ProviderID: providerId,
			Date:       timeStamp,
			CreatedAt:  time.Now(),
			UpdatedAt:  time.Now(),
		},
	}

	dto := dtos.FindAllInDayFromProviderDTO{ProviderID: providerId, Day: 11, Month: 3, Year: 2025}

	mockAppointmentRepository.On("FindAllInDayFromProvider", &dto).Return(appointments)

	testService := service.ListProviderDayAvailabilityService{mockAppointmentRepository}

	resp, err := testService.Execute(&dto)

	assert.Nil(t, err)
	assert.NotEmpty(t, resp)
}
