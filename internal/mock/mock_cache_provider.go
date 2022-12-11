package mock

import (
	"context"
	"fmt"
	"time"

	"github.com/stretchr/testify/mock"
)

type MockCacheProvider struct {
	mock.Mock
}

func (mock *MockCacheProvider) Save(ctx context.Context, key string, value interface{}, ttl time.Duration) error {
	args := mock.Called(ctx, key, value, ttl)
	return args.Error(0)
}

func (mock *MockCacheProvider) Recover(ctx context.Context, key string) (string, error) {
	args := mock.Called(ctx, key)
	result := args.Get(0)
	return fmt.Sprint(result), nil
}

func (mock *MockCacheProvider) Invalidate(ctx context.Context, key string) error {
	args := mock.Called(ctx, key)
	return args.Error(0)
}

func (mock *MockCacheProvider) InvalidatePrefix(ctx context.Context, prefix string) error {
	args := mock.Called(ctx, prefix)
	return args.Error(0)
}
