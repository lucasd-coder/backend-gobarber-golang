//go:build wireinject
// +build wireinject

package app

import (
	"backend-gobarber-golang/internal/infra/repository"
	"backend-gobarber-golang/internal/infra/storage"
	"backend-gobarber-golang/internal/pkg/database"
	"backend-gobarber-golang/internal/service"

	"github.com/google/wire"
)

func InitializeUserRepository() *repository.UserRepository {
	wire.Build(database.GetDatabase, repository.NewUserRepository)
	return &repository.UserRepository{}
}

func InitializeCreateUsersService() *service.CreateUsersService {
	wire.Build(InitializeUserRepository, service.NewCreateUsersService)
	return &service.CreateUsersService{}
}

func InitializeShowProfileService() *service.ShowProfileService {
	wire.Build(InitializeUserRepository, service.NewShowProfileService)
	return &service.ShowProfileService{}
}

func InitializeUpdateProfileService() *service.UpdateProfileService {
	wire.Build(InitializeUserRepository, service.NewUpdateProfileService)
	return &service.UpdateProfileService{}
}

func InitializeJWTService() *service.JWTService {
	wire.Build(service.NewJWTService)
	return &service.JWTService{}
}

func InitializeAuthenticateUserService() *service.AuthenticateUserService {
	wire.Build(InitializeUserRepository, InitializeJWTService, service.NewAuthenticateUserService)
	return &service.AuthenticateUserService{}
}

func InitializeDiskStorageProvider() *storage.DiskStorageProvider {
	wire.Build(storage.NewDiskStorageProvider)
	return &storage.DiskStorageProvider{}
}

func InitializeUpdateUserAvatarService() *service.UpdateUserAvatarService {
	wire.Build(InitializeUserRepository, InitializeDiskStorageProvider, service.NewUpdateUserAvatarService)
	return &service.UpdateUserAvatarService{}
}
