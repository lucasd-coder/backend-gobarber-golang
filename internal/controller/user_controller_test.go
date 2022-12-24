package controller

import (
	"bytes"
	"encoding/json"
	"image"
	"image/color"
	"image/png"
	"io"
	"mime/multipart"
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
	mockTestify "github.com/stretchr/testify/mock"
)

var (
	id        = "b5e672e0-55f8-4dc5-bd32-757289eedd3b"
	idInvalid = "Id invalid."
)

func SetUpRouterUserController() *gin.Engine {
	router := gin.Default()
	router.Use(middlewares.JSONAppErrorReporter())
	return router
}

func SetUpSetIdUserController(id string) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("id", id)

		c.Next()
	}
}

func mainUserController() {
	r := SetUpRouterUserController()
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

func buildUpdateUserProfileDTO() *dtos.UpdateUserProfileDTO {
	return &dtos.UpdateUserProfileDTO{
		Name:        "Maria",
		Email:       "maria@gmail.com",
		Password:    "1234567",
		OldPassword: "123456",
	}
}

func createImage() *image.RGBA {
	width := 200
	height := 100

	upLeft := image.Point{0, 0}
	lowRight := image.Point{width, height}

	img := image.NewRGBA(image.Rectangle{upLeft, lowRight})

	// Colors are defined by Red, Green, Blue, Alpha uint8 values.
	cyan := color.RGBA{100, 200, 200, 0xff}

	// Set color for each pixel.
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			switch {
			case x < width/2 && y < height/2: // upper left quadrant
				img.Set(x, y, cyan)
			case x >= width/2 && y >= height/2: // lower right quadrant
				img.Set(x, y, color.White)
			default:
				// Use zero value.
			}
		}
	}

	return img
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
	r := SetUpRouterUserController()
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
	r := SetUpRouterUserController()
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
	r := SetUpRouterUserController()
	r.POST("/", testController.CreateUser)
	req, _ := http.NewRequest("POST", "/", bytes.NewBuffer(jsonValue))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	errDecode := &errs.AppError{}

	util.ParseFromHttpResponse(w.Result(), errDecode)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, errDecode.Code, w.Code)
	assert.Equal(t, errDecode.Message, errInvalidId.Message)
}

func TestShowProfileSuccessfully(t *testing.T) {
	createUsersService := new(mock.MockCreateUsersService)
	showProfileService := new(mock.MockShowProfileService)
	updateProfileService := new(mock.MockUpdateProfileService)
	updateUserAvatarService := new(mock.MockUpdateUserAvatarService)

	showProfileService.On("Execute", id).Return(builderResponseProfileDTO(), nil)

	testController := UserController{
		createUsersService, showProfileService,
		updateProfileService, updateUserAvatarService,
	}

	r := SetUpRouterUserController()
	r.Use(SetUpSetIdUserController(id))
	r.GET("/", testController.ShowProfile)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)

	r.ServeHTTP(w, req)

	resProfileDTO := &dtos.ResponseProfileDTO{}

	util.ParseFromHttpResponse(w.Result(), resProfileDTO)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotEmpty(t, resProfileDTO)
}

func TestShowProfileInvalidID(t *testing.T) {
	createUsersService := new(mock.MockCreateUsersService)
	showProfileService := new(mock.MockShowProfileService)
	updateProfileService := new(mock.MockUpdateProfileService)
	updateUserAvatarService := new(mock.MockUpdateUserAvatarService)

	errInvalidId := &errs.AppError{
		Message: "Id invalid.",
		Code:    400,
	}

	showProfileService.On("Execute", idInvalid).Return(builderResponseProfileDTO(), errInvalidId)

	testController := UserController{
		createUsersService, showProfileService,
		updateProfileService, updateUserAvatarService,
	}

	r := SetUpRouterUserController()
	r.Use(SetUpSetIdUserController(idInvalid))
	r.GET("/", testController.ShowProfile)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)

	r.ServeHTTP(w, req)

	errDecode := &errs.AppError{}

	util.ParseFromHttpResponse(w.Result(), errDecode)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, errDecode.Code, w.Code)
	assert.Equal(t, errDecode.Message, errInvalidId.Message)
}

func TestUpdateProfileSuccessfully(t *testing.T) {
	createUsersService := new(mock.MockCreateUsersService)
	showProfileService := new(mock.MockShowProfileService)
	updateProfileService := new(mock.MockUpdateProfileService)
	updateUserAvatarService := new(mock.MockUpdateUserAvatarService)

	updateProfileService.On("Execute", id, buildUpdateUserProfileDTO()).Return(builderResponseProfileDTO(), nil)

	testController := UserController{
		createUsersService, showProfileService,
		updateProfileService, updateUserAvatarService,
	}

	jsonValue, _ := json.Marshal(buildUpdateUserProfileDTO())
	r := SetUpRouterUserController()
	r.Use(SetUpSetIdUserController(id))
	r.PUT("/", testController.UpdateProfile)
	req, _ := http.NewRequest("PUT", "/", bytes.NewBuffer(jsonValue))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	resProfileDTO := &dtos.ResponseProfileDTO{}

	util.ParseFromHttpResponse(w.Result(), resProfileDTO)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotEmpty(t, resProfileDTO)
}

