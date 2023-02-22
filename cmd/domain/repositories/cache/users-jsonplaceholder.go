package domainusersjsonplaceholdercache

import (
	response_jsonplaceholder "go-clean-api/cmd/domain/dto"
)

type (
	UsersJsonPlaceholderCache interface {
		Set(key string, user []response_jsonplaceholder.User, time int)
		Get(key string) ([]response_jsonplaceholder.User, error)
	}
)
