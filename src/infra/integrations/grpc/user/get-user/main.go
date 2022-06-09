package getuserservice

import (
	ports "go-api/src/core/ports"
)

type getUserService struct {
	service GetUserService
}

func New(service GetUserService) ports.GetUserService {

	return &getUserService{
		service: service,
	}
}
