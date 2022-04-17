package service_test

import (
	"testing"

	"backend-gobarber-golang/internal/dtos"
	"backend-gobarber-golang/internal/model"
	"backend-gobarber-golang/internal/service"
	"backend-gobarber-golang/internal/test"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestUpdateUserAvatarService_InvalidID(t *testing.T) {
	mockRepo := new(test.MockUserRepository)
	mockDiskStorage := new(test.MockDiskStorageProvider)

	testService := service.UpdateUserAvatarService{mockRepo, mockDiskStorage}

	avatar := dtos.Form{}

	idInvalid := " "

	_, err := testService.Execute(idInvalid, &avatar)

	assert.NotNil(t, err)
	assert.Equal(t, "Id invalid.", err.Error())
}

func TestUpdateUserAvatarService_UserNotFound(t *testing.T) {
	mockRepo := new(test.MockUserRepository)
	mockDiskStorage := new(test.MockDiskStorageProvider)

	testService := service.UpdateUserAvatarService{mockRepo, mockDiskStorage}

	var id uuid.UUID = uuid.MustParse("0399e631-e2f0-4df5-b1d0-ca6d567a318c")

	mockRepo.On("FindById", id.String()).Return(&model.User{})

	avatar := dtos.Form{}

	_, err := testService.Execute(id.String(), &avatar)

	assert.NotNil(t, err)
	assert.Equal(t, "User not found.", err.Error())
}

func TestUpdateUserAvatarService_UpdatedSuccess(t *testing.T) {
	mockRepo := new(test.MockUserRepository)
	mockDiskStorage := new(test.MockDiskStorageProvider)

	testService := service.UpdateUserAvatarService{mockRepo, mockDiskStorage}

	var id uuid.UUID = uuid.MustParse("0399e631-e2f0-4df5-b1d0-ca6d567a318c")
	user := model.User{ID: id, Name: "lucas", Email: "lucas@gmail.com", Password: "$2a$08$7xWE3NGEXhnKHYi7wcFUw.wMDtisIPK4T78lmjnSOsYEO.6gTuy1W"}

	mockRepo.On("FindById", id.String()).Return(&user)

	avatar := dtos.Form{}

	response, err := testService.Execute(id.String(), &avatar)

	assert.Nil(t, err)
	assert.Equal(t, user.Email, response.Email)
	assert.Equal(t, user.Name, response.Name)
}
