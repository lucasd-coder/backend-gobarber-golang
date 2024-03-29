// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

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
)

// Injectors from wire.go:

func InitializeUserRepository() *repository.UserRepository {
	db := database.GetDatabase()
	userRepository := repository.NewUserRepository(db)
	return userRepository
}

func InitializeUserTokenRepository() *repository.UserTokenRepository {
	db := database.GetDatabase()
	userTokenRepository := repository.NewUserTokenRepository(db)
	return userTokenRepository
}

func InitializeAppointmentRepository() *repository.AppointmentsRepository {
	db := database.GetDatabase()
	appointmentsRepository := repository.NewAppointmentsRepository(db)
	return appointmentsRepository
}

func InitializeNotificationsRepository() *repository.NotificationsRepository {
	configConfig := config.GetConfig()
	client := mongodb.GetClientMongoDB()
	notificationsRepository := repository.NewNotificationsRepository(configConfig, client)
	return notificationsRepository
}

func InitializeCreateUsersService() *service.CreateUsersService {
	userRepository := InitializeUserRepository()
	createUsersService := service.NewCreateUsersService(userRepository)
	return createUsersService
}

func InitializeShowProfileService() *service.ShowProfileService {
	userRepository := InitializeUserRepository()
	showProfileService := service.NewShowProfileService(userRepository)
	return showProfileService
}

func InitializeUpdateProfileService() *service.UpdateProfileService {
	userRepository := InitializeUserRepository()
	updateProfileService := service.NewUpdateProfileService(userRepository)
	return updateProfileService
}

func InitializeJWTService() *service.JWTService {
	jwtService := service.NewJWTService()
	return jwtService
}

func InitializeAuthenticateUserService() *service.AuthenticateUserService {
	userRepository := InitializeUserRepository()
	jwtService := InitializeJWTService()
	authenticateUserService := service.NewAuthenticateUserService(userRepository, jwtService)
	return authenticateUserService
}

func InitializeUpdateUserAvatarService() *service.UpdateUserAvatarService {
	userRepository := InitializeUserRepository()
	diskStorageProvider := InitializeDiskStorageProvider()
	updateUserAvatarService := service.NewUpdateUserAvatarService(userRepository, diskStorageProvider)
	return updateUserAvatarService
}

func InitializeSendForgotPasswordEmailService() *service.SendForgotPasswordEmailService {
	userRepository := InitializeUserRepository()
	userTokenRepository := InitializeUserTokenRepository()
	etherealMailProvider := InitializeEtherealMailProvider()
	renderForgotPasswordTemplate := InitializeRenderForgotPasswordTemplate()
	sendForgotPasswordEmailService := service.NewSendForgotPasswordEmailService(userRepository, userTokenRepository, etherealMailProvider, renderForgotPasswordTemplate)
	return sendForgotPasswordEmailService
}

func InitializeCreateAppointmentService() *service.CreateAppointmentService {
	appointmentsRepository := InitializeAppointmentRepository()
	notificationsRepository := InitializeNotificationsRepository()
	cacheProvider := InitializeCacheProvider()
	createAppointmentService := service.NewCreateAppointmentService(appointmentsRepository, notificationsRepository, cacheProvider)
	return createAppointmentService
}

func InitializeDiskStorageProvider() *storage.DiskStorageProvider {
	diskStorageProvider := storage.NewDiskStorageProvider()
	return diskStorageProvider
}

func InitializeEtherealMailProvider() *storage.EtherealMailProvider {
	etherealMailProvider := storage.NewEtherealMailProvider()
	return etherealMailProvider
}

func InitializeCacheProvider() *storage.CacheProvider {
	client := cache.GetClient()
	cacheProvider := storage.NewCacheProvider(client)
	return cacheProvider
}

func InitializeRenderForgotPasswordTemplate() *template.RenderForgotPasswordTemplate {
	renderForgotPasswordTemplate := template.NewRenderForgotPasswordTemplate()
	return renderForgotPasswordTemplate
}

func InitializeListProviderAppointmentsService() *service.ListProviderAppointmentsService {
	appointmentsRepository := InitializeAppointmentRepository()
	cacheProvider := InitializeCacheProvider()
	listProviderAppointmentsService := service.NewListProviderAppointmentsService(appointmentsRepository, cacheProvider)
	return listProviderAppointmentsService
}

func InitializeListProviderDayAvailabilityService() *service.ListProviderDayAvailabilityService {
	appointmentsRepository := InitializeAppointmentRepository()
	listProviderDayAvailabilityService := service.NewListProviderDayAvailabilityService(appointmentsRepository)
	return listProviderDayAvailabilityService
}

func InitializeListProviderMonthAvailabilityService() *service.ListProviderMonthAvailabilityService {
	appointmentsRepository := InitializeAppointmentRepository()
	listProviderMonthAvailabilityService := service.NewListProviderMonthAvailabilityService(appointmentsRepository)
	return listProviderMonthAvailabilityService
}
