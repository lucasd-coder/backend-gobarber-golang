package service_test

import (
	"testing"

	"github.com/lucasd-coder/backend-gobarber-golang/internal/model"
	"github.com/lucasd-coder/backend-gobarber-golang/internal/service"
	"github.com/lucasd-coder/backend-gobarber-golang/internal/test"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestShowProfileService_InvalidID(t *testing.T) {
	mockRepo := new(test.MockUserRepository)

	testService := service.ShowProfileService{mockRepo}
	idInvalid := " "

	_, err := testService.Execute(idInvalid)

	assert.NotNil(t, err)
	assert.Equal(t, "Id invalid.", err.Error())
}

func TestShowProfileService_UserNotFound(t *testing.T) {
	mockRepo := new(test.MockUserRepository)

	var id uuid.UUID = uuid.MustParse("0399e631-e2f0-4df5-b1d0-ca6d567a318c")
	user := model.User{ID: id, Name: "lucas", Email: "lucas@gmail.com", Password: "123456"}

	mockRepo.On("FindById", user.ID.String()).Return(&model.User{})

	testService := service.ShowProfileService{mockRepo}

	_, err := testService.Execute(id.String())

	assert.NotNil(t, err)
	assert.Equal(t, "User not found.", err.Error())
}

func TestShowProfileService_ShowProfileSuccessfully(t *testing.T) {
	mockRepo := new(test.MockUserRepository)

	var id uuid.UUID = uuid.MustParse("0399e631-e2f0-4df5-b1d0-ca6d567a318c")
	user := model.User{ID: id, Name: "lucas", Email: "lucas@gmail.com", Password: "123456"}

	mockRepo.On("FindById", user.ID.String()).Return(&user)

	testService := service.ShowProfileService{mockRepo}

	result, err := testService.Execute(id.String())

	assert.Nil(t, err)
	assert.Equal(t, result.Email, user.Email)
	assert.Equal(t, result.Name, user.Name)
}
