package delete_user

type (
	DeleteUserUseCase interface {
		DeleteUser(id string) error
	}
)
