package create_user_producer

import (
	dto "go-api/src/presentation/http/controllers/v1/users/create/dto"
)

func (service *createUserUseCase) CreateUser(dto dto.CreateDto) error {
	err := service.ProducerUser.CreateUser(dto)

	if err != nil {
		return err
	}

	return nil
}
