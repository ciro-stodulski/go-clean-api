package usersjsonplaceholdercache

import integration "go-api/cmd/infra/integrations/http/jsonplaceholder/responses"

type (
	UsersJsonPlaceholderCache interface {
		Set(key string, user []integration.User, time int)
		Get(key string) ([]integration.User, error)
	}
)
