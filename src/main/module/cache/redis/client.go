package redis

import (
	"context"
	cache_client "go-api/src/main/module/cache"
	"time"

	"github.com/go-redis/redis/v8"
)

type redisClient struct {
	redis *redis.Client
}

var ctx = context.Background()

func New() cache_client.CacheClient {
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	return &redisClient{redis: rdb}
}

func (rc *redisClient) Get(key string) (string, error) {
	return rc.redis.Get(ctx, key).Result()
}

func (rc *redisClient) Set(key string, value string, timeEx int) error {
	return rc.redis.Set(ctx, key, value, time.Duration(time.Duration(timeEx).Milliseconds())).Err()
}
