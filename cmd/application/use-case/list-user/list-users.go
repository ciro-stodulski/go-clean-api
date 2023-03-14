package listusersusecase

import (
	response_jsonplaceholder "go-clean-api/cmd/domain/dto"
	portsservice "go-clean-api/cmd/domain/services"
	domainusecases "go-clean-api/cmd/domain/use-cases"
	"log"
)

type (
	listUsersUseCase struct {
		UserService portsservice.UserService
	}
)

func New(us portsservice.UserService) domainusecases.ListUsersUseCase {
	return &listUsersUseCase{
		UserService: us,
	}
}

func printUsers(ujs []response_jsonplaceholder.User) {
	for _, u := range ujs {
		log.Default().Print("-user:" + u.Username + "-email:" + u.Email + "-")
	}
}

func (luuc *listUsersUseCase) ListUsers() {
	ujs, errApp, err := luuc.UserService.ListUsers()

	if errApp != nil || err != nil {
		log.Default().Panic(errApp, err)
	}

	printUsers(ujs)
}
