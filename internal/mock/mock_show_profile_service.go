package mock

import (
	"github.com/lucasd-coder/backend-gobarber-golang/internal/dtos"
	"github.com/stretchr/testify/mock"
)

type MockShowProfileService struct {
	mock.Mock
}

func (mock *MockShowProfileService) Execute(id string) (*dtos.ResponseProfileDTO, error) {
	args := mock.Called(id)

	r0 := &dtos.ResponseProfileDTO{}

	if rf, ok := args.Get(0).(func(string) *dtos.ResponseProfileDTO); ok {
		r0 = rf(id)
	} else {
		if args.Get(0) != nil {
			r0 = args.Get(0).(*dtos.ResponseProfileDTO)
		}
	}

	var r1 error
	if rf, ok := args.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = args.Error(1)
	}

	return r0, r1
}
