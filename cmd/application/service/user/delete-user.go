package userservice

import (
	entity "go-clean-api/cmd/domain/entity"
	exception "go-clean-api/cmd/domain/exception"
	"log"

	"github.com/google/uuid"
)

func (duuc *userService) DeleteUser(id string) error {
	id_uuid := entity.ConvertId(id)

	u, err := duuc.SqlUser.GetById(id_uuid)

	if u.ID == uuid.Nil {
		log.Default().Print("Not found user with id:" + id)
		return exception.UserNotFound()
	}

	if err != nil {
		return err
	}

	duuc.SqlUser.DeleteById(id_uuid)

	return nil
}
