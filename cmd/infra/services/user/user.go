package userservice

import (
	portsservice "go-api/cmd/core/ports"
	"go-api/cmd/infra/integrations/http/jsonplaceholder"
	usersjsonplaceholdercache "go-api/cmd/infra/repositories/cache/users-jsonplaceholder"
	userepository "go-api/cmd/infra/repositories/sql/user"
)

type userService struct {
	SqlUser                    userepository.UserSql
	IntegrationJsonPlaceHolder jsonplaceholder.JsonPlaceholderIntegration
	UsersJsonPlaceholderCache  usersjsonplaceholdercache.UsersJsonPlaceholderCache
}

func New(
	ur userepository.UserSql,
	ji jsonplaceholder.JsonPlaceholderIntegration,
	ujc usersjsonplaceholdercache.UsersJsonPlaceholderCache,
) portsservice.UserService {
	return &userService{
		SqlUser:                    ur,
		IntegrationJsonPlaceHolder: ji,
		UsersJsonPlaceholderCache:  ujc,
	}
}
