package jsonplaceholder

import (
	ports "go-api/cmd/core/ports"
	http_service "go-api/cmd/infra/integrations/http/client"
	"go-api/cmd/shared/env"
)

type JsonPlaceholderIntegration struct {
	Http    http_service.HttpClient
	rootUrl string
}

func New(http http_service.HttpClient) ports.JsonPlaceholderIntegration {

	return &JsonPlaceholderIntegration{
		Http:    http,
		rootUrl: env.Env().JsonPlaceOlderIntegrationUrl,
	}
}
