package get_user

import (
	entity_root "go-api/src/core/entities"
	entity "go-api/src/core/entities/user"
	"log"
	"strconv"
	"time"

	"github.com/google/uuid"
)

func (service *getUserUseCase) GetUser(id string) (*entity.User, error) {
	// criar um exemplo melhor parar a integração do grpc
	user, err := service.GetUserService.GetUser(id)

	if err != nil {
		return nil, err
	}

	if user.ID == uuid.Nil {
		userJson, err := service.IntegrationJsonPlaceHolder.GetUsers()

		if err != nil {
			return nil, err
		}

		for _, user := range userJson {
			id_string := strconv.Itoa(user.Id)
			if id_string == id {
				log.Default().Print("Found user in integration:" + id)

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
