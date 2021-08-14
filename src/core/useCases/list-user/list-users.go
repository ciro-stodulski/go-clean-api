package list_users

import (
	"log"
)

func (service *listUsersUseCase) ListUsers() {
	userJson, err := service.IntegrationJsonPlaceHolder.GetUsers()

	if err != nil {
		log.Default().Print("***Job failed***")
	}

	for _, user := range userJson {
		log.Default().Print("---Job for integration user:" + user.Username + "email:" + user.Email + "---")
	}
}
