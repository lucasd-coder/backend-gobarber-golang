package test

import (
	"backend-gobarber-golang/internal/dtos"
	"backend-gobarber-golang/internal/model/external"

	"github.com/stretchr/testify/mock"
)

type MockRenderForgotPasswordTemplate struct {
	mock.Mock
}

func (mock *MockRenderForgotPasswordTemplate) Render(variables *external.Variables, email string) *dtos.SendMailDTO {
	return &dtos.SendMailDTO{}
}
