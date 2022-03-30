package controller

import "github.com/gin-gonic/gin"

type ErrorResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func HandleError(ctx *gin.Context, code, message string, httpStatus int) {
	ctx.JSON(httpStatus, ErrorResponse{
		Code:    code,
		Message: message,
	})
}
