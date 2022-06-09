package createuserproducerusecase

import (
	"go-api/src/core/ports"
)

type createUserProducerUseCase struct {
	ProducerUser ports.UserProducer
}

func New(up ports.UserProducer) CreateUserUseCase {
	return &createUserProducerUseCase{
		ProducerUser: up,
	}
}
