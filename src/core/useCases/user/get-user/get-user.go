package get_user

import (
	entity_root "go-api/src/core/entities"
	entity "go-api/src/core/entities/user"
	"log"
	"strconv"
	"time"

	"github.com/google/uuid"
)

func (service *Service) GetUser(id string) (*entity.User, error) {
	id_uuid := entity_root.ConvertId(id)

	user, err := service.RepositoryUser.GetById(id_uuid)

	if err != nil {
		return nil, err
	}

	userJson, err := service.IntegrationJsonPlaceHolder.GetUsers()

	if err != nil {
		return nil, err
	}

	if user.ID == uuid.Nil {
		for _, user := range userJson {
			id_string := strconv.Itoa(user.Id)
			if id_string == id {
				log.Default().Print("found user in integration:" + id)

				return &entity.User{
					ID:        entity_root.NewID(),
					Name:      user.Username,
					Email:     user.Email,
					Password:  "test_for_integration",
					CreatedAt: time.Now(),
				}, nil
			}
		}

		log.Default().Print("not found user with id:" + id)
		return nil, entity.ErrUserNotFound
	}

	return user, err
}
