package usersjsonplaceholdercache

import (
	"encoding/json"
	response_jsonplaceholder "go-api/cmd/infra/integrations/http/jsonplaceholder/responses"
	cache_client "go-api/cmd/infra/repositories/cache"
)

type (
	UsersJsonPlaceholderCache interface {
		Set(key string, user []response_jsonplaceholder.User, time int)
		Get(key string) ([]response_jsonplaceholder.User, error)
	}
	usersJsonplaceholderCache struct {
		client cache_client.CacheClient
	}
)

func New(cli cache_client.CacheClient) UsersJsonPlaceholderCache {
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
	err = json.Unmarshal([]byte(val), &users)

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
