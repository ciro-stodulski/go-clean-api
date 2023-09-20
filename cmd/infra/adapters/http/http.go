package httpadapter

import (
	httpclient "go-clean-api/cmd/infra/integration/http"
	"io"
	"net/http"
)

type (
	httpClient struct {
		client http.Client
	}
)

func New() httpclient.HttpClient {
	return &httpClient{
		client: http.Client{},
	}
}

func (https *httpClient) Do(req *http.Request) (*httpclient.HttpResponse, error) {
	res, err := https.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, _ := io.ReadAll(res.Body)

	return &httpclient.HttpResponse{
		Body:       body,
		StatusCode: res.StatusCode,
		Header:     res.Header,
	}, nil
}
