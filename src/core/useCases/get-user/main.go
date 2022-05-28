package get_user

import (
	interfaces "go-api/src/core/ports"
)

type getUserUseCase struct {
	RepositoryUser             interfaces.UserRepository
	IntegrationJsonPlaceHolder interfaces.JsonPlaceholderIntegration
	GetUserService             interfaces.GetUserService
}

func NewUseCase(
	repository interfaces.UserRepository,
	jsonPlaceholderIntegration interfaces.JsonPlaceholderIntegration,
	getUserService interfaces.GetUserService,
) GetUserUseCase {
	return &getUserUseCase{
		RepositoryUser:             repository,
		IntegrationJsonPlaceHolder: jsonPlaceholderIntegration,
		GetUserService:             getUserService,
	}
}
