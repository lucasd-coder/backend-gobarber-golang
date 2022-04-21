package model_test

import (
	"testing"

	"backend-gobarber-golang/internal/model"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestModel_NewUserToken(t *testing.T) {
	var id uuid.UUID = uuid.MustParse("0399e631-e2f0-4df5-b1d0-ca6d567a318c")

	userToken := model.NewUserToken(id.String())

	require.Equal(t, userToken.UserID, id.String())
}
