package mock

import (
	"github.com/lucasd-coder/backend-gobarber-golang/internal/dtos"
	"github.com/lucasd-coder/backend-gobarber-golang/internal/model/external"

	"github.com/stretchr/testify/mock"
)

type MockRenderForgotPasswordTemplate struct {
	mock.Mock
}

func (mock *MockRenderForgotPasswordTemplate) Render(variables *external.Variables, email string) *dtos.SendMailDTO {
	args := mock.Called(variables, email)

	var r0 *dtos.SendMailDTO
	if rf, ok := args.Get(0).(func(variables *external.Variables, email string) *dtos.SendMailDTO); ok {
		r0 = rf(variables, email)
	} else {
		if args.Get(0) != nil {
			r0 = args.Get(0).(*dtos.SendMailDTO)
		}
	}

	return r0
}
