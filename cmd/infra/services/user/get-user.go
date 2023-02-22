package userservice

import (
	entity "go-clean-api/cmd/domain/entities"
	"go-clean-api/cmd/domain/entities/user"
	domainexceptions "go-clean-api/cmd/domain/exceptions"
	"log"
	"strconv"
	"time"

	"github.com/google/uuid"
)

func (us *userService) GetUser(id string) (*user.User, error) {
	iu := entity.ConvertId(id)

	u, err := us.SqlUser.GetById(iu)

	if u.ID == uuid.Nil {
		ujs, err := us.IntegrationJsonPlaceHolder.GetUsers()

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
		return nil, domainexceptions.ErrUserNotFound
	}
	return u, err
}
