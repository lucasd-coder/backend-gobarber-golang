package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/lucasd-coder/backend-gobarber-golang/internal/dtos"
	"github.com/lucasd-coder/backend-gobarber-golang/internal/interfaces"
	"github.com/lucasd-coder/backend-gobarber-golang/internal/middlewares"
	"github.com/lucasd-coder/backend-gobarber-golang/internal/service"
	"github.com/lucasd-coder/backend-gobarber-golang/pkg/logger"
)

type ProviderAppointments struct {
	providerAppointments interfaces.ListProviderAppointmentsService
}

func NewProviderAppointmentsController(providerAppointments *service.ListProviderAppointmentsService) *ProviderAppointments {
	return &ProviderAppointments{
		providerAppointments,
	}
}

func (providerAppointments *ProviderAppointments) InitRoutes(group *gin.RouterGroup) {
	appointments := group.Group("/providers", middlewares.EnsureAuthenticated())
	{
		appointments.GET("", providerAppointments.ListProviders)
	}
}

func (providerAppointments *ProviderAppointments) ListProviders(ctx *gin.Context) {
	id := ctx.MustGet("id")

	day := ctx.Query("day")
	intDay, err := strconv.Atoi(day)
	if err != nil {
		logger.Log.Error(err.Error())
		HandleError(ctx, "BAD_REQUEST", err.Error(), http.StatusBadRequest)
		return
	}

	month := ctx.Query("month")
	intMonth, err := strconv.Atoi(month)
	if err != nil {
		logger.Log.Error(err.Error())
		HandleError(ctx, "BAD_REQUEST", err.Error(), http.StatusBadRequest)
		return
	}

	year := ctx.Query("year")
	intYear, err := strconv.Atoi(year)
	if err != nil {
		logger.Log.Error(err.Error())
		HandleError(ctx, "BAD_REQUEST", err.Error(), http.StatusBadRequest)
		return
	}

	dto := dtos.FindAllInDayFromProviderDTO{
		ProviderID: fmt.Sprintf("%v", id),
		Day:        intDay,
		Month:      intMonth,
		Year:       intYear,
	}

	resp, err := providerAppointments.providerAppointments.Execute(&dto)
	if err != nil {
		logger.Log.Error(err.Error())
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, resp)
}
