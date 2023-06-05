package userservice

import (
	domainjsonplaceholder "go-clean-api/cmd/domain/integrations/http"
	domainusersjsonplaceholdercache "go-clean-api/cmd/domain/repositories/cache"
	domainusersql "go-clean-api/cmd/domain/repositories/sql"
	portsservice "go-clean-api/cmd/domain/services"
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
) portsservice.UserService {
	return &userService{
		SqlUser:                    ur,
		IntegrationJsonPlaceHolder: ji,
		UsersJsonPlaceholderCache:  ujc,
	}
}
