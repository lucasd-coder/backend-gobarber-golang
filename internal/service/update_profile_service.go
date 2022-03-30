package service

import (
	"backend-gobarber-golang/internal/dtos"
	"backend-gobarber-golang/internal/infra/errs"
	"backend-gobarber-golang/internal/infra/repository"
	"backend-gobarber-golang/internal/interfaces"
	"backend-gobarber-golang/internal/util"
)

type UpdateProfileService struct {
	UserRepository interfaces.UserRepository
}

func NewUpdateProfileService(userRepository *repository.UserRepository) *UpdateProfileService {
	return &UpdateProfileService{
		UserRepository: userRepository,
	}
}

func (service *UpdateProfileService) Execute(id string, userDto *dtos.UpdateUserProfileDTO) (*dtos.ResponseProfileDTO, error) {
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

	userWithUpdateEmail := service.UserRepository.FindByEmail(userDto.Email)

	if userWithUpdateEmail.Email != "" && userWithUpdateEmail.ID.String() != id {
		return nil, &errs.AppError{
			Message: "E-mail already in use.",
			Code:    400,
		}
	}

	user.Name = userDto.Name
	user.Email = userDto.Email

	if !util.CheckPasswordHash(userDto.OldPassword, user.Password) {
		return nil, &errs.AppError{
			Message: "You need to inform the old password to set a new password.",
			Code:    400,
		}
	}

	hash, err := util.HashPassword(userDto.Password)
	if err != nil {
		return nil, err
	}

	user.Password = hash

	service.UserRepository.Save(user)

	userResponse := dtos.NewResponseProfileDTO(user.ID.String(), user.Name, user.Email,
		user.Avatar, user.CreatedAt, user.UpdatedAt)

	return userResponse, err
}
