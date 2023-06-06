package listusersusecase

import (
	response_jsonplaceholder "go-clean-api/cmd/domain/dto"
	portsservice "go-clean-api/cmd/domain/services"
	domainusecases "go-clean-api/cmd/domain/use-case"
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
	ujs, err := luuc.UserService.ListUsers()

	if err != nil {
		log.Default().Panic(err)
	}

	printUsers(ujs)
}
