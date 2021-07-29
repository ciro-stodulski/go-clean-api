package user_use_case

import (
	entity_root "go-api/src/core/entities"
	entity "go-api/src/core/entities/user"
	"log"

	"github.com/google/uuid"
)

//GetUser Get an user
func (service *Service) GetUser(id string) (*entity.User, error) {
	id_uuid := entity_root.ConvertId(id)

	user, err := service.RepositoryUser.GetById(id_uuid)

	if err != nil {
		return nil, err
	}

	if user.ID == uuid.Nil {
		log.Default().Print("not found user with id:" + id_uuid.String())
		return nil, entity.ErrUserNotFound
	}

	return user, err
}
