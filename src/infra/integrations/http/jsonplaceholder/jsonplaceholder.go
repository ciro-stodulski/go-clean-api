package jsonplaceholder

import (
	response_jsonplaceholder "go-api/src/infra/integrations/http/jsonplaceholder/responses"
	"log"

	"encoding/json"
)

func (jpi *JsonPlaceholderIntegration) GetUsers() ([]response_jsonplaceholder.User, error) {
	response, err := jpi.Http.Get(jpi.rootUrl + "/users")

	if err != nil {
		log.Default().Printf("Error %s", err)

		return nil, err
	}

	var us []response_jsonplaceholder.User
	err = json.Unmarshal(response, &us)

	if err != nil {
		panic(err)
	}

	return us, nil
}
