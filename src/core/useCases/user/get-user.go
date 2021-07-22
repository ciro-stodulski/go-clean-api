package user_use_case

import (
	entity_root "go-api/src/core/entities"
	entity "go-api/src/core/entities/user"
	"log"

	"github.com/google/uuid"
)

//GetUser Get an user
func (service *Service) GetUser(id entity_root.ID) (*entity.User, error) {
	user, err := service.RepositoryUser.GetById(id)

	if err != nil {
		return nil, err
	}

	if user.ID == uuid.Nil {
		log.Default().Print("not found user with id:" + id.String())
		return nil, entity.ErrUserNotFound
	}

	return user, err
}
