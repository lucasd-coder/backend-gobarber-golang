package service_test

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/lucasd-coder/backend-gobarber-golang/internal/dtos"
	test "github.com/lucasd-coder/backend-gobarber-golang/internal/mock"
	"github.com/lucasd-coder/backend-gobarber-golang/internal/model"
	"github.com/lucasd-coder/backend-gobarber-golang/internal/service"
	"github.com/stretchr/testify/assert"
)

func TestProviderMonthAvailabilityService_InvalidID(t *testing.T) {
	mockAppointmentRepository := new(test.MockAppointmentsRepository)

	var providerId string = " "

	dto := dtos.FindAllInMonthFromProviderDTO{ProviderID: providerId, Month: 3, Year: 2025}

	testService := service.ListProviderMonthAvailabilityService{mockAppointmentRepository}

	_, err := testService.Execute(&dto)

	assert.NotNil(t, err)
	assert.Equal(t, "provider_id invalid.", err.Error())
}

func TestProviderMonthAvailabilityService_FindAllInMonthFromProviderListAppointments(t *testing.T) {
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

	dto := dtos.FindAllInMonthFromProviderDTO{ProviderID: providerId, Month: 3, Year: 2025}

	mockAppointmentRepository.On("FindAllInMonthFromProvider", &dto).Return(appointments)

	testService := service.ListProviderMonthAvailabilityService{mockAppointmentRepository}

	resp, err := testService.Execute(&dto)

	assert.Nil(t, err)
	assert.NotEmpty(t, resp)
}
