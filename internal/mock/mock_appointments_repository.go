package mock

import (
	"time"

	"github.com/lucasd-coder/backend-gobarber-golang/internal/dtos"
	"github.com/lucasd-coder/backend-gobarber-golang/internal/model"

	"github.com/stretchr/testify/mock"
)

type MockAppointmentsRepository struct {
	mock.Mock
}

func (mock *MockAppointmentsRepository) Save(appointment *model.Appointment) {
}

func (mock *MockAppointmentsRepository) FindByDate(date *time.Time, providerId string) *model.Appointment {
	args := mock.Called(date, providerId)
	result := args.Get(0)
	return result.(*model.Appointment)
}

func (mock *MockAppointmentsRepository) FindAllInMonthFromProvider(data *dtos.FindAllInMonthFromProviderDTO) []*model.Appointment {
	args := mock.Called(data)
	result := args.Get(0)
	return result.([]*model.Appointment)
}

func (mock *MockAppointmentsRepository) FindAllInDayFromProvider(data *dtos.FindAllInDayFromProviderDTO) []*model.Appointment {
	args := mock.Called(data)
	result := args.Get(0)
	return result.([]*model.Appointment)
}
