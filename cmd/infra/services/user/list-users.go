package userservice

import (
	response_jsonplaceholder "go-clean-api/cmd/infra/integrations/http/jsonplaceholder/responses"
	"log"
)

func (us *userService) ListUsers() []response_jsonplaceholder.User {
	ujs, err := us.IntegrationJsonPlaceHolder.GetUsers()

	if err != nil {
		log.Default().Print("###Error:Job failed, fail integration ###")
		panic(err)

	}

	if ujs == nil {
		log.Fatalln("###Error: error for get user in cache and integration###")
		panic("error for get user in cache and integration###")
	}

	log.Default().Print("***Set users in cache***")

	return ujs
}
