package registeruserusecase

import (
	"go-clean-api/cmd/core/entities/user"
	portsservice "go-clean-api/cmd/core/ports"
	"log"
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

func (cuuc *registerUserUseCase) Register(dto Dto) (*user.User, error) {
	u, err := user.New(dto.Email, dto.Password, dto.Name)

	if err != nil {
		return nil, err
	}

	new_u, err := cuuc.UserService.Register(u)

	if err != nil {
		return nil, err
	}

	notification := portsservice.Dto{
		Name:  "REGISTERED_USER",
		Event: "USER",
	}

	err = cuuc.NotificationService.SendNotify(notification)

	id := cuuc.NotificationService.SaveNotify(notification)

	notification_mongo := cuuc.NotificationService.FindById(id)
	log.Default().Println("notification save in mongo", notification_mongo)

	if err != nil {
		return nil, err
	}

	return new_u, nil
}
