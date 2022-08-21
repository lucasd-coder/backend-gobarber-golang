package service_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/lucasd-coder/backend-gobarber-golang/internal/dtos"
	"github.com/lucasd-coder/backend-gobarber-golang/internal/model"
	"github.com/lucasd-coder/backend-gobarber-golang/internal/service"
	"github.com/lucasd-coder/backend-gobarber-golang/internal/test"
	"github.com/lucasd-coder/backend-gobarber-golang/internal/util"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateAppointmentService_InvalidID(t *testing.T) {
	mockAppointmentRepository := new(test.MockAppointmentsRepository)
	mockNotificationsRepository := new(test.MockNotificationsRepository)
	mockCacheProvider := new(test.MockCacheProvider)

	testService := service.CreateAppointmentService{mockAppointmentRepository, mockNotificationsRepository, mockCacheProvider}

	timeStamp := time.Now().AddDate(0, 0, 1)

	var id string = "0399e631-e2f0-4df5-b1d0-ca6d567a318c"

	appointmentDto := dtos.AppointmentDTO{Date: timeStamp, ProviderID: " "}

	_, err := testService.Execute(id, &appointmentDto)

	assert.Equal(t, "provider_id invalid.", err.Error())
}

func TestCreateAppointmentService_NotCreateOnePostDate(t *testing.T) {
	mockAppointmentRepository := new(test.MockAppointmentsRepository)
	mockNotificationsRepository := new(test.MockNotificationsRepository)
	mockCacheProvider := new(test.MockCacheProvider)

	var id string = "0399e631-e2f0-4df5-b1d0-ca6d567a318c"

	timeStamp := time.Now()

	appointmentDto := dtos.AppointmentDTO{Date: timeStamp, ProviderID: id}

	testService := service.CreateAppointmentService{mockAppointmentRepository, mockNotificationsRepository, mockCacheProvider}

	_, err := testService.Execute(id, &appointmentDto)

	assert.Equal(t, "You can't create an appointment on a post date.", err.Error())
}

func TestCreateAppointmentService_NotAbleToBeCreateAnAppointmentWithSameUserAsProvider(t *testing.T) {
	mockAppointmentRepository := new(test.MockAppointmentsRepository)
	mockNotificationsRepository := new(test.MockNotificationsRepository)
	mockCacheProvider := new(test.MockCacheProvider)

	var id string = "0399e631-e2f0-4df5-b1d0-ca6d567a318c"

	timeStamp := time.Now().AddDate(0, 0, 1)

	appointmentDto := dtos.AppointmentDTO{ProviderID: id, Date: timeStamp}

	testService := service.CreateAppointmentService{mockAppointmentRepository, mockNotificationsRepository, mockCacheProvider}

	_, err := testService.Execute(id, &appointmentDto)

	assert.Equal(t, "You can't create an appointment with yourself.", err.Error())
}

func TestCreateAppointmentService_NotBeAbleToCreateAnAppointmentBefore8anAndAfter5pm(t *testing.T) {
	mockAppointmentRepository := new(test.MockAppointmentsRepository)
	mockNotificationsRepository := new(test.MockNotificationsRepository)
	mockCacheProvider := new(test.MockCacheProvider)

	var id string = uuid.NewString()

	appointmentDto := dtos.AppointmentDTO{Date: dateTest(), ProviderID: "0399e631-e2f0-4df5-b1d0-ca6d567a318c"}

	testService := service.CreateAppointmentService{mockAppointmentRepository, mockNotificationsRepository, mockCacheProvider}

	_, err := testService.Execute(id, &appointmentDto)

	assert.Equal(t, "You can only create appointments between 8am and 5pm", err.Error())
}

func TestCreateAppointmentService_NotBeAbleToCreateAnAppointmentWithSameUserAsProvider(t *testing.T) {
	mockAppointmentRepository := new(test.MockAppointmentsRepository)
	mockNotificationsRepository := new(test.MockNotificationsRepository)
	mockCacheProvider := new(test.MockCacheProvider)

	var userId string = uuid.NewString()
	var id string = uuid.NewString()

	var providerId string = "0399e631-e2f0-4df5-b1d0-ca6d567a318c"

	timeStamp := time.Date(2025, time.April,
		11, 14, 34, 0, 0, time.UTC)

	appointmentDto := dtos.AppointmentDTO{Date: timeStamp, ProviderID: providerId}

	appointment := model.Appointment{
		ID: uuid.MustParse(id), UserID: userId, ProviderID: providerId,
		Date: timeStamp, CreatedAt: time.Now(), UpdatedAt: time.Now(),
	}

	mockAppointmentRepository.On("FindByDate", &timeStamp, appointmentDto.ProviderID).Return(&appointment)

	testService := service.CreateAppointmentService{mockAppointmentRepository, mockNotificationsRepository, mockCacheProvider}

	_, err := testService.Execute(userId, &appointmentDto)

	assert.Equal(t, "This appointment is already backed", err.Error())
}

func TestCreateAppointmentService_ErrorSaveNotification(t *testing.T) {
	mockAppointmentRepository := new(test.MockAppointmentsRepository)
	mockNotificationsRepository := new(test.MockNotificationsRepository)
	mockCacheProvider := new(test.MockCacheProvider)

	var id string = uuid.NewString()

	providerId := "0399e631-e2f0-4df5-b1d0-ca6d567a318c"

	timeStamp := time.Date(2025, time.April,
		11, 14, 34, 0, 0, time.UTC)

	appointmentDto := dtos.AppointmentDTO{Date: timeStamp, ProviderID: providerId}

	mockAppointmentRepository.On("FindByDate", &timeStamp, appointmentDto.ProviderID).Return(&model.Appointment{})

	mockNotificationsRepository.On("Save", mock.Anything).Return(errors.New("error"))

	testService := service.CreateAppointmentService{mockAppointmentRepository, mockNotificationsRepository, mockCacheProvider}

	_, err := testService.Execute(id, &appointmentDto)

	assert.Equal(t, "Internal Server Error", err.Error())
}

func TestCreateAppointmentService_ReturnAppointment(t *testing.T) {
	mockAppointmentRepository := new(test.MockAppointmentsRepository)
	mockNotificationsRepository := new(test.MockNotificationsRepository)
	mockCacheProvider := new(test.MockCacheProvider)

	var id string = uuid.NewString()

	providerId := "0399e631-e2f0-4df5-b1d0-ca6d567a318c"

	timeStamp := time.Date(2025, time.April,
		11, 14, 34, 0, 0, time.UTC)

	appointmentDto := dtos.AppointmentDTO{Date: timeStamp, ProviderID: providerId}

	mockAppointmentRepository.On("FindByDate", &timeStamp, appointmentDto.ProviderID).Return(&model.Appointment{})

	mockNotificationsRepository.On("Save", mock.Anything).Return(nil)

	mockCacheProvider.On("Invalidate", context.TODO(), mock.Anything).Return(nil)

	testService := service.CreateAppointmentService{mockAppointmentRepository, mockNotificationsRepository, mockCacheProvider}

	resp, err := testService.Execute(id, &appointmentDto)

	assert.Nil(t, err)
	assert.Equal(t, providerId, resp.ProviderID)
	assert.Equal(t, timeStamp, resp.Date)
}

func dateTest() time.Time {
	appointmentDate := time.Now().Local()

	timeStamp, _ := util.DateUtils(appointmentDate, "2006-01-02 15:04:05")

	hr, _, _ := timeStamp.Clock()

	if hr > 8 || hr < 17 {
		return appointmentDate.Add(time.Hour * time.Duration(10))
	}

	return appointmentDate
}
