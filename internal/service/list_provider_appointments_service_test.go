package service_test

import (
	"context"
	"testing"
	"time"

	"backend-gobarber-golang/internal/dtos"
	"backend-gobarber-golang/internal/model"
	"backend-gobarber-golang/internal/service"
	"backend-gobarber-golang/internal/test"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestProviderAppointmentsService_InvalidID(t *testing.T) {
	mockAppointmentRepository := new(test.MockAppointmentsRepository)
	mockCacheProvider := new(test.MockCacheProvider)

	var providerId string = " "

	dto := dtos.FindAllInDayFromProviderDTO{ProviderID: providerId, Day: 11, Month: 3, Year: 2025}

	testService := service.ListProviderAppointmentsService{mockAppointmentRepository, mockCacheProvider}

	_, err := testService.Execute(&dto)

	assert.NotNil(t, err)
	assert.Equal(t, "provider_id invalid.", err.Error())
}

func TestListProviderAppointmentsService_RecoverReturnListAppointments(t *testing.T) {
	mockAppointmentRepository := new(test.MockAppointmentsRepository)
	mockCacheProvider := new(test.MockCacheProvider)

	var userId string = uuid.NewString()
	var id string = uuid.NewString()

	timeStamp := time.Date(2025, time.April,
		11, 14, 34, 0, 0, time.UTC)

	var providerId string = "0399e631-e2f0-4df5-b1d0-ca6d567a318c"

	appointments := []model.Appointment{
		{
			ID:         uuid.MustParse(id),
			UserID:     userId,
			ProviderID: providerId,
			Date:       timeStamp,
			CreatedAt:  time.Now(),
			UpdatedAt:  time.Now(),
		},
	}

	mockCacheProvider.On("Recover", context.TODO(), mock.Anything).Return(appointments)

	dto := dtos.FindAllInDayFromProviderDTO{ProviderID: providerId, Day: 11, Month: 3, Year: 2025}

	testService := service.ListProviderAppointmentsService{mockAppointmentRepository, mockCacheProvider}

	resp, err := testService.Execute(&dto)

	assert.Nil(t, err)
	assert.Equal(t, appointments[0].ProviderID, resp[0].ProviderID)
	assert.Equal(t, appointments[0].UserID, resp[0].UserID)
	assert.Equal(t, appointments[0].Date, resp[0].Date)
}

func TestProviderAppointmentsService_FindAllInDayFromProviderListAppointments(t *testing.T) {
	mockAppointmentRepository := new(test.MockAppointmentsRepository)
	mockCacheProvider := new(test.MockCacheProvider)

	var userId string = uuid.NewString()
	var id string = uuid.NewString()

	timeStamp := time.Date(2025, time.April,
		11, 14, 34, 0, 0, time.UTC)

	var providerId string = "0399e631-e2f0-4df5-b1d0-ca6d567a318c"

	appointments := []model.Appointment{
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

	mockCacheProvider.On("Recover", context.TODO(), mock.Anything).Return([]model.Appointment{})
	mockAppointmentRepository.On("FindAllInDayFromProvider", &dto).Return(appointments)
	mockCacheProvider.On("Save", context.TODO(), mock.Anything, appointments).Return(nil)

	testService := service.ListProviderAppointmentsService{mockAppointmentRepository, mockCacheProvider}

	resp, err := testService.Execute(&dto)

	assert.Nil(t, err)
	assert.Equal(t, appointments[0].ProviderID, resp[0].ProviderID)
	assert.Equal(t, appointments[0].UserID, resp[0].UserID)
	assert.Equal(t, appointments[0].Date, resp[0].Date)
}
