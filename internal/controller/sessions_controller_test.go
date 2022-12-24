package controller

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/lucasd-coder/backend-gobarber-golang/internal/dtos"
	"github.com/lucasd-coder/backend-gobarber-golang/internal/infra/errs"
	"github.com/lucasd-coder/backend-gobarber-golang/internal/middlewares"
	"github.com/lucasd-coder/backend-gobarber-golang/internal/mock"
	"github.com/lucasd-coder/backend-gobarber-golang/internal/util"
	"github.com/stretchr/testify/assert"
)

var resp = dtos.ResponseProfileDTO{
	ID:        "b665ed26-8ae4-407f-a667-9bb093431caf",
	Name:      "Lindalva",
	Email:     "Lindalva@gmail.com",
	CreatedAt: time.Now(),
}

var token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"

func SetUpRouterSessionsController() *gin.Engine {
	router := gin.Default()
	router.Use(middlewares.JSONAppErrorReporter())
	return router
}

func SetUpSetIdSessionsController(id string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Set("id", id)

		ctx.Next()
	}
}

func mainSessionsController() {
	r := SetUpRouterSessionsController()
	r.Run(":8080")
}

func buildResponseUserAuthenticatedSuccessDTO() *dtos.ResponseUserAuthenticatedSuccessDTO {
	return &dtos.ResponseUserAuthenticatedSuccessDTO{}
}

func TestSessionsControllerSuccessfully(t *testing.T) {
	authenticateUserService := new(mock.MockAuthenticateUserService)

	dto := &dtos.Credentials{
		Email:    "lucas@gmail.com",
		Password: "123456",
	}

	responseUserAuthenticatedSuccessDTO := &dtos.ResponseUserAuthenticatedSuccessDTO{
		Response: resp,
		Token:    token,
	}

	authenticateUserService.On("Execute", dto).Return(responseUserAuthenticatedSuccessDTO, nil)

	testController := SessionsController{
		authenticateUserService,
	}

	jsonValue, _ := json.Marshal(dto)
	r := SetUpRouterSessionsController()
	r.POST("/", testController.AuthenticateUser)
	req, _ := http.NewRequest("POST", "/", bytes.NewBuffer(jsonValue))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	response := &dtos.ResponseUserAuthenticatedSuccessDTO{}

	util.ParseFromHttpResponse(w.Result(), response)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotEmpty(t, response)
}

func TestFailSessionsControllerRequiredFields(t *testing.T) {
	authenticateUserService := new(mock.MockAuthenticateUserService)

	dto := &dtos.Credentials{}

	authenticateUserService.On("Execute", dto).Return(nil, nil)

	testController := SessionsController{
		authenticateUserService,
	}

	jsonValue, _ := json.Marshal(dto)
	r := SetUpRouterSessionsController()
	r.POST("/", testController.AuthenticateUser)
	req, _ := http.NewRequest("POST", "/", bytes.NewBuffer(jsonValue))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestFailSessionsControllerUserUnauthorized(t *testing.T) {
	authenticateUserService := new(mock.MockAuthenticateUserService)

	dto := &dtos.Credentials{
		Email:    "lucas@gmail.com",
		Password: "123456",
	}

	responseUserAuthenticatedSuccessDTO := &dtos.ResponseUserAuthenticatedSuccessDTO{
		Response: resp,
		Token:    token,
	}

	errAuthorization := &errs.AppError{
		Message: "Incorrect email/password combination.",
		Code:    401,
	}

	authenticateUserService.On("Execute", dto).Return(responseUserAuthenticatedSuccessDTO, errAuthorization)

	testController := SessionsController{
		authenticateUserService,
	}

	jsonValue, _ := json.Marshal(dto)
	r := SetUpRouterSessionsController()
	r.POST("/", testController.AuthenticateUser)
	req, _ := http.NewRequest("POST", "/", bytes.NewBuffer(jsonValue))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	errDecode := &errs.AppError{}

	util.ParseFromHttpResponse(w.Result(), errDecode)

	assert.Equal(t, http.StatusUnauthorized, w.Code)
	assert.Equal(t, errDecode.Code, w.Code)
	assert.Equal(t, errDecode.Message, errAuthorization.Message)
}
