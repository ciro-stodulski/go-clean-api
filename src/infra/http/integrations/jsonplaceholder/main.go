package jsonplaceholder

import (
	http_service "go-api/src/main/module/http/client"
)

type JsonPlaceholderIntegration struct {
	http *http_service.HttpService
}

func New(http *http_service.HttpService) *JsonPlaceholderIntegration {
	return &JsonPlaceholderIntegration{
		http: http,
	}
}
