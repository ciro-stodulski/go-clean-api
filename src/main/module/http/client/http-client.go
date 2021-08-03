package http_service

import (
	"io/ioutil"
	"net/http"
)

type HttpService struct {
	Engine *http.Client
}

func New() *HttpService {
	return &HttpService{
		Engine: &http.Client{},
	}
}

func (http *HttpService) Get(url string) (interface{}, error) {
	result, err := http.Engine.Get(url)
	if err != nil {
		return nil, err
	}
	defer result.Body.Close()

	body, err := ioutil.ReadAll(result.Body)

	return body, nil
}
