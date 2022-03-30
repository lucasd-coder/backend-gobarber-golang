package service

import (
	"backend-gobarber-golang/internal/dtos"
	"backend-gobarber-golang/internal/infra/errs"
	"backend-gobarber-golang/internal/infra/repository"
	"backend-gobarber-golang/internal/interfaces"
	"backend-gobarber-golang/internal/util"
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
		return nil, &errs.AppError{
			Message: "Id invalid.",
			Code:    400,
		}
	}

	result := service.UserRepository.FindById(id)

	if result.Email == "" {
		return nil, &errs.AppError{
			Message: "User not found.",
			Code:    404,
		}
	}

	user := dtos.NewResponseProfileDTO(result.ID.String(), result.Name, result.Email,
		result.Avatar, result.CreatedAt, result.UpdatedAt)

	return user, nil
}
