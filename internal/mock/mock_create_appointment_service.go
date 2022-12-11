package mock

import (
	"github.com/lucasd-coder/backend-gobarber-golang/internal/dtos"
	"github.com/lucasd-coder/backend-gobarber-golang/internal/model"
	"github.com/stretchr/testify/mock"
)

type MockCreateAppointmentService struct {
	mock.Mock
}

func NewMockCreateAppointmentService() *MockCreateAppointmentService {
	return &MockCreateAppointmentService{}
}

func (mock *MockCreateAppointmentService) Execute(userId string, dto *dtos.AppointmentDTO) (*model.Appointment, error) {
	args := mock.Called(userId, dto)

	r0 := &model.Appointment{}

	if rf, ok := args.Get(0).(func(string, *dtos.AppointmentDTO) *model.Appointment); ok {
		r0 = rf(userId, dto)
	} else {
		if args.Get(0) != nil {
			r0 = args.Get(0).(*model.Appointment)
		}
	}

	var r1 error
	if rf, ok := args.Get(1).(func(string, *dtos.AppointmentDTO) error); ok {
		r1 = rf(userId, dto)
	} else {
		r1 = args.Error(1)
	}

	return r0, r1
}
