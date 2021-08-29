package create_user_producer

import (
	interfaces "go-api/src/core/ports"
)

type createUserUseCase struct {
	ProducerUser interfaces.UserProducer
}

func NewUseCase(producer interfaces.UserProducer) CreateUserUseCase {
	return &createUserUseCase{
		ProducerUser: producer,
	}
}
