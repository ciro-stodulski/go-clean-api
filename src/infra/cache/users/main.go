package users_cache

import (
	"go-api/src/core/ports"
	cache_client "go-api/src/infra/cache"
)

type usersCache struct {
	client cache_client.CacheClient
}

func New(cli cache_client.CacheClient) ports.UsersCache {
	return &usersCache{
		client: cli,
	}
}
