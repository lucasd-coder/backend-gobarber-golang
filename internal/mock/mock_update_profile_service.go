package mock

import (
	"github.com/lucasd-coder/backend-gobarber-golang/internal/dtos"
	"github.com/stretchr/testify/mock"
)

type MockUpdateProfileService struct {
	mock.Mock
}

func (mock *MockUpdateProfileService) Execute(id string, userDto *dtos.UpdateUserProfileDTO) (*dtos.ResponseProfileDTO, error) {
	args := mock.Called(id, userDto)

	r0 := &dtos.ResponseProfileDTO{}

	if rf, ok := args.Get(0).(func(string, *dtos.UpdateUserProfileDTO) *dtos.ResponseProfileDTO); ok {
		r0 = rf(id, userDto)
	} else {
		if args.Get(0) != nil {
			r0 = args.Get(0).(*dtos.ResponseProfileDTO)
		}
	}

	var r1 error
	if rf, ok := args.Get(1).(func(string, *dtos.UpdateUserProfileDTO) error); ok {
		r1 = rf(id, userDto)
	} else {
		r1 = args.Error(1)
	}

	return r0, r1
}
