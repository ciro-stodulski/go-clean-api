package get_user

import (
	interfaces "go-api/src/core/ports"
)

type getUserUseCase struct {
	RepositoryUser             interfaces.UserRepository
	IntegrationJsonPlaceHolder interfaces.JsonPlaceholderIntegration
}

func NewUseCase(
	repository interfaces.UserRepository,
	jsonPlaceholderIntegration interfaces.JsonPlaceholderIntegration,
) GetUserUseCase {
	return &getUserUseCase{
		RepositoryUser:             repository,
		IntegrationJsonPlaceHolder: jsonPlaceholderIntegration,
	}
}
