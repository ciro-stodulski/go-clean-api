package ports

import create_dto "go-api/src/presentation/http/controllers/v1/users/create/dto"

type (
	UserProducer interface {
		CreateUser(dto create_dto.CreateDto) error
	}
)
