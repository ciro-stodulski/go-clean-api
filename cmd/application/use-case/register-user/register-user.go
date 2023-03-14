package registeruserusecase

import (
	domaindto "go-clean-api/cmd/domain/dto"
	"go-clean-api/cmd/domain/entities/user"
	domainexceptions "go-clean-api/cmd/domain/exceptions"
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

func (cuuc *registerUserUseCase) Register(dto domaindto.Dto) (*user.User, *domainexceptions.ApplicationException, error) {
	u, err := user.New(dto.Email, dto.Password, dto.Name)

	if err != nil {
		return nil, nil, err
	}

	new_u, errApp, err := cuuc.UserService.Register(u)

	if err != nil || errApp != nil {
		return nil, errApp, err
	}

	notification := domaindto.Event{
		Name:  "REGISTERED_USER",
		Event: "USER",
	}

	errApp, err = cuuc.NotificationService.SendNotify(notification)

	if err != nil || errApp != nil {
		return nil, errApp, err
	}

	id := cuuc.NotificationService.SaveNotify(notification)

	notification_mongo, errApp, err := cuuc.NotificationService.FindById(id)
	log.Default().Println("notification save in mongo", notification_mongo)

	return new_u, errApp, err
}
