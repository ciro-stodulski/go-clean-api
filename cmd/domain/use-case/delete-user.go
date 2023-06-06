package domainusecases

type (
	DeleteUserUseCase interface {
		DeleteUser(id string) error
	}
)
