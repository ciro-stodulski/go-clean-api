package listusersusecase

import (
	response_jsonplaceholder "go-api/src/infra/integrations/http/jsonplaceholder/responses"
	"log"
	"reflect"
)

func (usecase *listUsersUseCase) ListUsers() {
	ujs, err := usecase.UsersCache.Get("users")

	if err != nil {
		log.Default().Print("###Error:Job failed, fail cache ###")
		return
	}

	if reflect.DeepEqual(ujs, []response_jsonplaceholder.User{}) {
		ujs, err := usecase.IntegrationJsonPlaceHolder.GetUsers()

		if err != nil {
			log.Default().Print("###Error:Job failed, fail integration ###")
			return
		}

		if ujs == nil {
			log.Fatalln("###Error: error for get user in cache and integration###")
			return
		}

		usecase.UsersCache.Set("users", ujs, 100)
		log.Default().Print("***Set users in cache***")

		printUsers(ujs)
	} else {
		log.Default().Print("***Get users by cache***")

		printUsers(ujs)
	}
}

func printUsers(ujs []response_jsonplaceholder.User) {
	for _, u := range ujs {
		log.Default().Print("-user:" + u.Username + "-email:" + u.Email + "-")
	}
}
