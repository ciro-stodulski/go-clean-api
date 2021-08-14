package get_user

import (
	interfaces "go-api/src/core/ports"
)

type getUserUseCase struct {
	RepositoryUser             interfaces.Repository
	IntegrationJsonPlaceHolder interfaces.JsonPlaceholderIntegration
}

func NewUseCase(repository interfaces.Repository, jsonPlaceholderIntegration interfaces.JsonPlaceholderIntegration) GetUserUseCase {
	return &getUserUseCase{
		RepositoryUser:             repository,
		IntegrationJsonPlaceHolder: jsonPlaceholderIntegration,
	}
}
