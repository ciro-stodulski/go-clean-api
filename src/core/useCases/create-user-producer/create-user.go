package create_user_producer

import (
	"fmt"
	dto "go-api/src/presentation/http/controllers/v1/users/create/dto"
)

func (service *createUserUseCase) CreateUser(dto dto.CreateDto) error {
	err := service.ProducerUser.CreateUser(dto)
	fmt.Println("Send message to create user in producer mode")

	if err != nil {
		return err
	}

	return nil
}
