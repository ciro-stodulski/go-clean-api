package registeruserusecase

import (
	"go-api/cmd/core/entities/user"
	portsservice "go-api/cmd/core/ports"
)

type (
	Dto struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	RegisterUserUseCase interface {
		Register(dto Dto) (*user.User, error)
	}
	registerUserUseCase struct {
		UserService         portsservice.UserService
		NotificationService portsservice.NotificationService
	}
)

func New(us portsservice.UserService, ns portsservice.NotificationService) RegisterUserUseCase {
	return &registerUserUseCase{
		UserService:         us,
		NotificationService: ns,
	}
}
