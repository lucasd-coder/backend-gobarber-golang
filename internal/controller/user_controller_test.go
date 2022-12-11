package controller

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/lucasd-coder/backend-gobarber-golang/internal/dtos"
	"github.com/lucasd-coder/backend-gobarber-golang/internal/infra/errs"
	"github.com/lucasd-coder/backend-gobarber-golang/internal/middlewares"
	"github.com/lucasd-coder/backend-gobarber-golang/internal/mock"
	"github.com/stretchr/testify/assert"
)

func SetUpRouter() *gin.Engine {
	router := gin.Default()
	router.Use(middlewares.JSONAppErrorReporter())
	return router
}

func main() {
	r := SetUpRouter()
	r.Run(":8080")
}

func builderResponseProfileDTO() *dtos.ResponseProfileDTO {
	return &dtos.ResponseProfileDTO{
		ID:        "b5e672e0-55f8-4dc5-bd32-757289eedd3b",
		Name:      "Maria",
		Email:     "Maria@gmail.com",
		Avatar:    "",
		CreatedAt: time.Now(),
	}
}

func builderResponseCreateUserDTO() *dtos.ResponseCreateUserDTO {
	return &dtos.ResponseCreateUserDTO{
		Name:  "Maria",
		Email: "maria@gmail.com",
	}
}

func builderUserDTO() *dtos.UserDTO {
	return &dtos.UserDTO{
		Name:     "Maria",
		Email:    "maria@gmail.com",
		Password: "123456",
	}
}

func TestCreateUserSuccessfully(t *testing.T) {
	createUsersService := new(mock.MockCreateUsersService)
	showProfileService := new(mock.MockShowProfileService)
	updateProfileService := new(mock.MockUpdateProfileService)
	updateUserAvatarService := new(mock.MockUpdateUserAvatarService)

	createUsersService.On("Execute", builderUserDTO()).Return(builderResponseCreateUserDTO(), nil)

	testController := UserController{
		createUsersService, showProfileService,
		updateProfileService, updateUserAvatarService,
	}

	jsonValue, _ := json.Marshal(builderUserDTO())
	r := SetUpRouter()
	r.POST("/", testController.CreateUser)
	req, _ := http.NewRequest("POST", "/", bytes.NewBuffer(jsonValue))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
}

func TestFailCreateUserRequiredFields(t *testing.T) {
	createUsersService := new(mock.MockCreateUsersService)
	showProfileService := new(mock.MockShowProfileService)
	updateProfileService := new(mock.MockUpdateProfileService)
	updateUserAvatarService := new(mock.MockUpdateUserAvatarService)

	user := &dtos.UserDTO{
		Name:     "Maria",
		Email:    "email invalid",
		Password: "123456",
	}

	createUsersService.On("Execute", user).Return(nil, nil)

	testController := UserController{
		createUsersService, showProfileService,
		updateProfileService, updateUserAvatarService,
	}

	jsonValue, _ := json.Marshal(user)
	r := SetUpRouter()
	r.POST("/", testController.CreateUser)
	req, _ := http.NewRequest("POST", "/", bytes.NewBuffer(jsonValue))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestFailCreateUseEmailAddressAlready(t *testing.T) {
	createUsersService := new(mock.MockCreateUsersService)
	showProfileService := new(mock.MockShowProfileService)
	updateProfileService := new(mock.MockUpdateProfileService)
	updateUserAvatarService := new(mock.MockUpdateUserAvatarService)

	errInvalidId := &errs.AppError{
		Message: "Email address already used by another",
		Code:    400,
	}

	createUsersService.On("Execute", builderUserDTO()).Return(nil, errInvalidId)

	testController := UserController{
		createUsersService, showProfileService,
		updateProfileService, updateUserAvatarService,
	}

	jsonValue, _ := json.Marshal(builderUserDTO())
	r := SetUpRouter()
	r.POST("/", testController.CreateUser)
	req, _ := http.NewRequest("POST", "/", bytes.NewBuffer(jsonValue))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	errDecode := &errs.AppError{}

	ParseFromHttpResponse(w.Result(), errDecode)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, errDecode.Code, w.Code)
	assert.Equal(t, errDecode.Message, errInvalidId.Message)
}

func TestShowProfileSuccessfully(t *testing.T) {
	createUsersService := new(mock.MockCreateUsersService)
	showProfileService := new(mock.MockShowProfileService)
	updateProfileService := new(mock.MockUpdateProfileService)
	updateUserAvatarService := new(mock.MockUpdateUserAvatarService)

	id := "b5e672e0-55f8-4dc5-bd32-757289eedd3b"

	showProfileService.On("Execute", id).Return(builderResponseProfileDTO(), nil)

	testController := UserController{
		createUsersService, showProfileService,
		updateProfileService, updateUserAvatarService,
	}

	r := SetUpRouter()
	r.Use(SetUpSetId(id))
	r.GET("/", testController.ShowProfile)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestShowProfileInvalidId(t *testing.T) {
	createUsersService := new(mock.MockCreateUsersService)
	showProfileService := new(mock.MockShowProfileService)
	updateProfileService := new(mock.MockUpdateProfileService)
	updateUserAvatarService := new(mock.MockUpdateUserAvatarService)

	id := "invalid id"

	errInvalidId := &errs.AppError{
		Message: "Id invalid.",
		Code:    400,
	}

	showProfileService.On("Execute", id).Return(builderResponseProfileDTO(), errInvalidId)

	testController := UserController{
		createUsersService, showProfileService,
		updateProfileService, updateUserAvatarService,
	}

	r := SetUpRouter()
	r.Use(SetUpSetId(id))
	r.GET("/", testController.ShowProfile)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)

	r.ServeHTTP(w, req)

	errDecode := &errs.AppError{}

	ParseFromHttpResponse(w.Result(), errDecode)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, errDecode.Code, w.Code)
	assert.Equal(t, errDecode.Message, errInvalidId.Message)
}

func SetUpSetId(id string) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("id", id)

		c.Next()
	}
}

func ParseFromHttpResponse(resp *http.Response, model interface{}) error {
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	err = json.Unmarshal(bodyBytes, &model)

	if err != nil {
		return err
	}

	return nil
}
