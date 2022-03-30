package controller

import (
	"net/http"

	"backend-gobarber-golang/internal/dtos"
	"backend-gobarber-golang/internal/interfaces"
	"backend-gobarber-golang/internal/service"

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
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, resp)
}
