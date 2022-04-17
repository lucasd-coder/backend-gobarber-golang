package test

import (
	"backend-gobarber-golang/internal/model"

	"github.com/stretchr/testify/mock"
)

type MockUserRepository struct {
	mock.Mock
}

func (mock *MockUserRepository) Save(user *model.User) {
}

func (mock *MockUserRepository) FindByEmail(email string) *model.User {
	args := mock.Called(email)
	result := args.Get(0)
	return result.(*model.User)
}

func (mock *MockUserRepository) Update(user *model.User) {
}

func (mock *MockUserRepository) FindById(id string) *model.User {
	args := mock.Called(id)
	result := args.Get(0)
	return result.(*model.User)
}
