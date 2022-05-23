package delete_user

import (
	entity_root "go-api/src/core/entities"
	user_entity "go-api/src/core/entities/user"
	"log"

	"github.com/google/uuid"
)

func (service *deleteUserUseCase) DeleteUser(id string) error {
	id_uuid := entity_root.ConvertId(id)

	user, err := service.RepositoryUser.GetById(id_uuid)

	if user.ID == uuid.Nil {
		log.Default().Print("not found user with id:" + id)
		return user_entity.ErrUserNotFound
	}

	if err != nil {
		return err
	}

	service.RepositoryUser.DeleteById(id_uuid)

	return nil
}
