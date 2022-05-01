package controller

import (
	"fmt"
	"net/http"

	"backend-gobarber-golang/internal/dtos"
	"backend-gobarber-golang/internal/interfaces"
	"backend-gobarber-golang/internal/middlewares"
	"backend-gobarber-golang/internal/service"
	"backend-gobarber-golang/pkg/logger"

	"github.com/gin-gonic/gin"
)

type AppointmentsController struct {
	createAppointment interfaces.CreateAppointmentService
}

func NewAppointmentsController(createAppointment *service.CreateAppointmentService) *AppointmentsController {
	return &AppointmentsController{
		createAppointment,
	}
}

func (appointment *AppointmentsController) InitRoutes(group *gin.RouterGroup) {
	appointments := group.Group("/appointments", middlewares.EnsureAuthenticated())
	{
		appointments.POST("", appointment.CreateAppointment)
	}
}

func (appointment *AppointmentsController) CreateAppointment(ctx *gin.Context) {
	var body dtos.AppointmentDTO
	if err := ctx.ShouldBindJSON(&body); err != nil {
		HandleError(ctx, "BAD_REQUEST", err.Error(), http.StatusBadRequest)
		return
	}

	id := ctx.MustGet("id")

	resp, err := appointment.createAppointment.Execute(fmt.Sprintf("%v", id), &body)
	if err != nil {
		logger.Log.Error(err.Error())
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusCreated, resp)
}
