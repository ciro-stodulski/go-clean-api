package getuserusecase

import (
	entity "go-api/cmd/core/entities"
	"go-api/cmd/core/entities/user"
	"log"
	"strconv"
	"time"

	"github.com/google/uuid"
)

func (guuc *getUserUseCase) GetUser(id string) (*user.User, error) {
	iu := entity.ConvertId(id)

	u, err := guuc.RepositoryUser.GetById(iu)

	if err != nil {
		return nil, err
	}

	if u.ID == uuid.Nil {
		ujs, err := guuc.IntegrationJsonPlaceHolder.GetUsers()

		if err != nil {
			return nil, err
		}

		for _, uj := range ujs {
			id_string := strconv.Itoa(uj.Id)
			if id_string == id {
				log.Default().Print("Found user in integration:" + id)

				return &user.User{
					ID:        entity.NewID(),
					Name:      uj.Username,
					Email:     uj.Email,
					Password:  "test_for_integration",
					CreatedAt: time.Now(),
				}, nil
			}
		}

		log.Default().Print("not found user with id:" + id)
		return nil, user.ErrUserNotFound
	}

	return u, err
}
