package mockhttpclient

import (
	httpclient "go-clean-api/cmd/infra/integration/http"
	"net/http"

	"github.com/stretchr/testify/mock"
)

type MockHttpClient struct {
	mock.Mock
}

func (mock *MockHttpClient) Request(req *http.Request) (*httpclient.HttpResponse, error) {
	arg := mock.Called(req)
	result := arg.Get(0)
	return result.(*httpclient.HttpResponse), arg.Error(1)
}
