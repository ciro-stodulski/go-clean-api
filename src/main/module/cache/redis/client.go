package redis

import (
	"context"
	cache_client "go-api/src/main/module/cache"
	"log"
	"os"
	"time"

	"github.com/go-redis/redis/v8"
)

type redisClient struct {
	redis *redis.Client
	ctx   context.Context
}

func New() cache_client.CacheClient {
	var ctx = context.Background()

	rdb := redis.NewClient(&redis.Options{
		Addr: os.Getenv("REDIS_HOST") + ":" + os.Getenv("REDIS_PORT"),
	})

	err := rdb.Set(ctx, "key", "value", 0).Err()

	if err != nil {
		log.Default().Print("Redis: Connection error")

		panic(err)
	}

	log.Default().Print("Redis: Connection with sucessfully")
	return &redisClient{redis: rdb, ctx: ctx}
}

func (rc *redisClient) Get(key string) (string, error) {
	result, err := rc.redis.Get(rc.ctx, key).Result()
	if err == redis.Nil {
		return "", nil
	}

	return result, nil
}

func (rc *redisClient) Set(key string, value string, timeEx int) error {
	return rc.redis.Set(rc.ctx, key, value, time.Duration(time.Duration(timeEx).Milliseconds())).Err()
}
