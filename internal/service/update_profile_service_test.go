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

func TestUpdateProfileService_InvalidID(t *testing.T) {
	mockRepo := new(test.MockUserRepository)

	var id uuid.UUID = uuid.MustParse("0399e631-e2f0-4df5-b1d0-ca6d567a318c")
	user := model.User{ID: id, Name: "lucas", Email: "lucas@gmail.com", Password: "123456"}

	userDto := dtos.UpdateUserProfileDTO{Name: user.Name, Email: user.Email, Password: user.Password}

	testService := service.UpdateProfileService{mockRepo}
	idInvalid := " "

	_, err := testService.Execute(idInvalid, &userDto)

	assert.NotNil(t, err)
	assert.Equal(t, "Id invalid.", err.Error())
}

func TestUpdateProfileService_UserNotFound(t *testing.T) {
	mockRepo := new(test.MockUserRepository)

	var id uuid.UUID = uuid.MustParse("0399e631-e2f0-4df5-b1d0-ca6d567a318c")
	user := model.User{ID: id, Name: "lucas", Email: "lucas@gmail.com", Password: "123456"}

	mockRepo.On("FindById", user.ID.String()).Return(&model.User{})

	userDto := dtos.UpdateUserProfileDTO{Name: user.Name, Email: user.Email, Password: user.Password}

	testService := service.UpdateProfileService{mockRepo}

	_, err := testService.Execute(id.String(), &userDto)

	assert.NotNil(t, err)
	assert.Equal(t, "User not found.", err.Error())
}

func TestUpdateProfileService_EmailAlreadyInUse(t *testing.T) {
	mockRepo := new(test.MockUserRepository)

	var id uuid.UUID = uuid.MustParse("0399e631-e2f0-4df5-b1d0-ca6d567a318c")
	var idExists uuid.UUID = uuid.MustParse("b7ca51e4-cd89-4bd0-ac1e-a6d09bcf0e10")

	user := model.User{ID: id, Name: "lucas", Email: "lucas@gmail.com", Password: "123456"}
	userExists := model.User{ID: idExists, Name: "lucas", Email: "lucas123@gmail.com", Password: "123456"}

	mockRepo.On("FindById", user.ID.String()).Return(&user)
	mockRepo.On("FindByEmail", user.Email).Return(&userExists)

	userDto := dtos.UpdateUserProfileDTO{
		Name: user.Name, Email: user.Email,
		Password: user.Password, OldPassword: "123456",
	}

	testService := service.UpdateProfileService{mockRepo}

	_, err := testService.Execute(user.ID.String(), &userDto)

	assert.NotNil(t, err)
	assert.Equal(t, "E-mail already in use.", err.Error())
}

func TestUpdateProfileService_InvalidOldPassword(t *testing.T) {
	mockRepo := new(test.MockUserRepository)

	var id uuid.UUID = uuid.MustParse("0399e631-e2f0-4df5-b1d0-ca6d567a318c")

	user := model.User{ID: id, Name: "lucas", Email: "lucas@gmail.com", Password: "123456"}

	mockRepo.On("FindById", user.ID.String()).Return(&user)
	mockRepo.On("FindByEmail", user.Email).Return(&user)

	userDto := dtos.UpdateUserProfileDTO{
		Name: user.Name, Email: user.Email,
		Password: user.Password, OldPassword: "invalid",
	}

	testService := service.UpdateProfileService{mockRepo}

	_, err := testService.Execute(user.ID.String(), &userDto)

	assert.NotNil(t, err)
	assert.Equal(t, "You need to inform the old password to set a new password.", err.Error())
}

func TestUpdateProfileService_UpdatedSuccess(t *testing.T) {
	mockRepo := new(test.MockUserRepository)

	var id uuid.UUID = uuid.MustParse("0399e631-e2f0-4df5-b1d0-ca6d567a318c")

	user := model.User{ID: id, Name: "lucas", Email: "lucas@gmail.com", Password: "$2a$08$7xWE3NGEXhnKHYi7wcFUw.wMDtisIPK4T78lmjnSOsYEO.6gTuy1W"}

	mockRepo.On("FindById", user.ID.String()).Return(&user)
	mockRepo.On("FindByEmail", user.Email).Return(&user)

	userDto := dtos.UpdateUserProfileDTO{
		Name: user.Name, Email: user.Email,
		Password: "12345678", OldPassword: "123456",
	}

	testService := service.UpdateProfileService{mockRepo}

	response, err := testService.Execute(user.ID.String(), &userDto)

	assert.Nil(t, err)
	assert.Equal(t, userDto.Email, response.Email)
	assert.Equal(t, userDto.Name, response.Name)
}
