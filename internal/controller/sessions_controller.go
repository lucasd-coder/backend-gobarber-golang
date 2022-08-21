package controller

import (
	"net/http"

	"github.com/lucasd-coder/backend-gobarber-golang/internal/dtos"
	"github.com/lucasd-coder/backend-gobarber-golang/internal/interfaces"
	"github.com/lucasd-coder/backend-gobarber-golang/internal/service"
	"github.com/lucasd-coder/backend-gobarber-golang/pkg/logger"

	"github.com/gin-gonic/gin"
)

type SessionsController struct {
	authenticateUserService interfaces.AuthenticateUserService
}

func NewSessionsController(authenticateUserService *service.AuthenticateUserService) *SessionsController {
	return &SessionsController{
		authenticateUserService,
	}
}

func (sessions *SessionsController) InitRoutes(group *gin.RouterGroup) {
	group.POST("/sessions", sessions.AuthenticateUser)
}

func (sessions *SessionsController) AuthenticateUser(ctx *gin.Context) {
	var body dtos.Credentials
	if err := ctx.ShouldBindJSON(&body); err != nil {
		HandleError(ctx, "BAD_REQUEST", err.Error(), http.StatusBadRequest)
		return

	}

	resp, err := sessions.authenticateUserService.Execute(&body)
	if err != nil {
		logger.Log.Error(err.Error())
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, resp)
}
