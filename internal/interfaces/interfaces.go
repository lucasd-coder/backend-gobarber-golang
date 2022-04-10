package interfaces

import (
	"mime/multipart"

	"backend-gobarber-golang/internal/dtos"
	"backend-gobarber-golang/internal/model"
	"backend-gobarber-golang/internal/model/external"

	"github.com/golang-jwt/jwt/v4"
)

type (
	CreateUsersService interface {
		Execute(user *dtos.UserDTO) (*dtos.ResponseCreateUserDTO, error)
	}

	ShowProfileService interface {
		Execute(id string) (*dtos.ResponseProfileDTO, error)
	}

	UpdateProfileService interface {
		Execute(id string, user *dtos.UpdateUserProfileDTO) (*dtos.ResponseProfileDTO, error)
	}

	JWTService interface {
		GenerateToken(email string) string
		ValidateToken(tokenString string) (*jwt.Token, error)
	}

	AuthenticateUserService interface {
		Execute(dto *dtos.Credentials) (*dtos.ResponseUserAuthenticatedSuccessDTO, error)
	}

	UpdateUserAvatarService interface {
		Execute(id string, file *dtos.Form) (*dtos.ResponseProfileDTO, error)
	}

	SendForgotPasswordEmailService interface {
		Execute(dto *dtos.ForgotPasswordEmail) error
	}

	UserRepository interface {
		Save(user *model.User)
		Update(user *model.User)
		FindByEmail(email string) *model.User
		FindById(id string) *model.User
	}

	UserTokenRepository interface {
		Generate(userToken *model.UserToken) *model.UserToken
		FindByToken(token string) *model.UserToken
	}

	DiskStorageProvider interface {
		SaveFile(file *multipart.FileHeader) string
		DeleteFile(file string)
	}

	EtherealMailProvider interface {
		SendMail(authSmtp *external.AuthSmtpSendEmail, dto *dtos.SendMailDTO) error
	}

	RenderForgotPasswordTemplate interface {
		Render(variables *external.Variables, email string) *dtos.SendMailDTO
	}
)
