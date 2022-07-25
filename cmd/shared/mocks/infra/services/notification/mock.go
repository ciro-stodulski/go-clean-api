package mockservicesnotification

import (
	notificationproducer "go-api/cmd/infra/integrations/amqp/notification"

	"github.com/stretchr/testify/mock"
)

type MockNotificationServices struct {
	mock.Mock
}

func (mock *MockNotificationServices) SendNotify(dto notificationproducer.Dto) error {
	arg := mock.Called()
	return arg.Error(0)
}

func (mock *MockNotificationServices) CheckNotify(msg string) (string error) {
	arg := mock.Called()
	return arg.Error(0)
}
