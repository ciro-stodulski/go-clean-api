package mockamqpnotification

import (
	amqpclient "go-clean-api/cmd/infra/integrations/amqp"

	"github.com/stretchr/testify/mock"
)

type MockAmqpClient struct {
	mock.Mock
}

func (mock *MockAmqpClient) Publish(body []byte, config amqpclient.ConfigAmqpClient) error {
	arg := mock.Called(body, config)

	return arg.Error(0)
}
