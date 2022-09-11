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
