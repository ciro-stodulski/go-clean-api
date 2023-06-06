package mockclientcache

import (
	"github.com/stretchr/testify/mock"
)

type MockCache struct {
	mock.Mock
}

func (mock *MockCache) Get(key string) (string, error) {
	arg := mock.Called(key)
	result := arg.Get(0)
	return result.(string), arg.Error(1)
}

func (mock *MockCache) Set(key string, value string, timeEx int) error {
	arg := mock.Called(key, value, timeEx)
	return arg.Error(0)
}
