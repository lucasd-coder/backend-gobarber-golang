package service

import (
	"github.com/lucasd-coder/backend-gobarber-golang/internal/dtos"
	"github.com/lucasd-coder/backend-gobarber-golang/internal/infra/errs"
	"github.com/lucasd-coder/backend-gobarber-golang/internal/infra/repository"
	"github.com/lucasd-coder/backend-gobarber-golang/internal/interfaces"
	"github.com/lucasd-coder/backend-gobarber-golang/internal/util"
)

type AuthenticateUserService struct {
	UserRepository interfaces.UserRepository
	JWTService     interfaces.JWTService
}

func NewAuthenticateUserService(userRepository *repository.UserRepository, jwtService *JWTService) *AuthenticateUserService {
	return &AuthenticateUserService{
		UserRepository: userRepository,
		JWTService:     jwtService,
	}
}

func (service *AuthenticateUserService) Execute(dto *dtos.Credentials) (*dtos.ResponseUserAuthenticatedSuccessDTO, error) {
	user := service.UserRepository.FindByEmail(dto.Email)

	if user.Email == "" {
		return &dtos.ResponseUserAuthenticatedSuccessDTO{}, &errs.AppError{
			Message: "Incorrect email/password combination.",
			Code:    401,
		}
	}

	if !util.CheckPasswordHash(dto.Password, user.Password) {
		return &dtos.ResponseUserAuthenticatedSuccessDTO{}, &errs.AppError{
			Message: "Incorrect email/password combination.",
			Code:    401,
		}
	}

	token := service.JWTService.GenerateToken(user.ID.String())

	userResponse := dtos.NewResponseProfileDTO(user.ID.String(), user.Name, user.Email,
		user.Avatar, user.CreatedAt, user.UpdatedAt)

	response := dtos.NewResponseUserAuthenticatedSuccessDTO(*userResponse, token)

	return response, nil
}
