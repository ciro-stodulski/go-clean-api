package createuserproducerusecase

import (
	port "go-api/src/core/ports"
	"log"
)

func (cupuc *createUserProducerUseCase) CreateUser(dto port.CreateDto) error {
	err := cupuc.ProducerUser.CreateUser(dto)
	log.Default().Print("Send message to create user in producer mode")

	if err != nil {
		return err
	}

	return nil
}
