package deleteuserusecase

import (
	entity "go-api/src/core/entities"
	user "go-api/src/core/entities/user"
	"log"

	"github.com/google/uuid"
)

func (duuc *deleteUserUseCase) DeleteUser(id string) error {
	id_uuid := entity.ConvertId(id)

	u, err := duuc.RepositoryUser.GetById(id_uuid)

	if u.ID == uuid.Nil {
		log.Default().Print("Not found user with id:" + id)
		return user.ErrUserNotFound
	}

	if err != nil {
		return err
	}

	duuc.RepositoryUser.DeleteById(id_uuid)

	return nil
}
