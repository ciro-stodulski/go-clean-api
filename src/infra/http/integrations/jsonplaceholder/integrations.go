package jsonplaceholder

import (
	response_jsonplaceholder "go-api/src/infra/http/integrations/jsonplaceholder/responses"
	"log"

	"encoding/json"
)

func (intergration *JsonPlaceholderIntegration) GetUsers() ([]response_jsonplaceholder.User, error) {
	response, err := intergration.Http.Get(intergration.rootUrl + "/users")

	if err != nil {
		log.Default().Printf("Error %s", err)

		return nil, err
	}

	var users []response_jsonplaceholder.User
	err = json.Unmarshal(response, &users)

	if err != nil {
		panic(err)
	}

	return users, nil
}
