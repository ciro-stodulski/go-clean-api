package httpadapter

import (
	httpclient "go-clean-api/cmd/infra/integration/http"
	"io"
	"net/http"
)

type (
	httpClient struct {
		http.Client
	}
)

func New() httpclient.HttpClient {
	return &httpClient{}
}

func (https *httpClient) Request(req *http.Request) (*httpclient.HttpResponse, error) {
	res, err := https.Do(req)
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
