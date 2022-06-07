package createuserproducerusecase

import (
	"go-api/src/core/ports"
)

type createUserUseCase struct {
	ProducerUser ports.UserProducer
}

func New(up ports.UserProducer) CreateUserUseCase {
	return &createUserUseCase{
		ProducerUser: up,
	}
}
