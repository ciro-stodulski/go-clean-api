package list_users

import (
	response_jsonplaceholder "go-api/src/infra/http/integrations/jsonplaceholder/responses"
	"log"
	"reflect"
)

func (service *listUsersUseCase) ListUsers() {
	userJson := service.UsersCache.Get("users")

	if reflect.DeepEqual(userJson, []response_jsonplaceholder.User{}) {
		userJson, err := service.IntegrationJsonPlaceHolder.GetUsers()

		service.UsersCache.Set("users", userJson, 100)
		log.Default().Print("***Set users in cache***")

		if err != nil {
			log.Default().Print("***Job failed***")
		}
	} else {
		log.Default().Print("***Get users by cache***")
	}

	for _, user := range userJson {
		log.Default().Print("---Job for integration user:" + user.Username + "email:" + user.Email + "---")
	}
}
