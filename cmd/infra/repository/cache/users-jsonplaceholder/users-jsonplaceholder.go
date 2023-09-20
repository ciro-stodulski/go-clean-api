package usersjsonplaceholdercache

import (
	"encoding/json"
	response_jsonplaceholder "go-clean-api/cmd/domain/dto"
	domainusersjsonplaceholdercache "go-clean-api/cmd/domain/repository/cache"
	cache_client "go-clean-api/cmd/infra/repository/cache"
)

type (
	usersJsonplaceholderCache struct {
		client cache_client.CacheClient
	}
)

func New(cli cache_client.CacheClient) domainusersjsonplaceholdercache.UsersJsonPlaceholderCache {
	return &usersJsonplaceholderCache{
		client: cli,
	}
}

func (uc *usersJsonplaceholderCache) Get(key string) ([]response_jsonplaceholder.User, error) {
	val, err := uc.client.Get(key)

	if err != nil {
		panic(err)
	}

	if val == "" {
		return []response_jsonplaceholder.User{}, nil
	}

	var users []response_jsonplaceholder.User
	err = json.Unmarshal([]byte(val.(string)), &users)

	if err != nil {
		panic(err)
	}

	return users, nil
}

func (uc *usersJsonplaceholderCache) Set(key string, value []response_jsonplaceholder.User, timeEx int) {
	out, err := json.Marshal(value)

	if err != nil {
		panic(err)
	}

	err_client := uc.client.Set(key, string(out), timeEx)

	if err_client != nil {
		panic(err)
	}
}
