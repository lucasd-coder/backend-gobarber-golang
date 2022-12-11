package mock

import (
	"github.com/lucasd-coder/backend-gobarber-golang/internal/model"

	"github.com/stretchr/testify/mock"
)

type MockNotificationsRepository struct {
	mock.Mock
}

func (mock *MockNotificationsRepository) Save(notification *model.Notification) error {
	args := mock.Called(notification)
	return args.Error(0)
}
