package controller

import (
	"fmt"
	"net/http"
	"time"

	"backend-gobarber-golang/internal/dtos"
	"backend-gobarber-golang/internal/interfaces"
	"backend-gobarber-golang/internal/middlewares"
	"backend-gobarber-golang/internal/service"
	"backend-gobarber-golang/internal/util"
	"backend-gobarber-golang/pkg/logger"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	createUsersService      interfaces.CreateUsersService
	showProfileService      interfaces.ShowProfileService
	updateProfileService    interfaces.UpdateProfileService
	updateUserAvatarService interfaces.UpdateUserAvatarService
}

func NewUserController(createUsersService *service.CreateUsersService,
	showProfileService *service.ShowProfileService, updateProfileService *service.UpdateProfileService,
	updateUserAvatarService *service.UpdateUserAvatarService,
) *UserController {
	return &UserController{createUsersService, showProfileService, updateProfileService, updateUserAvatarService}
}

func (user *UserController) InitRoutes(group *gin.RouterGroup) {
	group.POST("/users", user.CreateUser)
	profile := group.Group("/profile", middlewares.EnsureAuthenticated())
	{
		profile.GET("", user.ShowProfile)
		profile.PUT("", user.UpdateProfile)
		profile.PATCH("/avatar", user.UpdateUserAvatar)
	}
}

func (user *UserController) CreateUser(ctx *gin.Context) {
	var body dtos.UserDTO
	if err := ctx.ShouldBindJSON(&body); err != nil {
		logger.Log.Infof("Payload received in invalid. Payload: %+v\n", util.JsonLog(&body))
		HandleError(ctx, "BAD_REQUEST", err.Error(), http.StatusBadRequest)
		return

	}

	start := time.Now()
	resp, err := user.createUsersService.Execute(&body)
	if err != nil {
		logger.Log.Error(err.Error())
		ctx.Error(err)
		return
	}

	logger.Log.Infof("Response received from downstream service. responseTime: %d response: %+v\n ",
		time.Since(start).Milliseconds(), util.JsonLog(resp))

	ctx.JSON(http.StatusCreated, resp)
}

func (user *UserController) ShowProfile(ctx *gin.Context) {
	id := ctx.MustGet("id")

	resp, err := user.showProfileService.Execute(fmt.Sprintf("%v", id))
	if err != nil {
		logger.Log.Error(err.Error())
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

func (user *UserController) UpdateProfile(ctx *gin.Context) {
	var body dtos.UpdateUserProfileDTO
	if err := ctx.ShouldBindJSON(&body); err != nil {
		logger.Log.Infof("Payload received in invalid. Payload: %+v\n", util.JsonLog(&body))
		HandleError(ctx, "BAD_REQUEST", err.Error(), http.StatusBadRequest)
		return

	}

	id := ctx.MustGet("id")

	start := time.Now()
	resp, err := user.updateProfileService.Execute(fmt.Sprintf("%v", id), &body)
	if err != nil {
		logger.Log.Error(err.Error())
		ctx.Error(err)
		return
	}

	logger.Log.Infof("Response received from downstream service. responseTime: %d response: %+v\n ",
		time.Since(start).Milliseconds(), util.JsonLog(resp))

	ctx.JSON(http.StatusOK, resp)
}

func (user *UserController) UpdateUserAvatar(ctx *gin.Context) {
	var form dtos.Form
	if err := ctx.ShouldBind(&form); err != nil {
		logger.Log.Infof("Avatar received in invalid.")
		HandleError(ctx, "BAD_REQUEST", err.Error(), http.StatusBadRequest)
		return
	}

	id := ctx.MustGet("id")

	start := time.Now()
	resp, err := user.updateUserAvatarService.Execute(fmt.Sprintf("%v", id), &form)
	if err != nil {
		logger.Log.Error(err.Error())
		ctx.Error(err)
		return
	}

	logger.Log.Infof("Response received from downstream service. responseTime: %d response: %+v\n ",
		time.Since(start).Milliseconds(), util.JsonLog(resp))

	ctx.JSON(http.StatusOK, resp)
}
