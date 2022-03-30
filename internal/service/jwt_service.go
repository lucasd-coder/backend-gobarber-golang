package service

import (
	"fmt"
	"time"

	"backend-gobarber-golang/config"

	"github.com/golang-jwt/jwt/v4"
)

type JWTService struct {
	secretKey string
	issuer    string
}

func NewJWTService() *JWTService {
	cfg := config.GetConfig()
	return &JWTService{
		secretKey: cfg.Secret,
		issuer:    cfg.Issuer,
	}
}

type jwtCustomClaims struct {
	Sum string `json:"sum"`
	jwt.StandardClaims
}

func (service *JWTService) GenerateToken(id string) string {
	claims := &jwtCustomClaims{
		id,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 2).Unix(),
			Issuer:    service.issuer,
			IssuedAt:  time.Now().Unix(),
		},
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token using the secret signing key
	t, err := token.SignedString([]byte(service.secretKey))
	if err != nil {
		panic(err)
	}
	return t
}

func (service *JWTService) ValidateToken(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Signing method validation
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		// Return the secret signing key
		return []byte(service.secretKey), nil
	})
}
