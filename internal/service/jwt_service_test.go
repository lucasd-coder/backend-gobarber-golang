package service_test

import (
	"testing"

	"github.com/lucasd-coder/backend-gobarber-golang/internal/service"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestJWTService_GenerateTokenSuccessfully(t *testing.T) {
	var id uuid.UUID = uuid.MustParse("0399e631-e2f0-4df5-b1d0-ca6d567a318c")

	testService := service.JWTService{}

	token := testService.GenerateToken(id.String())

	assert.NotNil(t, token)
}

func TestJWTService_ValidateTokenSuccessfully(t *testing.T) {
	var id uuid.UUID = uuid.MustParse("0399e631-e2f0-4df5-b1d0-ca6d567a318c")

	testService := service.JWTService{}

	token := testService.GenerateToken(id.String())

	resp, err := testService.ValidateToken(token)

	assert.Nil(t, err)
	assert.True(t, resp.Valid)
}

func TestJWTService_ErrorUnauthorizedValidateToken(t *testing.T) {
	testService := service.JWTService{}

	var testToken string = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG" +
		"4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"

	resp, err := testService.ValidateToken(testToken)

	assert.NotNil(t, err)
	assert.False(t, resp.Valid)
}
