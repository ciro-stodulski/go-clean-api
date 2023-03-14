package userservice

import (
	entity "go-clean-api/cmd/domain/entities"
	domainexceptions "go-clean-api/cmd/domain/exceptions"
	"log"

	"github.com/google/uuid"
)

func (duuc *userService) DeleteUser(id string) (*domainexceptions.ApplicationException, error) {
	id_uuid := entity.ConvertId(id)

	u, err := duuc.SqlUser.GetById(id_uuid)

	if u.ID == uuid.Nil {
		log.Default().Print("Not found user with id:" + id)
		return domainexceptions.UserNotFound(), nil
	}

	if err != nil {
		return nil, err
	}

	duuc.SqlUser.DeleteById(id_uuid)

	return nil, err
}
