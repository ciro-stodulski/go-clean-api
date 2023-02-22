package jsonplaceholder

import (
	response_jsonplaceholder "go-clean-api/cmd/domain/dto"
	domainjsonplaceholder "go-clean-api/cmd/domain/integrations/http"
	http_service "go-clean-api/cmd/infra/integrations/http"
	"go-clean-api/cmd/shared/env"
	"log"

	"encoding/json"
)

type (
	jsonPlaceholderIntegration struct {
		Http    http_service.HttpClient
		rootUrl string
	}
)

func New(http http_service.HttpClient) domainjsonplaceholder.JsonPlaceholderIntegration {

	return &jsonPlaceholderIntegration{
		Http:    http,
		rootUrl: env.Env().JsonPlaceOlderIntegrationUrl,
	}
}

func (jpi *jsonPlaceholderIntegration) GetUsers() ([]response_jsonplaceholder.User, error) {
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
