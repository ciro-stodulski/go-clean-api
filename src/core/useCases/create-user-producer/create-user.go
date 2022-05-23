package create_user_producer

import (
	dto "go-api/src/presentation/http/controllers/v1/users/create/dto"
	"log"
)

func (service *createUserUseCase) CreateUser(dto dto.CreateDto) error {
	err := service.ProducerUser.CreateUser(dto)
	log.Default().Print("Send message to create user in producer mode")

	if err != nil {
		return err
	}

	return nil
}
