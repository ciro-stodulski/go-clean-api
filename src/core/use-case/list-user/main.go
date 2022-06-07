package listusersusecase

import (
	"go-api/src/core/ports"
)

type listUsersUseCase struct {
	IntegrationJsonPlaceHolder ports.JsonPlaceholderIntegration
	UsersCache                 ports.UsersCache
}

func New(ji ports.JsonPlaceholderIntegration, uc ports.UsersCache) ListUsersUseCase {
	return &listUsersUseCase{
		IntegrationJsonPlaceHolder: ji,
		UsersCache:                 uc,
	}
}
