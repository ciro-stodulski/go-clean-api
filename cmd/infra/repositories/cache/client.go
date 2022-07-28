package cacheclient

import (
	"context"
	"go-api/cmd/shared/env"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
)

type (
	CacheClient interface {
		Set(key string, value string, timeEx int) error
		Get(key string) (string, error)
	}
	redisClient struct {
		redis *redis.Client
		ctx   context.Context
	}
)

func New() CacheClient {
	var ctx = context.Background()

	rdb := redis.NewClient(&redis.Options{
		Addr: env.Env().RedisHost + ":" + env.Env().RedisPort,
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
