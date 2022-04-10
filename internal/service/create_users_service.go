package service

import (
	"backend-gobarber-golang/internal/dtos"
	"backend-gobarber-golang/internal/infra/errs"
	"backend-gobarber-golang/internal/infra/repository"
	"backend-gobarber-golang/internal/interfaces"
	"backend-gobarber-golang/internal/model"
	"backend-gobarber-golang/internal/util"
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
