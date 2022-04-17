package test

import (
	"backend-gobarber-golang/internal/model"

	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
)

type MockUserTokenRepository struct {
	mock.Mock
}

func (mock *MockUserTokenRepository) Generate(userToken *model.UserToken) *model.UserToken {
	uuidGenerate, _ := uuid.NewRandom()
	return &model.UserToken{ID: uuidGenerate, Token: uuidGenerate}
}

func (mock *MockUserTokenRepository) FindByToken(token string) *model.UserToken {
	return &model.UserToken{}
}
