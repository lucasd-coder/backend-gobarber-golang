package service_test

import (
	"testing"

	"github.com/lucasd-coder/backend-gobarber-golang/internal/dtos"
	"github.com/lucasd-coder/backend-gobarber-golang/internal/model"
	"github.com/lucasd-coder/backend-gobarber-golang/internal/service"
	"github.com/lucasd-coder/backend-gobarber-golang/internal/test"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestCreateUsersService_UserAddressOfAlreadyExisting(t *testing.T) {
	mockRepo := new(test.MockUserRepository)

	var id uuid.UUID = uuid.MustParse("0399e631-e2f0-4df5-b1d0-ca6d567a318c")
	user := model.User{ID: id, Name: "lucas", Email: "lucas@gmail.com", Password: "123456"}

	mockRepo.On("FindByEmail", user.Email).Return(&user)

	userDto := dtos.UserDTO{Name: user.Name, Email: user.Email, Password: user.Password}

	testService := service.CreateUsersService{mockRepo}

	_, err := testService.Execute(&userDto)

	assert.NotNil(t, err)
	assert.Equal(t, "Email address already used by another", err.Error())
}

func TestCreateUserService_UserValid(t *testing.T) {
	mockRepo := new(test.MockUserRepository)

	var id uuid.UUID = uuid.MustParse("0399e631-e2f0-4df5-b1d0-ca6d567a318c")
	user := model.User{ID: id, Name: "lucas", Email: "lucas@gmail.com", Password: "123456"}

	mockRepo.On("FindByEmail", user.Email).Return(&model.User{})
	mockRepo.On("Save", &user).Return(nil)

	userDto := dtos.UserDTO{Name: user.Name, Email: user.Email, Password: user.Password}

	testService := service.CreateUsersService{mockRepo}

	response, err := testService.Execute(&userDto)

	assert.Equal(t, userDto.Email, response.Email)
	assert.Equal(t, userDto.Name, response.Name)
	assert.Nil(t, err)
}
