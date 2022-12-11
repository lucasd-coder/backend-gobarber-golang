package service_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/lucasd-coder/backend-gobarber-golang/internal/dtos"
	test "github.com/lucasd-coder/backend-gobarber-golang/internal/mock"
	"github.com/lucasd-coder/backend-gobarber-golang/internal/model"
	"github.com/lucasd-coder/backend-gobarber-golang/internal/service"

	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
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

func TestProviderAppointmentsService_FindAllInDayFromProviderListAppointments(t *testing.T) {
	mockAppointmentRepository := new(test.MockAppointmentsRepository)
	mockCacheProvider := new(test.MockCacheProvider)

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
	cacheKey := fmt.Sprintf("provider-appointments:%s:%d-%d-%d", dto.ProviderID, dto.Year, dto.Month, dto.Day)

	mockCacheProvider.On("Recover", context.TODO(), cacheKey).Return(nil, redis.Nil)
	mockAppointmentRepository.On("FindAllInDayFromProvider", &dto).Return(appointments)
	mockCacheProvider.On("Save", context.TODO(), cacheKey, appointments, 4*time.Hour).Return(nil)

	testService := service.ListProviderAppointmentsService{mockAppointmentRepository, mockCacheProvider}

	resp, err := testService.Execute(&dto)

	assert.Nil(t, err)
	assert.Equal(t, appointments[0].ProviderID, resp[0].ProviderID)
	assert.Equal(t, appointments[0].UserID, resp[0].UserID)
	assert.Equal(t, appointments[0].Date, resp[0].Date)
}
