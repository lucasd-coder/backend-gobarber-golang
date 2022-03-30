package middlewares

import (
	"net/http"

	"backend-gobarber-golang/internal/service"
	"backend-gobarber-golang/pkg/logger"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func EnsureAuthenticated() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if authHeader == "" {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		const BEARER_SCHEMA = "Bearer "

		tokenString := authHeader[len(BEARER_SCHEMA):]

		token, err := service.NewJWTService().ValidateToken(tokenString)
		if token.Valid {
			claims := token.Claims.(jwt.MapClaims)
			id := claims["sum"]

			c.Set("id", id)

			c.Next()
		} else {
			logger.Log.Error(err)
			c.AbortWithStatus(http.StatusUnauthorized)
		}
	}
}
