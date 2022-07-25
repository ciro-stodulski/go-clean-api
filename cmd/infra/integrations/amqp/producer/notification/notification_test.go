package notificationproducer

import (
	"encoding/json"
	types_client "go-api/cmd/infra/integrations/amqp/client/types"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockAmqpClient struct {
	mock.Mock
}

func (mock *MockAmqpClient) Publish(body []byte, config types_client.ConfigAmqpClient) error {
	arg := mock.Called(body, config)

	return arg.Error(0)
}

func Test_User_Create_Producer(t *testing.T) {
	t.Run("succeffully", func(t *testing.T) {
		mac := new(MockAmqpClient)
		dto := Dto{
			Name:  "test",
			Event: "test",
		}

		config := types_client.ConfigAmqpClient{
			Exchange:    "notification.dx",
			Routing_key: "notify.create",
		}

		result, _ := json.Marshal(&dto)

		mac.On("Publish", []byte(string(result)), config).Return(nil)

		testService := New(mac)

		err := testService.SendNotify(dto)

		assert.Nil(t, err)
		mac.AssertCalled(t, "Publish", []byte(string(result)), config)
	})
}
