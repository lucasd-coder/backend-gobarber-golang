package test

import (
	"context"
	"time"

	"github.com/stretchr/testify/mock"
)

type MockCacheProvider struct {
	mock.Mock
}

func (mock *MockCacheProvider) Save(ctx context.Context, key string, value interface{}, ttl time.Duration) error {
	args := mock.Called(ctx, key, value)
	return args.Error(0)
}

func (mock *MockCacheProvider) Recover(ctx context.Context, key string) interface{} {
	args := mock.Called(ctx, key)
	result := args.Get(0)
	return result
}

func (mock *MockCacheProvider) Invalidate(ctx context.Context, key string) error {
	args := mock.Called(ctx, key)
	return args.Error(0)
}

func (mock *MockCacheProvider) InvalidatePrefix(ctx context.Context, prefix string) error {
	args := mock.Called(ctx, prefix)
	return args.Error(0)
}
