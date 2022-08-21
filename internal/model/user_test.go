package model_test

import (
	"testing"

	"github.com/lucasd-coder/backend-gobarber-golang/internal/model"

	"github.com/stretchr/testify/require"
)

func TestModel_NewUser(t *testing.T) {
	name := "lucas"
	email := "lucas@gmail.com"
	password := "123456"

	user := model.NewUser(name, email, password)

	require.Equal(t, user.Name, name)
	require.Equal(t, user.Email, email)

	user.Email = ""
	require.NotEmpty(t, user)
}
