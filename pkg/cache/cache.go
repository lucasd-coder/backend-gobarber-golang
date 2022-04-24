package cache

import (
	"backend-gobarber-golang/config"

	"github.com/go-redis/redis/v8"
)

var client *redis.Client

func SetUpRedis(cfg *config.Config) {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     cfg.RedisUrl,
		DB:       cfg.RedisDb,
		Password: cfg.RedisPassword,
	})

	client = redisClient
}

func GetClient() *redis.Client {
	return client
}
