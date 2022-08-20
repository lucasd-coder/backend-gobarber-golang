package storage

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

type CacheProvider struct {
	Client *redis.Client
}

func NewCacheProvider(redisClient *redis.Client) *CacheProvider {
	return &CacheProvider{
		Client: redisClient,
	}
}

func (cacheProvider *CacheProvider) Save(ctx context.Context, key string, value interface{}, ttl time.Duration) error {
	return nil
}

func (cacheProvider *CacheProvider) Recover(ctx context.Context, key string) (string, error) {
	return cacheProvider.Client.Get(ctx, key).Result()
}

func (cacheProvider *CacheProvider) Invalidate(ctx context.Context, key string) error {
	err := cacheProvider.Client.Del(ctx, key).Err()
	if err != nil {
		return err
	}
	return nil
}

func (cacheProvider *CacheProvider) InvalidatePrefix(ctx context.Context, prefix string) error {
	return nil
}
