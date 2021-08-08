package get_user

import (
	interfaces "go-api/src/core/ports"
)

type Service struct {
	RepositoryUser             interfaces.Repository
	IntegrationJsonPlaceHolder interfaces.JsonPlaceholderIntegration
}

func NewService(repository interfaces.Repository, jsonPlaceholderIntegration interfaces.JsonPlaceholderIntegration) UseCase {
	return &Service{
		RepositoryUser:             repository,
		IntegrationJsonPlaceHolder: jsonPlaceholderIntegration,
	}
}
