package jsonplaceholder

import (
	ports "go-api/src/core/ports"
	http_service "go-api/src/infra/integrations/http/client"
	"os"
)

type JsonPlaceholderIntegration struct {
	Http    http_service.HttpClient
	rootUrl string
}

func New(http http_service.HttpClient) ports.JsonPlaceholderIntegration {

	return &JsonPlaceholderIntegration{
		Http:    http,
		rootUrl: os.Getenv("JSON_PLACE_OLDER_INTEGRATION_URL"),
	}
}
