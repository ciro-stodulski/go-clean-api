package user_use_case

import (
	"fmt"
	entity_root "go-api/src/core/entities"
	entity "go-api/src/core/entities/user"

	"github.com/google/uuid"
)

//GetUser Get an user
func (service *Service) GetUser(id entity_root.ID) (*entity.User, error) {
	user, err := service.RepositoryUser.GetById(id)

	if err != nil {
		return nil, err
	}

	fmt.Println(user)

	if user.ID == uuid.Nil {
		return nil, entity.ErrUserNotFound
	}

	return user, err
}
