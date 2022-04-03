package service

import (
	"backend-gobarber-golang/internal/dtos"
	"backend-gobarber-golang/internal/infra/errs"
	"backend-gobarber-golang/internal/infra/repository"
	"backend-gobarber-golang/internal/infra/storage"
	"backend-gobarber-golang/internal/interfaces"
	"backend-gobarber-golang/internal/util"
)

type UpdateUserAvatarService struct {
	UserRepository      interfaces.UserRepository
	DiskStorageProvider interfaces.DiskStorageProvider
}

func NewUpdateUserAvatarService(userRepository *repository.UserRepository, diskStorageProvider *storage.DiskStorageProvider) *UpdateUserAvatarService {
	return &UpdateUserAvatarService{
		UserRepository:      userRepository,
		DiskStorageProvider: diskStorageProvider,
	}
}

func (service *UpdateUserAvatarService) Execute(id string, file *dtos.Form) (*dtos.ResponseProfileDTO, error) {
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

	if user.Avatar != "" {
		service.DiskStorageProvider.DeleteFile(user.Avatar)
	}

	filename := service.DiskStorageProvider.SaveFile(file.Avatar)

	user.Avatar = filename

	service.UserRepository.Update(user)

	userResponse := dtos.NewResponseProfileDTO(user.ID.String(), user.Name, user.Email,
		user.Avatar, user.CreatedAt, user.UpdatedAt)

	return userResponse, nil
}
