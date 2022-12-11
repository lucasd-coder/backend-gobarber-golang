package service_test

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/lucasd-coder/backend-gobarber-golang/config"
	"github.com/lucasd-coder/backend-gobarber-golang/internal/dtos"
	test "github.com/lucasd-coder/backend-gobarber-golang/internal/mock"
	"github.com/lucasd-coder/backend-gobarber-golang/internal/model"
	"github.com/lucasd-coder/backend-gobarber-golang/internal/model/external"
	"github.com/lucasd-coder/backend-gobarber-golang/internal/service"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var env_value = "SendForgot"

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

func TestSuccessfullySendForgotPassword(t *testing.T) {
	mockRepo := new(test.MockUserRepository)
	mockRepoToken := new(test.MockUserTokenRepository)
	mockEtherealMailProvi := new(test.MockEtherealMailProvider)
	mockRenderForgotTemp := new(test.MockRenderForgotPasswordTemplate)

	var id uuid.UUID = uuid.MustParse("0399e631-e2f0-4df5-b1d0-ca6d567a318c")
	user := model.User{ID: id, Name: "lucas", Email: "lucas@gmail.com", Password: "123456"}

	userToken := model.NewUserToken(user.ID.String())

	userTokenReturn := model.UserToken{
		ID:        uuid.New(),
		Token:     uuid.New(),
		UserID:    user.ID.String(),
		User:      user,
		CreatedAt: time.Now(),
	}

	cfg := SetUpConfig()
	cfg.AppWebUrl = "http://localhost:8080"
	config.ExportConfig(cfg)

	sendMail := &dtos.SendMailDTO{
		From:    "Equipe GoBarber",
		To:      []string{"teste@gmail.com"},
		Message: []byte("Hello world"),
	}

	auth := &external.AuthSmtpSendEmail{
		Host:     cfg.EtherealMail.Host,
		Port:     cfg.EtherealMail.SmtpPort,
		Password: cfg.EtherealMail.Password,
		Username: cfg.EtherealMail.Username,
	}

	mockRepo.On("FindByEmail", user.Email).Return(&user)
	mockRepoToken.On("Generate", &userToken).Return(&userTokenReturn)
	mockRenderForgotTemp.On("Render", mock.Anything, user.Email).Return(sendMail)
	mockEtherealMailProvi.On("SendMail", &auth, &sendMail).Return(nil)

	testService := service.SendForgotPasswordEmailService{
		mockRepo, mockRepoToken,
		mockEtherealMailProvi, mockRenderForgotTemp,
	}

	userDto := dtos.ForgotPasswordEmail{Email: user.Email}

	err := testService.Execute(&userDto)

	assert.Nil(t, err)
}

func SetUpConfig() *config.Config {
	err := setEnvValues()
	if err != nil {
		panic(err)
	}
	var cfg config.Config
	err = cleanenv.ReadEnv(&cfg)
	if err != nil {
		panic(err)
	}

	return &cfg
}

func setEnvValues() error {
	err := os.Setenv("USERNAME_DB", env_value)
	if err != nil {
		return fmt.Errorf("Error setting USERNAME_DB, err = %v", err)
	}

	err = os.Setenv("PASSWORD_DB", env_value)
	if err != nil {
		return fmt.Errorf("Error setting PASSWORD_DB, err = %v", err)
	}

	err = os.Setenv("JWT_SECRET", env_value)
	if err != nil {
		return fmt.Errorf("Error setting JWT_SECRET, err = %v", err)
	}

	err = os.Setenv("JWT_ISSUER", env_value)
	if err != nil {
		return fmt.Errorf("Error setting JWT_ISSUER, err = %v", err)
	}

	err = os.Setenv("HOST_ETHEREAL_MAIL", env_value)
	if err != nil {
		return fmt.Errorf("Error setting HOST_ETHEREAL_MAIL, err = %v", err)
	}

	err = os.Setenv("PORT_ETHEREAL_MAIL", env_value)
	if err != nil {
		return fmt.Errorf("Error setting PORT_ETHEREAL_MAIL, err = %v", err)
	}

	err = os.Setenv("USERNAME_ETHEREAL_MAIL", env_value)
	if err != nil {
		return fmt.Errorf("Error setting USERNAME_ETHEREAL_MAIL, err = %v", err)
	}

	err = os.Setenv("PASSWORD_ETHEREAL_MAIL", env_value)
	if err != nil {
		return fmt.Errorf("Error setting PASSWORD_ETHEREAL_MAIL, err = %v", err)
	}

	err = os.Setenv("APP_NAME", env_value)
	if err != nil {
		return fmt.Errorf("Error setting APP_NAME, err = %v", err)
	}

	err = os.Setenv("APP_VERSION", env_value)
	if err != nil {
		return fmt.Errorf("Error setting APP_VERSION, err = %v", err)
	}

	err = os.Setenv("HTTP_PORT", "8080")
	if err != nil {
		return fmt.Errorf("Error setting HTTP_PORT, err = %v", err)
	}

	err = os.Setenv("LOG_LEVEL", "info")
	if err != nil {
		return fmt.Errorf("Error setting LOG_LEVEL, err = %v", err)
	}

	err = os.Setenv("APP_WEB_URL", "http://localhost:8080")
	if err != nil {
		return fmt.Errorf("Error setting APP_WEB_URL, err = %v", err)
	}

	err = os.Setenv("HOST_DB", "http://localhost:8080")
	if err != nil {
		return fmt.Errorf("Error setting HOST_DB, err = %v", err)
	}

	err = os.Setenv("HOST_MONGODB", "http://localhost:8080")
	if err != nil {
		return fmt.Errorf("Error setting HOST_MONGODB, err = %v", err)
	}

	err = os.Setenv("DATABASE_MONGODB", "http://localhost:8080")
	if err != nil {
		return fmt.Errorf("Error setting DATABASE_MONGODB, err = %v", err)
	}

	err = os.Setenv("PORT_MONGODB", "8080")
	if err != nil {
		return fmt.Errorf("Error setting PORT_MONGODB, err = %v", err)
	}

	err = os.Setenv("PORT_DB", "8080")
	if err != nil {
		return fmt.Errorf("Error setting PORT_DB, err = %v", err)
	}

	err = os.Setenv("PORT_MONGODB", "8080")
	if err != nil {
		return fmt.Errorf("Error setting PORT_MONGODB, err = %v", err)
	}

	err = os.Setenv("REDIS_DB", "3")
	if err != nil {
		return fmt.Errorf("Error setting REDIS_DB, err = %v", err)
	}

	err = os.Setenv("REDIS_PORT", "8080")
	if err != nil {
		return fmt.Errorf("Error setting REDIS_PORT, err = %v", err)
	}

	return nil
}
