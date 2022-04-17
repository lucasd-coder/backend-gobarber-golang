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

func TestSendForgotPasswordEmailService_UserNotFound(t *testing.T) {
	mockRepo := new(test.MockUserRepository)
	mockRepoToken := new(test.MockUserTokenRepository)
	mockEtherealMailProvi := new(test.MockEtherealMailProvider)
	mockRenderForgotTemp := new(test.MockRenderForgotPasswordTemplate)

	var id uuid.UUID = uuid.MustParse("0399e631-e2f0-4df5-b1d0-ca6d567a318c")
	user := model.User{ID: id, Name: "lucas", Email: "lucas@gmail.com", Password: "123456"}

	mockRepo.On("FindByEmail", user.Email).Return(&model.User{})

	testService := service.SendForgotPasswordEmailService{
		mockRepo, mockRepoToken,
		mockEtherealMailProvi, mockRenderForgotTemp,
	}

	userDto := dtos.ForgotPasswordEmail{Email: user.Email}

	err := testService.Execute(&userDto)

	assert.NotNil(t, err)
	assert.Equal(t, "User not found.", err.Error())
}