func TestUpdateProfileInvalidID(t *testing.T) {
	createUsersService := new(mock.MockCreateUsersService)
	showProfileService := new(mock.MockShowProfileService)
	updateProfileService := new(mock.MockUpdateProfileService)
	updateUserAvatarService := new(mock.MockUpdateUserAvatarService)

	errInvalidId := &errs.AppError{
		Message: "Id invalid.",
		Code:    400,
	}

	updateProfileService.On("Execute", idInvalid, buildUpdateUserProfileDTO()).Return(builderResponseProfileDTO(), errInvalidId)

	testController := UserController{
		createUsersService, showProfileService,
		updateProfileService, updateUserAvatarService,
	}

	jsonValue, _ := json.Marshal(buildUpdateUserProfileDTO())
	r := SetUpRouterUserController()
	r.Use(SetUpSetIdUserController(idInvalid))
	r.PUT("/", testController.UpdateProfile)
	req, _ := http.NewRequest("PUT", "/", bytes.NewBuffer(jsonValue))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	errDecode := &errs.AppError{}

	util.ParseFromHttpResponse(w.Result(), errDecode)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, errDecode.Code, w.Code)
	assert.Equal(t, errDecode.Message, errInvalidId.Message)
}

func TestFailUpdateProfileRequiredFields(t *testing.T) {
	createUsersService := new(mock.MockCreateUsersService)
	showProfileService := new(mock.MockShowProfileService)
	updateProfileService := new(mock.MockUpdateProfileService)
	updateUserAvatarService := new(mock.MockUpdateUserAvatarService)

	id := "b5e672e0-55f8-4dc5-bd32-757289eedd3b"

	userProfile := &dtos.UpdateUserProfileDTO{
		Email:       "Email invalid",
		Password:    "1234567",
		OldPassword: "123456",
	}

	updateProfileService.On("Execute", id, userProfile).Return(nil, nil)

	testController := UserController{
		createUsersService, showProfileService,
		updateProfileService, updateUserAvatarService,
	}

	jsonValue, _ := json.Marshal(userProfile)
	r := SetUpRouterUserController()
	r.PUT("/", testController.UpdateProfile)
	req, _ := http.NewRequest("PUT", "/", bytes.NewBuffer(jsonValue))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestUpdateUserAvatarSuccessfully(t *testing.T) {
	createUsersService := new(mock.MockCreateUsersService)
	showProfileService := new(mock.MockShowProfileService)
	updateProfileService := new(mock.MockUpdateProfileService)
	updateUserAvatarService := new(mock.MockUpdateUserAvatarService)

	testController := UserController{
		createUsersService, showProfileService,
		updateProfileService, updateUserAvatarService,
	}

	updateUserAvatarService.On("Execute", id, mockTestify.Anything).Return(builderResponseProfileDTO(), nil)

	r := SetUpRouterUserController()
	r.Use(SetUpSetIdUserController(id))

	pr, pw := io.Pipe()

	writer := multipart.NewWriter(pw)

	go func() {
		defer writer.Close()

		part, err := writer.CreateFormFile("avatar", "avatar.png")
		if err != nil {
			t.Error(err)
		}

		img := createImage()

		err = png.Encode(part, img)
		if err != nil {
			t.Error(err)
		}
	}()

	r.PATCH("/", testController.UpdateUserAvatar)

	req, _ := http.NewRequest("PATCH", "/", pr)

	req.Header.Add("Content-Type", writer.FormDataContentType())
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	resProfileDTO := &dtos.ResponseProfileDTO{}

	util.ParseFromHttpResponse(w.Result(), resProfileDTO)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotEmpty(t, resProfileDTO)
}

func TestUpdateUserAvatarInvalidID(t *testing.T) {
	createUsersService := new(mock.MockCreateUsersService)
	showProfileService := new(mock.MockShowProfileService)
	updateProfileService := new(mock.MockUpdateProfileService)
	updateUserAvatarService := new(mock.MockUpdateUserAvatarService)

	testController := UserController{
		createUsersService, showProfileService,
		updateProfileService, updateUserAvatarService,
	}

	errInvalidId := &errs.AppError{
		Message: "Id invalid.",
		Code:    400,
	}

	updateUserAvatarService.On("Execute", idInvalid, mockTestify.Anything).Return(nil, errInvalidId)

	r := SetUpRouterUserController()
	r.Use(SetUpSetIdUserController(idInvalid))

	pr, pw := io.Pipe()

	writer := multipart.NewWriter(pw)

	go func() {
		defer writer.Close()

		part, err := writer.CreateFormFile("avatar", "avatar.png")
		if err != nil {
			t.Error(err)
		}

		img := createImage()

		err = png.Encode(part, img)
		if err != nil {
			t.Error(err)
		}
	}()

	r.PATCH("/", testController.UpdateUserAvatar)

	req, _ := http.NewRequest("PATCH", "/", pr)

	req.Header.Add("Content-Type", writer.FormDataContentType())
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	errDecode := &errs.AppError{}

	util.ParseFromHttpResponse(w.Result(), errDecode)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, errDecode.Code, w.Code)
	assert.Equal(t, errDecode.Message, errInvalidId.Message)
}

func TestFailUpdateUserAvatarRequiredFields(t *testing.T) {
	createUsersService := new(mock.MockCreateUsersService)
	showProfileService := new(mock.MockShowProfileService)
	updateProfileService := new(mock.MockUpdateProfileService)
	updateUserAvatarService := new(mock.MockUpdateUserAvatarService)

	testController := UserController{
		createUsersService, showProfileService,
		updateProfileService, updateUserAvatarService,
	}

	updateUserAvatarService.On("Execute", id, mockTestify.Anything).Return(builderResponseProfileDTO(), nil)

	r := SetUpRouterUserController()
	r.Use(SetUpSetIdUserController(id))

	r.PATCH("/", testController.UpdateUserAvatar)

	req, _ := http.NewRequest("PATCH", "/", nil)

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}
