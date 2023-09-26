package httpclient

import (
	"net/http"
)

type (
	HttpClient interface {
		Request(req *http.Request) (*HttpResponse, error)
	}

	HttpResponse struct {
		Body       []byte
		StatusCode int
		Header     http.Header
	}
)
