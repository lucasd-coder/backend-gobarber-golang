package service

import (
	"github.com/lucasd-coder/backend-gobarber-golang/internal/dtos"
	"github.com/lucasd-coder/backend-gobarber-golang/internal/infra/errs"
	"github.com/lucasd-coder/backend-gobarber-golang/internal/infra/repository"
	"github.com/lucasd-coder/backend-gobarber-golang/internal/interfaces"
	"github.com/lucasd-coder/backend-gobarber-golang/internal/util"
)

type ShowProfileService struct {
	UserRepository interfaces.UserRepository
}

func NewShowProfileService(userRepository *repository.UserRepository) *ShowProfileService {
	return &ShowProfileService{
		UserRepository: userRepository,
	}
}

func (service *ShowProfileService) Execute(id string) (*dtos.ResponseProfileDTO, error) {
	if !util.IsValidUUID(id) {
		return &dtos.ResponseProfileDTO{}, &errs.AppError{
			Message: "Id invalid.",
			Code:    400,
		}
	}

	result := service.UserRepository.FindById(id)

	if result.Email == "" {
		return &dtos.ResponseProfileDTO{}, &errs.AppError{
			Message: "User not found.",
			Code:    404,
		}
	}

	user := dtos.NewResponseProfileDTO(result.ID.String(), result.Name, result.Email,
		result.Avatar, result.CreatedAt, result.UpdatedAt)

	return user, nil
}
