package usecase

type (
	DeleteUserUseCase interface {
		DeleteUser(id string) error
	}
)
