package mock

import (
	"github.com/lucasd-coder/backend-gobarber-golang/internal/dtos"
	"github.com/stretchr/testify/mock"
)

type MockAuthenticateUserService struct {
	mock.Mock
}

func NewAuthenticateUserService() *MockAuthenticateUserService {
	return &MockAuthenticateUserService{}
}

func (mock *MockAuthenticateUserService) Execute(dto *dtos.Credentials) (*dtos.ResponseUserAuthenticatedSuccessDTO, error) {
	args := mock.Called(dto)

	r0 := &dtos.ResponseUserAuthenticatedSuccessDTO{}

	if rf, ok := args.Get(0).(func(*dtos.Credentials) *dtos.ResponseUserAuthenticatedSuccessDTO); ok {
		r0 = rf(dto)
	} else {
		if args.Get(0) != nil {
			r0 = args.Get(0).(*dtos.ResponseUserAuthenticatedSuccessDTO)
		}
	}

	var r1 error
	if rf, ok := args.Get(1).(func(*dtos.Credentials) error); ok {
		r1 = rf(dto)
	} else {
		r1 = args.Error(1)
	}

	return r0, r1
}
