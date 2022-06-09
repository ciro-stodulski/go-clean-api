package http_service

import (
	"io/ioutil"
	"net/http"
)

type httpClient struct {
	client *http.Client
}

func New() HttpClient {
	return &httpClient{
		client: &http.Client{},
	}
}

func (http *httpClient) Get(url string) ([]byte, error) {
	res, err := http.client.Get(url)

	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)

	return body, nil
}
