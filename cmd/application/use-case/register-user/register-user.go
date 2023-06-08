package registeruserusecase

import (
	usecase "go-clean-api/cmd/domain/use-case"

	"go-clean-api/cmd/domain/dto"
	"go-clean-api/cmd/domain/entity/user"
	"go-clean-api/cmd/domain/service"
	"log"
)

type (
	registerUserUseCase struct {
		UserService         service.UserService
		NotificationService service.NotificationService
	}
)

func New(us service.UserService, ns service.NotificationService) usecase.RegisterUserUseCase {
	return &registerUserUseCase{
		UserService:         us,
		NotificationService: ns,
	}
}

func (cuuc *registerUserUseCase) Register(data dto.RegisterUser) (*user.User, error) {
	u, err := user.New(data.Email, data.Password, data.Name)

	if err != nil {
		return nil, err
	}

	new_u, err := cuuc.UserService.Register(u)

	if err != nil {
		return nil, err
	}

	notification := dto.Event{
		Name:  "REGISTERED_USER",
		Event: "USER",
	}

	err = cuuc.NotificationService.SendNotify(notification)

	if err != nil {
		return nil, err
	}

	id := cuuc.NotificationService.SaveNotify(notification)

	notification_mongo, err := cuuc.NotificationService.FindById(id)
	log.Default().Println("notification save in mongo", notification_mongo)

	return new_u, err
}
