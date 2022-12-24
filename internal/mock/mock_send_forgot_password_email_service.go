package mock

import (
	"github.com/lucasd-coder/backend-gobarber-golang/internal/dtos"
	"github.com/stretchr/testify/mock"
)

type MockSendForgotPasswordEmailService struct {
	mock.Mock
}

func (mock *MockSendForgotPasswordEmailService) Execute(dto *dtos.ForgotPasswordEmail) error {
	args := mock.Called(dto)

	var r0 error
	if rf, ok := args.Get(0).(func(*dtos.ForgotPasswordEmail) error); ok {
		r0 = rf(dto)
	} else {
		r0 = args.Error(0)
	}

	return r0
}
