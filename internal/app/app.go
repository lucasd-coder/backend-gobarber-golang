package app

import (
	"backend-gobarber-golang/config"
	"backend-gobarber-golang/internal/controller"
	"backend-gobarber-golang/internal/middlewares"
	"backend-gobarber-golang/internal/pkg/database"
	"backend-gobarber-golang/internal/pkg/mongodb"

	"backend-gobarber-golang/pkg/cache"
	"backend-gobarber-golang/pkg/logger"

	"github.com/gin-gonic/gin"
)

func Run(cfg *config.Config) {
	// Log config
	logger.SetUpLog(cfg)

	// Database Config
	database.StartDB(cfg)

	// Mongo Config
	mongodb.SetUpMongoDB(cfg)

	// Redis Config
	cache.SetUpRedis(cfg)

	// Close Database
	defer database.CloseConn()

	// Clode MongoDB
	defer mongodb.CloseConnMongoDB()

	// Http server
	engine := gin.New()
	engine.Use(gin.Recovery())
	engine.Use(middlewares.JSONAppErrorReporter())
	engine.MaxMultipartMemory = 8 << 20 // 8 MiB

	// Routers
	handler := engine.Group("/" + cfg.Name)
	handler.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "UP",
		})
	})

	createUsersService := InitializeCreateUsersService()
	showProfileService := InitializeShowProfileService()
	updateProfileService := InitializeUpdateProfileService()
	updateUserAvatarService := InitializeUpdateUserAvatarService()
	createAppointmentService := InitializeCreateAppointmentService()

	userController := controller.NewUserController(createUsersService, showProfileService,
		updateProfileService, updateUserAvatarService)

	userController.InitRoutes(handler)

	authenticateUserService := InitializeAuthenticateUserService()
	sessionsController := controller.NewSessionsController(authenticateUserService)
	sessionsController.InitRoutes(handler)

	sendForgotPasswordEmailService := InitializeSendForgotPasswordEmailService()
	forgotPasswordEmailController := controller.NewForgotPasswordController(sendForgotPasswordEmailService)
	forgotPasswordEmailController.InitRoutes(handler)

	createAppointmentController := controller.NewAppointmentsController(createAppointmentService)
	createAppointmentController.InitRoutes(handler)

	err := engine.Run(":" + cfg.Port)
	if err != nil {
		panic(err)
	}
}
