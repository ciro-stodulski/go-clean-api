package createuserproducerusecase

import (
	"go-api/cmd/core/ports"
)

type createUserProducerUseCase struct {
	ProducerUser ports.UserProducer
}

func New(up ports.UserProducer) CreateUserUseCase {
	return &createUserProducerUseCase{
		ProducerUser: up,
	}
}
