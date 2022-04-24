package test

import (
	"time"

	"github.com/stretchr/testify/mock"
)

type MockCacheProvider struct {
	mock.Mock
}

func (mock *MockCacheProvider) Save(key string, value interface{}, ttl time.Duration) error {
	args := mock.Called(key, value)
	return args.Error(0)
}

func (mock *MockCacheProvider) Recover(key string) interface{} {
	args := mock.Called(key)
	result := args.Get(0)
	return result
}

func (mock *MockCacheProvider) Invalidate(key string) error {
	args := mock.Called(key)
	return args.Error(0)
}

func (mock *MockCacheProvider) InvalidatePrefix(prefix string) error {
	args := mock.Called(prefix)
	return args.Error(0)
}
