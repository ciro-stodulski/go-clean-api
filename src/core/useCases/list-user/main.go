package list_users

import (
	interfaces "go-api/src/core/ports"
)

type listUsersUseCase struct {
	IntegrationJsonPlaceHolder interfaces.JsonPlaceholderIntegration
}

func NewUseCase(jsonPlaceholderIntegration interfaces.JsonPlaceholderIntegration) ListUsersUseCase {
	return &listUsersUseCase{
		IntegrationJsonPlaceHolder: jsonPlaceholderIntegration,
	}
}
