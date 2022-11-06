package test

import (
	"github.com/lucasd-coder/backend-gobarber-golang/internal/model"

	"github.com/stretchr/testify/mock"
)

type MockUserRepository struct {
	mock.Mock
}

func (mock *MockUserRepository) Save(user *model.User) {
}

func (mock *MockUserRepository) FindByEmail(email string) *model.User {
	args := mock.Called(email)

	r0 := &model.User{}
	if rf, ok := args.Get(0).(func(email string) *model.User); ok {
		r0 = rf(email)
	} else {
		if args.Get(0) != nil {
			r0 = args.Get(0).(*model.User)
		}
	}

	return r0
}

func (mock *MockUserRepository) Update(user *model.User) {
}

func (mock *MockUserRepository) FindById(id string) *model.User {
	args := mock.Called(id)
	result := args.Get(0)
	return result.(*model.User)
}
