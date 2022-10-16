package mockservicesnotification

import (
	portsservice "go-clean-api/cmd/core/ports"

	"github.com/stretchr/testify/mock"
)

type MockNotificationServices struct {
	mock.Mock
}

func (mock *MockNotificationServices) SendNotify(dto portsservice.Dto) error {
	arg := mock.Called(dto)
	return arg.Error(0)
}

func (mock *MockNotificationServices) CheckNotify(msg string) (string error) {
	arg := mock.Called(msg)
	return arg.Error(1)
}

func (mock *MockNotificationServices) SaveNotify(dto portsservice.Dto) string {
	arg := mock.Called(dto)
	result := arg.Get(0)

	return result.(string)
}

func (mock *MockNotificationServices) FindById(msg string) *portsservice.Dto {
	arg := mock.Called(msg)
	result := arg.Get(0)

	return result.(*portsservice.Dto)
}
