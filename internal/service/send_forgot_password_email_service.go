package service

import (
	"fmt"

	"github.com/lucasd-coder/backend-gobarber-golang/config"
	"github.com/lucasd-coder/backend-gobarber-golang/internal/dtos"
	"github.com/lucasd-coder/backend-gobarber-golang/internal/infra/errs"
	"github.com/lucasd-coder/backend-gobarber-golang/internal/infra/repository"
	"github.com/lucasd-coder/backend-gobarber-golang/internal/infra/storage"
	"github.com/lucasd-coder/backend-gobarber-golang/internal/interfaces"
	"github.com/lucasd-coder/backend-gobarber-golang/internal/model"
	"github.com/lucasd-coder/backend-gobarber-golang/internal/model/external"
	"github.com/lucasd-coder/backend-gobarber-golang/internal/template"
	"github.com/lucasd-coder/backend-gobarber-golang/pkg/logger"
)

type SendForgotPasswordEmailService struct {
	UserRepository               interfaces.UserRepository
	UserTokenRepository          interfaces.UserTokenRepository
	EtherealMailProvider         interfaces.EtherealMailProvider
	RenderForgotPasswordTemplate interfaces.RenderForgotPasswordTemplate
}

func NewSendForgotPasswordEmailService(userRepository *repository.UserRepository,
	userTokenRepository *repository.UserTokenRepository, etherealMailProvider *storage.EtherealMailProvider,
	renderForgotPasswordTemplate *template.RenderForgotPasswordTemplate,
) *SendForgotPasswordEmailService {
	return &SendForgotPasswordEmailService{
		UserRepository:               userRepository,
		UserTokenRepository:          userTokenRepository,
		EtherealMailProvider:         etherealMailProvider,
		RenderForgotPasswordTemplate: renderForgotPasswordTemplate,
	}
}

func (service *SendForgotPasswordEmailService) Execute(dto *dtos.ForgotPasswordEmail) error {
	user := service.UserRepository.FindByEmail(dto.Email)

	if user.Email == "" {
		return &errs.AppError{
			Message: "User not found.",
			Code:    404,
		}
	}

	userToken := model.NewUserToken(user.ID.String())

	token := service.UserTokenRepository.Generate(userToken)

	cfg := config.GetConfig()

	variables := &external.Variables{
		Name: user.Name,
		Link: fmt.Sprintf("%s/reset-password?token=%s`", cfg.AppWebUrl, token.Token.String()),
	}

	msg := service.RenderForgotPasswordTemplate.Render(variables, user.Email)

	auth := &external.AuthSmtpSendEmail{
		Host:     cfg.EtherealMail.Host,
		Port:     cfg.EtherealMail.SmtpPort,
		Password: cfg.EtherealMail.Password,
		Username: cfg.EtherealMail.Username,
	}

	err := service.EtherealMailProvider.SendMail(auth, msg)
	if err != nil {
		logger.Log.Error(err.Error())
		return err
	}

	return err
}
