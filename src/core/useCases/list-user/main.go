package list_users

import (
	interfaces "go-api/src/core/ports"
)

type listUsersUseCase struct {
	IntegrationJsonPlaceHolder interfaces.JsonPlaceholderIntegration
	UsersCache                 interfaces.UsersCache
}

func NewUseCase(jsonPlaceholderIntegration interfaces.JsonPlaceholderIntegration, usersCache interfaces.UsersCache) ListUsersUseCase {
	return &listUsersUseCase{
		IntegrationJsonPlaceHolder: jsonPlaceholderIntegration,
		UsersCache:                 usersCache,
	}
}
