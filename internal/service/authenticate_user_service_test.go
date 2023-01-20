package service_test

import (
	"testing"

	"github.com/lucasd-coder/backend-gobarber-golang/internal/dtos"
	test "github.com/lucasd-coder/backend-gobarber-golang/internal/mock"
	"github.com/lucasd-coder/backend-gobarber-golang/internal/model"
	"github.com/lucasd-coder/backend-gobarber-golang/internal/service"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestAuthenticateUserService_AuthenticateFailNonExistingUser(t *testing.T) {
	mockRepo := new(test.MockUserRepository)
	jwtService := service.JWTService{}

	email := "lucas@gmail.com"

	mockRepo.On("FindByEmail", email).Return(&model.User{})

	userDto := dtos.Credentials{Email: email, Password: "123456"}

	testService := service.AuthenticateUserService{mockRepo, &jwtService}

	_, err := testService.Execute(&userDto)

	assert.NotNil(t, err)
	assert.Equal(t, "Incorrect email/password combination.", err.Error())
}

func TestAuthenticateUserService_AuthenticateFailWithWrongPassword(t *testing.T) {
	mockRepo := new(test.MockUserRepository)
	jwtService := service.JWTService{}

	var id uuid.UUID = uuid.MustParse("0399e631-e2f0-4df5-b1d0-ca6d567a318c")
	user := model.User{ID: id, Name: "lucas", Email: "lucas@gmail.com", Password: "123456"}

	mockRepo.On("FindByEmail", user.Email).Return(&user)

	userDto := dtos.Credentials{Email: user.Email, Password: "wrong-password"}

	testService := service.AuthenticateUserService{mockRepo, &jwtService}

	_, err := testService.Execute(&userDto)

	assert.NotNil(t, err)
	assert.Equal(t, "Incorrect email/password combination.", err.Error())
}

func TestAuthenticateUserService_AuthenticateSuccessfully(t *testing.T) {
	mockRepo := new(test.MockUserRepository)
	jwtService := service.JWTService{}

	var id uuid.UUID = uuid.MustParse("0399e631-e2f0-4df5-b1d0-ca6d567a318c")

	user := model.User{
		ID:       id,
		Name:     "lucas",
		Email:    "lucas@gmail.com",
		Password: "$2a$08$7xWE3NGEXhnKHYi7wcFUw.wMDtisIPK4T78lmjnSOsYEO.6gTuy1W",
	}

	mockRepo.On("FindByEmail", user.Email).Return(&user)

	userDto := dtos.Credentials{Email: user.Email, Password: "123456"}

	testService := service.AuthenticateUserService{mockRepo, &jwtService}

	resp, err := testService.Execute(&userDto)

	assert.Nil(t, err)
	assert.NotEmpty(t, resp.Response)
	assert.NotEmpty(t, resp.Token)
}
