package deleteuserusecase

type (
	DeleteUserUseCase interface {
		DeleteUser(id string) error
	}
)
