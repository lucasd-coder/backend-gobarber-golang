package service

import (
	"github.com/lucasd-coder/backend-gobarber-golang/internal/dtos"
	"github.com/lucasd-coder/backend-gobarber-golang/internal/infra/errs"
	"github.com/lucasd-coder/backend-gobarber-golang/internal/infra/repository"
	"github.com/lucasd-coder/backend-gobarber-golang/internal/interfaces"
	"github.com/lucasd-coder/backend-gobarber-golang/internal/model"
	"github.com/lucasd-coder/backend-gobarber-golang/internal/util"
)

type CreateUsersService struct {
	UserRepository interfaces.UserRepository
}

func NewCreateUsersService(userRepository *repository.UserRepository) *CreateUsersService {
	return &CreateUsersService{
		UserRepository: userRepository,
	}
}

func (createUser *CreateUsersService) Execute(userDto *dtos.UserDTO) (*dtos.ResponseCreateUserDTO, error) {
	result := createUser.UserRepository.FindByEmail(userDto.Email)

	if result.Email != "" {
		return &dtos.ResponseCreateUserDTO{}, &errs.AppError{
			Message: "Email address already used by another",
			Code:    400,
		}
	}

	hash, err := util.HashPassword(userDto.Password)
	if err != nil {
		return &dtos.ResponseCreateUserDTO{}, err
	}

	user := model.NewUser(userDto.Name, userDto.Email, hash)

	createUser.UserRepository.Save(user)

	responseUserDto := dtos.NewResponseCreateUserDTO(user.Name, user.Email)

	return responseUserDto, err
}
