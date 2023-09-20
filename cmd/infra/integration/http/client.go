package httpclient

import (
	"net/http"
)

type (
	HttpClient interface {
		Do(req *http.Request) (*HttpResponse, error)
	}

	HttpResponse struct {
		Body       []byte
		StatusCode int
		Header     http.Header
	}
)
