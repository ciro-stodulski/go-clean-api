package registeruserusecase

import (
	domaindto "go-clean-api/cmd/domain/dto"
	"go-clean-api/cmd/domain/entities/user"
	portsservice "go-clean-api/cmd/domain/services"
	domainusecases "go-clean-api/cmd/domain/use-cases"
	"log"
)

type (
	registerUserUseCase struct {
		UserService         portsservice.UserService
		NotificationService portsservice.NotificationService
	}
)

func New(us portsservice.UserService, ns portsservice.NotificationService) domainusecases.RegisterUserUseCase {
	return &registerUserUseCase{
		UserService:         us,
		NotificationService: ns,
	}
}

func (cuuc *registerUserUseCase) Register(dto domaindto.Dto) (*user.User, error) {
	u, err := user.New(dto.Email, dto.Password, dto.Name)

	if err != nil {
		return nil, err
	}

	new_u, err := cuuc.UserService.Register(u)

	if err != nil {
		return nil, err
	}

	notification := domaindto.Event{
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
