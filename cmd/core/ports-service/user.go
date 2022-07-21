package portsservice

import (
	response_jsonplaceholder "go-api/cmd/infra/integrations/http/jsonplaceholder/responses"

	"go-api/cmd/core/entities/user"
	create_dto "go-api/cmd/interface/amqp/consumers/users/create/dto"
)

type (
	UserService interface {
		GetUser(id string) (*user.User, error)
		CreateUser(dto create_dto.CreateDto) (*user.User, error)
		DeleteUser(id string) error
		ListUsers() []response_jsonplaceholder.User
	}
)
