package storage_test

import (
	"context"
	"testing"
	"time"

	"github.com/alicebob/miniredis/v2"
	"github.com/go-redis/redis/v8"
	"github.com/lucasd-coder/backend-gobarber-golang/internal/infra/storage"
	"github.com/stretchr/testify/assert"
)

var redisServer *miniredis.Miniredis

func mockRedis() *miniredis.Miniredis {
	s, err := miniredis.Run()
	if err != nil {
		panic(err)
	}

	return s
}

func TestCacheSave(t *testing.T) {
	redisServer = mockRedis()
	redisClient := redis.NewClient(&redis.Options{
		Addr: redisServer.Addr(),
	})

	defer teardown()

	testCacheProvider := storage.CacheProvider{redisClient}

	err := testCacheProvider.Save(context.TODO(), "data", "something here", time.Minute)

	assert.Nil(t, err)
}

func TestCacheRecover(t *testing.T) {
	redisServer = mockRedis()
	redisClient := redis.NewClient(&redis.Options{
		Addr: redisServer.Addr(),
	})

	defer teardown()

	redisClient.Set(context.TODO(), "data", "something here", time.Minute)

	testCacheProvider := storage.CacheProvider{redisClient}

	result, err := testCacheProvider.Recover(context.TODO(), "data")

	assert.Nil(t, err)
	assert.Equal(t, result, "something here")
}

func TestCacheInvalidate(t *testing.T) {
	redisServer = mockRedis()
	redisClient := redis.NewClient(&redis.Options{
		Addr: redisServer.Addr(),
	})

	defer teardown()

	redisClient.Set(context.TODO(), "data", "something here", time.Minute)

	testCacheProvider := storage.CacheProvider{redisClient}

	err := testCacheProvider.Invalidate(context.TODO(), "data")

	assert.Nil(t, err)
}

func TestCacheInvalidatePrefix(t *testing.T) {
	redisServer = mockRedis()
	redisClient := redis.NewClient(&redis.Options{
		Addr: redisServer.Addr(),
	})

	defer teardown()

	redisClient.Set(context.TODO(), "data:prefix:delete", "something here", time.Minute)

	redisClient.Set(context.TODO(), "data:prefix:delete", "something here 2", time.Minute)

	testCacheProvider := storage.CacheProvider{redisClient}

	err := testCacheProvider.InvalidatePrefix(context.TODO(), "data")

	assert.Nil(t, err)
}

func teardown() {
	redisServer.Close()
}
