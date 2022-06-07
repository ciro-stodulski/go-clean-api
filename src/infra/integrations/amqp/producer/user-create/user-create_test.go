package user_create

import (
	"encoding/json"
	types_client "go-api/src/infra/integrations/amqp/client/types"
	create_dto "go-api/src/presentation/http/controllers/v1/users/create/dto"
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

func Test_CreateUser(t *testing.T) {
	t.Run("succeffully", func(t *testing.T) {
		MockAmqpClient := new(MockAmqpClient)
		dto := create_dto.CreateDto{
			Name:     "test",
			Email:    "test",
			Password: "test",
		}

		config := types_client.ConfigAmqpClient{
			Exchange:    "user.dx",
			Routing_key: "user.create",
		}

		result, _ := json.Marshal(&dto)

		MockAmqpClient.On("Publish", []byte(string(result)), config).Return(nil)

		testService := NewProdocer(MockAmqpClient)

		err := testService.CreateUser(dto)

		assert.Nil(t, err)
		MockAmqpClient.AssertCalled(t, "Publish", []byte(string(result)), config)
	})
}
