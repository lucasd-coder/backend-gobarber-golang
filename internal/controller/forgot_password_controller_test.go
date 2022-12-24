package controller

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/lucasd-coder/backend-gobarber-golang/internal/dtos"
	"github.com/lucasd-coder/backend-gobarber-golang/internal/infra/errs"
	"github.com/lucasd-coder/backend-gobarber-golang/internal/middlewares"
	"github.com/lucasd-coder/backend-gobarber-golang/internal/mock"
	"github.com/lucasd-coder/backend-gobarber-golang/internal/util"
	"github.com/stretchr/testify/assert"
)

func SetUpRouterForgotPasswordController() *gin.Engine {
	router := gin.Default()
	router.Use(middlewares.JSONAppErrorReporter())

	return router
}

func SetUpSetIdForgotPasswordController(id string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Set("id", id)

		ctx.Next()
	}
}

func mainForgotPasswordController() {
	r := SetUpRouterForgotPasswordController()
	r.Run(":8080")
}

func TestForgotPasswordControllerSuccessfully(t *testing.T) {
	sendForgotPasswordEmailService := new(mock.MockSendForgotPasswordEmailService)

	dto := dtos.ForgotPasswordEmail{
		Email: "lucas@gmail.com",
	}

	sendForgotPasswordEmailService.On("Execute", &dto).Return(nil)

	testController := ForgotPasswordController{
		sendForgotPasswordEmailService,
	}

	jsonValue, _ := json.Marshal(dto)
	r := SetUpRouterForgotPasswordController()
	r.POST("/", testController.ForgotPassword)
	req, _ := http.NewRequest("POST", "/", bytes.NewBuffer(jsonValue))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusAccepted, w.Code)
}

func TestFailForgotPasswordControllerRequiredFields(t *testing.T) {
	sendForgotPasswordEmailService := new(mock.MockSendForgotPasswordEmailService)

	dto := dtos.ForgotPasswordEmail{
		Email: "lucas@gmail.com",
	}

	errUserNotFound := &errs.AppError{
		Message: "User not found.",
		Code:    404,
	}

	sendForgotPasswordEmailService.On("Execute", &dto).Return(errUserNotFound)

	testController := ForgotPasswordController{
		sendForgotPasswordEmailService,
	}

	jsonValue, _ := json.Marshal(dto)
	r := SetUpRouterForgotPasswordController()
	r.POST("/", testController.ForgotPassword)
	req, _ := http.NewRequest("POST", "/", bytes.NewBuffer(jsonValue))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	errDecode := &errs.AppError{}

	util.ParseFromHttpResponse(w.Result(), errDecode)

	assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestFailForgotPasswordControllerUserNotFound(t *testing.T) {
	sendForgotPasswordEmailService := new(mock.MockSendForgotPasswordEmailService)

	dto := dtos.ForgotPasswordEmail{}

	sendForgotPasswordEmailService.On("Execute", &dto).Return(nil)

	testController := ForgotPasswordController{
		sendForgotPasswordEmailService,
	}

	jsonValue, _ := json.Marshal(dto)
	r := SetUpRouterForgotPasswordController()
	r.POST("/", testController.ForgotPassword)
	req, _ := http.NewRequest("POST", "/", bytes.NewBuffer(jsonValue))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}
