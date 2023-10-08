package userservice

import (
	domainjsonplaceholder "go-clean-api/cmd/domain/integration/http"
	domainusersjsonplaceholdercache "go-clean-api/cmd/domain/repository/cache"
	domainusersql "go-clean-api/cmd/domain/repository/sql"
	service "go-clean-api/cmd/domain/service"
)

type userService struct {
	SqlUser                    domainusersql.UserSql
	IntegrationJsonPlaceHolder domainjsonplaceholder.JsonPlaceholderIntegration
	UsersJsonPlaceholderCache  domainusersjsonplaceholdercache.UsersJsonPlaceholderCache
}

func New(
	ur domainusersql.UserSql,
	ji domainjsonplaceholder.JsonPlaceholderIntegration,
	ujc domainusersjsonplaceholdercache.UsersJsonPlaceholderCache,
) service.UserService {
	return &userService{
		SqlUser:                    ur,
		IntegrationJsonPlaceHolder: ji,
		UsersJsonPlaceholderCache:  ujc,
	}
}
