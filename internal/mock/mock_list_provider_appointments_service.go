package mock

import (
	"github.com/lucasd-coder/backend-gobarber-golang/internal/dtos"
	"github.com/lucasd-coder/backend-gobarber-golang/internal/model"
	"github.com/stretchr/testify/mock"
)

type MockListProviderAppointmentsService struct {
	mock.Mock
}

func NewMockListProviderAppointmentsService() *MockListProviderAppointmentsService {
	return &MockListProviderAppointmentsService{}
}

func (mock *MockListProviderAppointmentsService) Execute(dto *dtos.FindAllInDayFromProviderDTO) ([]*model.Appointment, error) {
	args := mock.Called(dto)

	r0 := []*model.Appointment{}

	if rf, ok := args.Get(0).(func(*dtos.FindAllInDayFromProviderDTO) []*model.Appointment); ok {
		r0 = rf(dto)
	} else {
		if args.Get(0) != nil {
			r0 = args.Get(0).([]*model.Appointment)
		}
	}

	var r1 error
	if rf, ok := args.Get(1).(func(*dtos.FindAllInDayFromProviderDTO) error); ok {
		r1 = rf(dto)
	} else {
		r1 = args.Error(1)
	}

	return r0, r1
}
