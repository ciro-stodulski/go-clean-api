package list_users

import (
	response_jsonplaceholder "go-api/src/infra/integrations/http/jsonplaceholder/responses"
	"log"
	"reflect"
)

func (service *listUsersUseCase) ListUsers() {
	userJson, err := service.UsersCache.Get("users")

	if err != nil {
		log.Default().Print("###Error:Job failed, fail cache ###")
		return
	}

	if reflect.DeepEqual(userJson, []response_jsonplaceholder.User{}) {
		userJson, err := service.IntegrationJsonPlaceHolder.GetUsers()

		if err != nil {
			log.Default().Print("###Error:Job failed, fail integration ###")
			return
		}

		if userJson == nil {
			log.Fatalln("###Error: error for get user in cache and integration###")
			return
		}

		service.UsersCache.Set("users", userJson, 100)
		log.Default().Print("***Set users in cache***")

		printUsers(userJson)
	} else {
		log.Default().Print("***Get users by cache***")

		printUsers(userJson)
	}
}

func printUsers(users []response_jsonplaceholder.User) {
	for _, user := range users {
		log.Default().Print("-user:" + user.Username + "-email:" + user.Email + "-")
	}
}
