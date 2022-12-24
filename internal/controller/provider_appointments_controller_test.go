package controller

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/lucasd-coder/backend-gobarber-golang/internal/dtos"
	"github.com/lucasd-coder/backend-gobarber-golang/internal/infra/errs"
	"github.com/lucasd-coder/backend-gobarber-golang/internal/middlewares"
	"github.com/lucasd-coder/backend-gobarber-golang/internal/mock"
	"github.com/lucasd-coder/backend-gobarber-golang/internal/model"
	"github.com/lucasd-coder/backend-gobarber-golang/internal/util"
	"github.com/stretchr/testify/assert"
	mockTestify "github.com/stretchr/testify/mock"
)

const (
	ID        = "b5e672e0-55f8-4dc5-bd32-757289eedd3b"
	IDInvalid = "Id invalid."
)

func SetUpRouterProviderAppointments() *gin.Engine {
	router := gin.Default()
	router.Use(middlewares.JSONAppErrorReporter())
	return router
}

func SetUpSetIdProviderAppointments(id string) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("id", id)

		c.Next()
	}
}

func mainProviderAppointments() {
	r := SetUpRouterProviderAppointments()
	r.Run(":8080")
}

func buildFindAllInDayFromProviderDTO() *dtos.FindAllInDayFromProviderDTO {
	return &dtos.FindAllInDayFromProviderDTO{
		ProviderID: "b4981a46-fe70-48be-8c98-f2250ccd1bdc",
		Day:        10,
		Month:      10,
		Year:       2022,
	}
}

func buildAppointments() []*model.Appointment {
	return []*model.Appointment{
		{
			ID:         uuid.New(),
			UserID:     "b4981a46-fe70-48be-8c98-f2250ccd1bdc",
			ProviderID: "b8b6683d-1ca8-4984-8160-08f2e96d4ff7",
			Date:       time.Now(),
			CreatedAt:  time.Now(),
			UpdatedAt:  time.Now(),
		}, {
			ID:         uuid.New(),
			UserID:     "e2d3f19e-0f19-47c5-9b2a-a70c6fb097b2",
			ProviderID: "b665ed26-8ae4-407f-a667-9bb093431caf",
			Date:       time.Now(),
			CreatedAt:  time.Now(),
			UpdatedAt:  time.Now(),
		},
	}
}

func TestProviderAppointmentsListProvidersSuccessfully(t *testing.T) {
	listProviderAppointmentsService := new(mock.MockListProviderAppointmentsService)

	listProviderAppointmentsService.On("Execute", mockTestify.Anything).Return(buildAppointments(), nil)

	testController := ProviderAppointments{
		listProviderAppointmentsService,
	}

	r := SetUpRouterProviderAppointments()
	r.Use(SetUpSetIdProviderAppointments(ID))
	query := fmt.Sprintf("?day=%d&month=%d&year=%d", 10, 10, 2022)
	r.GET("/", testController.ListProviders)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/"+query, nil)

	r.ServeHTTP(w, req)

	resp := &[]model.Appointment{}

	util.ParseFromHttpResponse(w.Result(), resp)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotEmpty(t, resp)
}

func TestFailProviderAppointmentsListProviderRequiredQueryFieldDay(t *testing.T) {
	listProviderAppointmentsService := new(mock.MockListProviderAppointmentsService)

	listProviderAppointmentsService.On("Execute", mockTestify.Anything).Return(buildAppointments(), nil)

	testController := ProviderAppointments{
		listProviderAppointmentsService,
	}

	r := SetUpRouterProviderAppointments()
	r.Use(SetUpSetIdProviderAppointments(ID))
	query := fmt.Sprintf("?month=%d&year=%d", 10, 2022)
	r.GET("/", testController.ListProviders)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/"+query, nil)

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestFailProviderAppointmentsListProviderRequiredQueryFieldMonth(t *testing.T) {
	listProviderAppointmentsService := new(mock.MockListProviderAppointmentsService)

	listProviderAppointmentsService.On("Execute", mockTestify.Anything).Return(buildAppointments(), nil)

	testController := ProviderAppointments{
		listProviderAppointmentsService,
	}

	r := SetUpRouterProviderAppointments()
	r.Use(SetUpSetIdProviderAppointments(ID))
	query := fmt.Sprintf("?day=%d&year=%d", 10, 2022)
	r.GET("/", testController.ListProviders)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/"+query, nil)

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestFailProviderAppointmentsListProviderRequiredQueryFieldYear(t *testing.T) {
	listProviderAppointmentsService := new(mock.MockListProviderAppointmentsService)

	listProviderAppointmentsService.On("Execute", mockTestify.Anything).Return(buildAppointments(), nil)

	testController := ProviderAppointments{
		listProviderAppointmentsService,
	}

	r := SetUpRouterProviderAppointments()
	r.Use(SetUpSetIdProviderAppointments(ID))
	query := fmt.Sprintf("?day=%d&month=%d", 10, 10)
	r.GET("/", testController.ListProviders)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/"+query, nil)

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestProviderAppointmentsListProviderInvalidID(t *testing.T) {
	listProviderAppointmentsService := new(mock.MockListProviderAppointmentsService)

	errInvalidId := &errs.AppError{
		Message: "Id invalid.",
		Code:    400,
	}

	listProviderAppointmentsService.On("Execute", mockTestify.Anything).Return(nil, errInvalidId)

	testController := ProviderAppointments{
		listProviderAppointmentsService,
	}

	r := SetUpRouterProviderAppointments()
	r.Use(SetUpSetIdProviderAppointments(IDInvalid))
	query := fmt.Sprintf("?day=%d&month=%d&year=%d", 10, 10, 2022)
	r.GET("/", testController.ListProviders)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/"+query, nil)

	r.ServeHTTP(w, req)

	errDecode := &errs.AppError{}

	util.ParseFromHttpResponse(w.Result(), errDecode)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, errDecode.Code, w.Code)
	assert.Equal(t, errDecode.Message, errInvalidId.Message)
}
