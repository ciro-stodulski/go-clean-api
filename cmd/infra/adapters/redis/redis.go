package redisadapter

import (
	"context"
	cacheclient "go-clean-api/cmd/infra/repository/cache"
	"go-clean-api/cmd/shared/env"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
)

type (
	RedisAdapter struct {
		redis *redis.Client
		ctx   context.Context
	}
)

func (rc *RedisAdapter) Get(key string) (any, error) {
	result, err := rc.redis.Get(rc.ctx, key).Result()
	if err == redis.Nil {
		// A chave não existe no Redis, você pode lidar com isso de acordo com suas necessidades
		return nil, nil // ou outro valor padrão ou erro personalizado, se necessário
	}
	return result, err
}

func (rc *RedisAdapter) Set(key string, value any, timeEx int) error {
	return rc.redis.Set(rc.ctx, key, value, time.Duration(time.Duration(timeEx).Milliseconds())).Err()
}

func New() cacheclient.CacheClient {
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
	return &RedisAdapter{redis: rdb, ctx: ctx}
}
