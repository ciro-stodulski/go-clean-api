package mockusercache

import (
	response_jsonplaceholder "go-clean-api/cmd/infra/integrations/http/jsonplaceholder/responses"

	"github.com/stretchr/testify/mock"
)

type MockCache struct {
	mock.Mock
}

func (mock *MockCache) Get(key string) ([]response_jsonplaceholder.User, error) {
	arg := mock.Called(key)
	result := arg.Get(0)
	return result.([]response_jsonplaceholder.User), arg.Error(1)
}

func (mock *MockCache) Set(key string, value []response_jsonplaceholder.User, timeEx int) {
	mock.Called(key, value, timeEx)
}
