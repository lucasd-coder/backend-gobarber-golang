package mock

import (
	"github.com/lucasd-coder/backend-gobarber-golang/internal/dtos"
	"github.com/stretchr/testify/mock"
)

type MockUpdateUserAvatarService struct {
	mock.Mock
}

func NewMockUpdateUserAvatarService() *MockUpdateUserAvatarService {
	return &MockUpdateUserAvatarService{}
}

func (mock *MockUpdateUserAvatarService) Execute(id string, file *dtos.Form) (*dtos.ResponseProfileDTO, error) {
	args := mock.Called(id, file)

	r0 := &dtos.ResponseProfileDTO{}

	if rf, ok := args.Get(0).(func(string, *dtos.Form) *dtos.ResponseProfileDTO); ok {
		r0 = rf(id, file)
	} else {
		if args.Get(0) != nil {
			r0 = args.Get(0).(*dtos.ResponseProfileDTO)
		}
	}

	var r1 error
	if rf, ok := args.Get(1).(func(string, *dtos.Form) error); ok {
		r1 = rf(id, file)
	} else {
		r1 = args.Error(1)
	}

	return r0, r1
}
