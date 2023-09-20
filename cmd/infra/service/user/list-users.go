package userservice

import (
	response_jsonplaceholder "go-clean-api/cmd/domain/dto"
	"log"
	"reflect"
)

func (us *userService) ListUsers() ([]response_jsonplaceholder.User, error) {
	ujs, err := us.UsersJsonPlaceholderCache.Get("users")

	if err != nil {
		log.Default().Print("###Error:Job failed, fail cache ###")
		panic(err)
	}

	if reflect.DeepEqual(ujs, []response_jsonplaceholder.User{}) {
		ujs, err := us.IntegrationJsonPlaceHolder.GetUsers()

		if err != nil {
			log.Default().Print("###Error:Job failed, fail integration ###")
			panic(err)

		}

		if ujs == nil {
			log.Fatalln("###Error: error for get user in cache and integration###")
			panic("error for get user in cache and integration###")
		}

		us.UsersJsonPlaceholderCache.Set("users", ujs, 1)
		log.Default().Print("***Set users in cache***")

		return ujs, nil
	}
	log.Default().Print("***Get users by cache***")

	return ujs, nil
}
