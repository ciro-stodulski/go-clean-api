package entity

import (
	"github.com/google/uuid"
)

type ID = uuid.UUID

func NewID() ID {
	return ID(uuid.New())
}

func ConvertId(id string) uuid.UUID {
	string, _ := uuid.Parse(id)

	return string
}
