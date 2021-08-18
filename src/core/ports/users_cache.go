package ports

import integration "go-api/src/infra/http/integrations/jsonplaceholder/responses"

type (
	UsersCache interface {
		Set(key string, value []integration.User, timeEx int)
		Get(key string) []integration.User
	}
)