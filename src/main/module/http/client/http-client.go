package http_service

import (
	"io/ioutil"
	"net/http"
)

type httpClient struct {
	Engine *http.Client
}

func New() HttpClient {
	return &httpClient{
		Engine: &http.Client{},
	}
}

func (http *httpClient) Get(url string) ([]byte, error) {
	result, err := http.Engine.Get(url)

	if err != nil {
		return nil, err
	}
	defer result.Body.Close()

	body, _ := ioutil.ReadAll(result.Body)

	return body, nil
}
