package listusersusecase

import (
	portsservice "go-api/cmd/core/ports-service"
	response_jsonplaceholder "go-api/cmd/infra/integrations/http/jsonplaceholder/responses"
	"log"
)

type listUsersUseCase struct {
	UserService portsservice.UserService
}

func New(us portsservice.UserService) ListUsersUseCase {
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
	ujs := luuc.UserService.ListUsers()

	printUsers(ujs)
}
