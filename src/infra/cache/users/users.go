package users_cache

import (
	"encoding/json"
	response_jsonplaceholder "go-api/src/infra/http/integrations/jsonplaceholder/responses"
)

func (userCache *usersCache) Get(key string) ([]response_jsonplaceholder.User, error) {
	val, err := userCache.client.Get(key)

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

func (userCache *usersCache) Set(key string, value []response_jsonplaceholder.User, timeEx int) {
	out, err := json.Marshal(value)

	if err != nil {
		panic(err)
	}

	err_client := userCache.client.Set(key, string(out), timeEx)

	if err_client != nil {
		panic(err)
	}

}
