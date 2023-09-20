package jsonplaceholder

import (
	"fmt"
	response_jsonplaceholder "go-clean-api/cmd/domain/dto"
	domainjsonplaceholder "go-clean-api/cmd/domain/integration/http"
	http_service "go-clean-api/cmd/infra/integration/http"
	"go-clean-api/cmd/shared/env"
	"log"
	"net/http"
	"net/url"

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
	response, err := jpi.Http.Do(
		&http.Request{
			Method: http.MethodGet,
			URL: &url.URL{
				Path:   "/users",
				Host:   jpi.rootUrl,
				Scheme: "https",
			},
		},
	)

	if err != nil {
		log.Default().Printf("[jsonPlaceholderIntegration] Error %s", err)

		return nil, err
	}

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("[jsonPlaceholderIntegration] Error status: %v", response.StatusCode)
	}

	var us []response_jsonplaceholder.User
	err = json.Unmarshal(response.Body, &us)

	if err != nil {
		panic(err)
	}

	return us, nil
}
