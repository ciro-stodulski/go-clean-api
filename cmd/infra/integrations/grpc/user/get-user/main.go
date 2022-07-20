package getuserservice

import (
	ports "go-api/cmd/core/ports"
)

type getUserService struct {
	service GetUserService
}

func New(service GetUserService) ports.GetUserService {

	return &getUserService{
		service: service,
	}
}
