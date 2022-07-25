package userservice

import (
	entity "go-api/cmd/core/entities"
	user "go-api/cmd/core/entities/user"
	"log"

	"github.com/google/uuid"
)

func (duuc *userService) DeleteUser(id string) error {
	id_uuid := entity.ConvertId(id)

	u, err := duuc.SqlUser.GetById(id_uuid)

	if u.ID == uuid.Nil {
		log.Default().Print("Not found user with id:" + id)
		return user.ErrUserNotFound
	}

	if err != nil {
		return err
	}

	duuc.SqlUser.DeleteById(id_uuid)

	return nil
}
