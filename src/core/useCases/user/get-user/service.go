package get_user

import (
	interfaces "go-api/src/core/ports"
)

//Service  interface
type Service struct {
	RepositoryUser             interfaces.Repository
	IntegrationJsonPlaceHolder interfaces.JsonPlaceholderIntegration
}

//NewService create new use case
func NewService(repository interfaces.Repository, jsonPlaceholderIntegration interfaces.JsonPlaceholderIntegration) UseCase {
	return &Service{
		RepositoryUser:             repository,
		IntegrationJsonPlaceHolder: jsonPlaceholderIntegration,
	}
}
