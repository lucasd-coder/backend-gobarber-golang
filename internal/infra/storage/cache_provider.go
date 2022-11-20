package storage

import (
	"context"
	"encoding/json"
	"fmt"
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
	val, err := json.Marshal(value)
	if err != nil {
		return err
	}

	r := cacheProvider.Client.Set(ctx, key, val, ttl)
	_, err = r.Result()

	if err != nil {
		return err
	}

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
	iter := cacheProvider.Client.Scan(ctx, 0, fmt.Sprintf("%s:*", prefix), 0).Iterator()

	pipe := cacheProvider.Client.Pipeline()

	for iter.Next(ctx) {
		pipe.Del(ctx, iter.Val())

		if pipe.Len() < 100 {
			continue
		}

		if _, err := pipe.Exec(ctx); err != nil {
			return err
		}
	}

	if _, err := pipe.Exec(ctx); err != nil {
		return err
	}

	return nil
}
