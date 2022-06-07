package deleteuserusecase

import (
	entity "go-api/src/core/entities"
	user "go-api/src/core/entities/user"
	"log"

	"github.com/google/uuid"
)

func (usecase *deleteUserUseCase) DeleteUser(id string) error {
	id_uuid := entity.ConvertId(id)

	u, err := usecase.RepositoryUser.GetById(id_uuid)

	if u.ID == uuid.Nil {
		log.Default().Print("Not found user with id:" + id)
		return user.ErrUserNotFound
	}

	if err != nil {
		return err
	}

	usecase.RepositoryUser.DeleteById(id_uuid)

	return nil
}
