package entity

import (
	"github.com/google/uuid"
)

//ID entity ID
type ID = uuid.UUID

//NewID create a new entity ID
func NewID() ID {
	return ID(uuid.New())
}

func ConvertId(id string) uuid.UUID {
	string, _ := uuid.Parse(id)

	return string
}
