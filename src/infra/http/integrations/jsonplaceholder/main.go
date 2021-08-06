package jsonplaceholder

import (
	ports "go-api/src/core/ports"
	http_service "go-api/src/main/module/http/client"
)

type JsonPlaceholderIntegration struct {
	Http    http_service.HttpClient
	rootUrl string
}

func New(http http_service.HttpClient, root_url string) ports.JsonPlaceholderIntegration {

	return &JsonPlaceholderIntegration{
		Http:    http,
		rootUrl: root_url,
	}
}
