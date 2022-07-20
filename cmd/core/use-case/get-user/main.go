package getuserusecase

import (
	"go-api/cmd/core/ports"
)

type getUserUseCase struct {
	RepositoryUser             ports.UserRepository
	IntegrationJsonPlaceHolder ports.JsonPlaceholderIntegration
}

func New(
	ur ports.UserRepository,
	ji ports.JsonPlaceholderIntegration,
) GetUserUseCase {
	return &getUserUseCase{
		RepositoryUser:             ur,
		IntegrationJsonPlaceHolder: ji,
	}
}
