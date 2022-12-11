package mock

import (
	"github.com/lucasd-coder/backend-gobarber-golang/internal/dtos"
	"github.com/stretchr/testify/mock"
)

type MockCreateUsersService struct {
	mock.Mock
}

func NewMockCreateUsersService() *MockCreateUsersService {
	return &MockCreateUsersService{}
}

func (mock *MockCreateUsersService) Execute(userDto *dtos.UserDTO) (*dtos.ResponseCreateUserDTO, error) {
	args := mock.Called(userDto)

	r0 := &dtos.ResponseCreateUserDTO{}

	if rf, ok := args.Get(0).(func(*dtos.UserDTO) *dtos.ResponseCreateUserDTO); ok {
		r0 = rf(userDto)
	} else {
		if args.Get(0) != nil {
			r0 = args.Get(0).(*dtos.ResponseCreateUserDTO)
		}
	}

	var r1 error
	if rf, ok := args.Get(1).(func(*dtos.UserDTO) error); ok {
		r1 = rf(userDto)
	} else {
		r1 = args.Error(1)
	}

	return r0, r1
}
