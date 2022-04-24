package test

import (
	"backend-gobarber-golang/internal/model"

	"github.com/stretchr/testify/mock"
)

type MockNotificationsRepository struct {
	mock.Mock
}

func (mock *MockNotificationsRepository) Save(notification *model.Notification) {
}
