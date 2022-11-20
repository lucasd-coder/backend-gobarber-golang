package app

import (
	"github.com/lucasd-coder/backend-gobarber-golang/config"
	"github.com/lucasd-coder/backend-gobarber-golang/internal/controller"
	"github.com/lucasd-coder/backend-gobarber-golang/internal/middlewares"
	"github.com/lucasd-coder/backend-gobarber-golang/internal/pkg/database"
	"github.com/lucasd-coder/backend-gobarber-golang/internal/pkg/mongodb"

	"github.com/lucasd-coder/backend-gobarber-golang/pkg/cache"
	"github.com/lucasd-coder/backend-gobarber-golang/pkg/logger"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
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

	// Close MongoDB
	defer mongodb.CloseConnMongoDB()

	// Http server
	engine := gin.New()
	engine.Use(gin.Recovery())
	engine.Use(middlewares.PrometheusHandler())
	engine.Use(middlewares.JSONAppErrorReporter())
	engine.MaxMultipartMemory = 8 << 20 // 8 MiB

	// Routers
	handler := engine.Group("/" + cfg.Name)
	handler.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "UP",
		})
	})
	handler.GET("/metrics", gin.WrapH(promhttp.Handler()))

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

	listProviderAppointments := InitializeListProviderAppointmentsService()
	listProviderAppointmentsController := controller.NewProviderAppointmentsController(listProviderAppointments)
	listProviderAppointmentsController.InitRoutes(handler)

	err := engine.Run(":" + cfg.Port)
	if err != nil {
		panic(err)
	}
}
