package app

import (
	"backend-gobarber-golang/config"
	"backend-gobarber-golang/internal/controller"
	"backend-gobarber-golang/internal/middlewares"
	"backend-gobarber-golang/internal/pkg/database"

	"backend-gobarber-golang/pkg/logger"

	"github.com/gin-gonic/gin"
)

func Run(cfg *config.Config) {
	// Log config
	logger.SetUpLog(cfg)

	// Database Config
	database.StartDB(cfg)

	// Close Database
	defer database.CloseConn()

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

	userController := controller.NewUserController(createUsersService, showProfileService,
		updateProfileService, updateUserAvatarService)

	userController.InitRoutes(handler)

	authenticateUserService := InitializeAuthenticateUserService()
	sessionsController := controller.NewSessionsController(authenticateUserService)
	sessionsController.InitRoutes(handler)

	err := engine.Run(":" + cfg.Port)
	if err != nil {
		panic(err)
	}
}
