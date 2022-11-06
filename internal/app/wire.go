//go:build wireinject
//+build wireinject

package app

import (
	"github.com/lucasd-coder/backend-gobarber-golang/config"
	"github.com/lucasd-coder/backend-gobarber-golang/internal/infra/repository"
	"github.com/lucasd-coder/backend-gobarber-golang/internal/infra/storage"
	"github.com/lucasd-coder/backend-gobarber-golang/internal/pkg/database"
	"github.com/lucasd-coder/backend-gobarber-golang/internal/pkg/mongodb"
	"github.com/lucasd-coder/backend-gobarber-golang/internal/service"
	"github.com/lucasd-coder/backend-gobarber-golang/internal/template"
	"github.com/lucasd-coder/backend-gobarber-golang/pkg/cache"

	"github.com/google/wire"
)

func InitializeUserRepository() *repository.UserRepository {
	wire.Build(database.GetDatabase, repository.NewUserRepository)
	return &repository.UserRepository{}
}

func InitializeUserTokenRepository() *repository.UserTokenRepository {
	wire.Build(database.GetDatabase, repository.NewUserTokenRepository)
	return &repository.UserTokenRepository{}
}

func InitializeAppointmentRepository() *repository.AppointmentsRepository {
	wire.Build(database.GetDatabase, repository.NewAppointmentsRepository)
	return &repository.AppointmentsRepository{}
}

func InitializeNotificationsRepository() *repository.NotificationsRepository {
	wire.Build(config.GetConfig, mongodb.GetClientMongoDB, repository.NewNotificationsRepository)
	return &repository.NotificationsRepository{}
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

func InitializeUpdateUserAvatarService() *service.UpdateUserAvatarService {
	wire.Build(InitializeUserRepository, InitializeDiskStorageProvider, service.NewUpdateUserAvatarService)
	return &service.UpdateUserAvatarService{}
}

func InitializeSendForgotPasswordEmailService() *service.SendForgotPasswordEmailService {
	wire.Build(InitializeUserRepository, InitializeUserTokenRepository, InitializeEtherealMailProvider, InitializeRenderForgotPasswordTemplate, service.NewSendForgotPasswordEmailService)
	return &service.SendForgotPasswordEmailService{}
}

func InitializeCreateAppointmentService() *service.CreateAppointmentService {
	wire.Build(InitializeAppointmentRepository, InitializeNotificationsRepository, InitializeCacheProvider, service.NewCreateAppointmentService)
	return &service.CreateAppointmentService{}
}

func InitializeDiskStorageProvider() *storage.DiskStorageProvider {
	wire.Build(storage.NewDiskStorageProvider)
	return &storage.DiskStorageProvider{}
}

func InitializeEtherealMailProvider() *storage.EtherealMailProvider {
	wire.Build(storage.NewEtherealMailProvider)
	return &storage.EtherealMailProvider{}
}

func InitializeCacheProvider() *storage.CacheProvider {
	wire.Build(cache.GetClient, storage.NewCacheProvider)
	return &storage.CacheProvider{}
}

func InitializeRenderForgotPasswordTemplate() *template.RenderForgotPasswordTemplate {
	wire.Build(template.NewRenderForgotPasswordTemplate)
	return &template.RenderForgotPasswordTemplate{}
}

func InitializeListProviderAppointmentsService() *service.ListProviderAppointmentsService {
	wire.Build(InitializeAppointmentRepository, InitializeCacheProvider, service.NewListProviderAppointmentsService)
	return &service.ListProviderAppointmentsService{}
}

func InitializeListProviderDayAvailabilityService() *service.ListProviderDayAvailabilityService {
	wire.Build(InitializeAppointmentRepository, service.NewListProviderDayAvailabilityService)
	return &service.ListProviderDayAvailabilityService{}
}

func InitializeListProviderMonthAvailabilityService() *service.ListProviderMonthAvailabilityService {
	wire.Build(InitializeAppointmentRepository, service.NewListProviderMonthAvailabilityService)
	return &service.ListProviderMonthAvailabilityService{}
}
