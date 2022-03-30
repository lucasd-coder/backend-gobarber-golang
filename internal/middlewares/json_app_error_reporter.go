package middlewares

import (
	"net/http"

	"backend-gobarber-golang/internal/infra/errs"

	"github.com/gin-gonic/gin"
)

//
// Middleware Error Handler in server package
//
func JSONAppErrorReporter() gin.HandlerFunc {
	return jsonAppErrorReporterT(gin.ErrorTypeAny)
}

func jsonAppErrorReporterT(errType gin.ErrorType) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		detectedErrors := c.Errors.ByType(errType)

		if len(detectedErrors) > 0 {
			err := detectedErrors[0].Err
			var parsedError *errs.AppError
			switch err.(type) {
			case *errs.AppError:
				parsedError = err.(*errs.AppError)
			default:
				parsedError = &errs.AppError{
					Code:    http.StatusInternalServerError,
					Message: "Internal Server Error",
				}
			}
			// Put the error into response
			c.IndentedJSON(parsedError.Code, parsedError)
			c.Abort()
			// or c.AbortWithStatusJSON(parsedError.Code, parsedError)
			return
		}
	}
}
