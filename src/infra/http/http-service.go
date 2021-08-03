package http_service

import "github.com/monaco-io/request"

type HttpService struct {
	Engine *request.Client
}

func (http_server *HttpService) CreateEngine(url string) *HttpService {
	return &HttpService{
		Engine: &request.Client{
			URL: url,
		},
	}
}
