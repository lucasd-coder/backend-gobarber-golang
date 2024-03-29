package controller

import (
	"net/http"

	"github.com/lucasd-coder/backend-gobarber-golang/internal/dtos"
	"github.com/lucasd-coder/backend-gobarber-golang/internal/interfaces"
	"github.com/lucasd-coder/backend-gobarber-golang/internal/service"
	"github.com/lucasd-coder/backend-gobarber-golang/pkg/logger"

	"github.com/gin-gonic/gin"
)

type ForgotPasswordController struct {
	sendForgotPasswordEmailService interfaces.SendForgotPasswordEmailService
}

func NewForgotPasswordController(sendForgotPasswordEmailService *service.SendForgotPasswordEmailService) *ForgotPasswordController {
	return &ForgotPasswordController{
		sendForgotPasswordEmailService,
	}
}

func (forgotPassword *ForgotPasswordController) InitRoutes(group *gin.RouterGroup) {
	group.POST("/forgot", forgotPassword.ForgotPassword)
}

func (forgotPassword *ForgotPasswordController) ForgotPassword(ctx *gin.Context) {
	var body dtos.ForgotPasswordEmail
	if err := ctx.ShouldBindJSON(&body); err != nil {
		HandleError(ctx, "BAD_REQUEST", err.Error(), http.StatusBadRequest)
		return

	}

	err := forgotPassword.sendForgotPasswordEmailService.Execute(&body)
	if err != nil {
		logger.Log.Error(err.Error())
		ctx.Error(err)
		return
	}

	ctx.Status(http.StatusAccepted)
}
