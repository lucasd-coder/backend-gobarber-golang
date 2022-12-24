package controller

import (
	"bytes"
	"encoding/json"
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

func SetUpRouterAppointmentsController() *gin.Engine {
	router := gin.Default()
	router.Use(middlewares.JSONAppErrorReporter())
	return router
}

func SetUpSetIdAppointmentsController(id string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Set("id", id)

		ctx.Next()
	}
}

func mainAppointmentsController() {
	r := SetUpRouterAppointmentsController()
	r.Run(":8080")
}

func buildAppointmentDTO() *dtos.AppointmentDTO {
	return &dtos.AppointmentDTO{
		Date:       time.Now(),
		ProviderID: "88980f09-e23e-45bc-a3be-e5471ba41c35",
	}
}

func buildAppointment() *model.Appointment {
	return &model.Appointment{
		ID:         uuid.New(),
		UserID:     "b4981a46-fe70-48be-8c98-f2250ccd1bdc",
		ProviderID: "b8b6683d-1ca8-4984-8160-08f2e96d4ff7",
		Date:       time.Now(),
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}
}

func TestAppointmentsControllerSuccessfully(t *testing.T) {
	id := "b5e672e0-55f8-4dc5-bd32-757289eedd3b"

	createAppointment := new(mock.MockCreateAppointmentService)

	createAppointment.On("Execute", id, mockTestify.Anything).Return(buildAppointment(), nil)

	testController := AppointmentsController{
		createAppointment,
	}

	jsonValue, _ := json.Marshal(buildAppointmentDTO())
	r := SetUpRouterAppointmentsController()
	r.Use(SetUpSetIdAppointmentsController(id))
	r.POST("/", testController.CreateAppointment)
	req, _ := http.NewRequest("POST", "/", bytes.NewBuffer(jsonValue))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	resp := &model.Appointment{}

	util.ParseFromHttpResponse(w.Result(), resp)

	assert.Equal(t, http.StatusCreated, w.Code)
	assert.NotEmpty(t, resp)
}

func TestFailAppointmentsControllerInvalidID(t *testing.T) {
	idInvalid := "Id invalid."

	errInvalidId := &errs.AppError{
		Message: "Id invalid.",
		Code:    400,
	}

	createAppointment := new(mock.MockCreateAppointmentService)

	createAppointment.On("Execute", idInvalid, mockTestify.Anything).Return(nil, errInvalidId)

	testController := AppointmentsController{
		createAppointment,
	}

	jsonValue, _ := json.Marshal(buildAppointmentDTO())
	r := SetUpRouterAppointmentsController()
	r.Use(SetUpSetIdAppointmentsController(idInvalid))
	r.POST("/", testController.CreateAppointment)
	req, _ := http.NewRequest("POST", "/", bytes.NewBuffer(jsonValue))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	errDecode := &errs.AppError{}

	util.ParseFromHttpResponse(w.Result(), errDecode)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, errDecode.Code, w.Code)
	assert.Equal(t, errDecode.Message, errInvalidId.Message)
}

func TestFailAppointmentsControllerRequiredFields(t *testing.T) {
	id := "b5e672e0-55f8-4dc5-bd32-757289eedd3b"

	createAppointment := new(mock.MockCreateAppointmentService)

	createAppointment.On("Execute", id, mockTestify.Anything).Return(nil, nil)

	testController := AppointmentsController{
		createAppointment,
	}

	dto := dtos.AppointmentDTO{}

	jsonValue, _ := json.Marshal(dto)
	r := SetUpRouterAppointmentsController()
	r.Use(SetUpSetIdAppointmentsController(id))
	r.POST("/", testController.CreateAppointment)
	req, _ := http.NewRequest("POST", "/", bytes.NewBuffer(jsonValue))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}
