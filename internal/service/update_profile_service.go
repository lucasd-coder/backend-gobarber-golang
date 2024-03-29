package service

import (
	"github.com/lucasd-coder/backend-gobarber-golang/internal/dtos"
	"github.com/lucasd-coder/backend-gobarber-golang/internal/infra/errs"
	"github.com/lucasd-coder/backend-gobarber-golang/internal/infra/repository"
	"github.com/lucasd-coder/backend-gobarber-golang/internal/interfaces"
	"github.com/lucasd-coder/backend-gobarber-golang/internal/util"
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
		return &dtos.ResponseProfileDTO{}, &errs.AppError{
			Message: "Id invalid.",
			Code:    400,
		}
	}

	user := service.UserRepository.FindById(id)

	if user.Email == "" {
		return &dtos.ResponseProfileDTO{}, &errs.AppError{
			Message: "User not found.",
			Code:    404,
		}
	}

	userWithUpdateEmail := service.UserRepository.FindByEmail(userDto.Email)

	if userWithUpdateEmail.Email != "" && userWithUpdateEmail.ID.String() != id {
		return &dtos.ResponseProfileDTO{}, &errs.AppError{
			Message: "E-mail already in use.",
			Code:    400,
		}
	}

	user.Name = userDto.Name
	user.Email = userDto.Email

	if !util.CheckPasswordHash(userDto.OldPassword, user.Password) {
		return &dtos.ResponseProfileDTO{}, &errs.AppError{
			Message: "You need to inform the old password to set a new password.",
			Code:    400,
		}
	}

	hash, err := util.HashPassword(userDto.Password)
	if err != nil {
		return &dtos.ResponseProfileDTO{}, err
	}

	user.Password = hash

	service.UserRepository.Update(user)

	userResponse := dtos.NewResponseProfileDTO(user.ID.String(), user.Name, user.Email,
		user.Avatar, user.CreatedAt, user.UpdatedAt)

	return userResponse, err
}
