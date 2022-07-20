package users_cache

import (
	"go-api/cmd/core/ports"
	cache_client "go-api/cmd/infra/repositories/cache"
)

type usersCache struct {
	client cache_client.CacheClient
}

func New(cli cache_client.CacheClient) ports.UsersCache {
	return &usersCache{
		client: cli,
	}
}
