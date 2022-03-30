package service

import (
	"backend-gobarber-golang/internal/dtos"
	"backend-gobarber-golang/internal/infra/errs"
	"backend-gobarber-golang/internal/infra/repository"
	"backend-gobarber-golang/internal/interfaces"
	"backend-gobarber-golang/internal/util"
)

type UpdateUserAvatarService struct {
	UserRepository interfaces.UserRepository
}

func NewUpdateUserAvatarService(userRepository *repository.UserRepository) *UpdateUserAvatarService {
	return &UpdateUserAvatarService{
		UserRepository: userRepository,
	}
}

func (service *UpdateUserAvatarService) Execute(id, avatarFilename string) (*dtos.ResponseProfileDTO, error) {
	if !util.IsValidUUID(id) {
		return nil, &errs.AppError{
			Message: "Id invalid.",
			Code:    400,
		}
	}

	user := service.UserRepository.FindById(id)

	if user.Email == "" {
		return nil, &errs.AppError{
			Message: "User not found.",
			Code:    404,
		}
	}

	user.Avatar = avatarFilename

	service.UserRepository.Save(user)

	userResponse := dtos.NewResponseProfileDTO(user.ID.String(), user.Name, user.Email,
		user.Avatar, user.CreatedAt, user.UpdatedAt)

	return userResponse, nil
}
