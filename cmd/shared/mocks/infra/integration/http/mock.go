package mockhttpclient

import "github.com/stretchr/testify/mock"

type MockHttpClient struct {
	mock.Mock
}

func (mock *MockHttpClient) Get(url string) ([]byte, error) {
	arg := mock.Called(url)
	result := arg.Get(0)
	return result.([]byte), arg.Error(1)
}
