package userservice

import (
	"go-api/cmd/core/ports"
	portsservice "go-api/cmd/core/ports-service"
	usersjsonplaceholdercache "go-api/cmd/infra/repositories/cache/users-jsonplaceholder"
	userepository "go-api/cmd/infra/repositories/sql/user"
)

type userService struct {
	RepositoryUser             userepository.UserRepository
	IntegrationJsonPlaceHolder ports.JsonPlaceholderIntegration
	UsersJsonPlaceholderCache  usersjsonplaceholdercache.UsersJsonPlaceholderCache
}

func New(ur userepository.UserRepository, ji ports.JsonPlaceholderIntegration, ujc usersjsonplaceholdercache.UsersJsonPlaceholderCache) portsservice.UserService {
	return &userService{
		RepositoryUser:             ur,
		IntegrationJsonPlaceHolder: ji,
		UsersJsonPlaceholderCache:  ujc,
	}
}
