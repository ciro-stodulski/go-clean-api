package notificationproducer

import (
	"encoding/json"
	portsservice "go-clean-api/cmd/domain/services"
	mockamqpnotification "go-clean-api/cmd/shared/mocks/infra/integrations/amqp/notification"

	amqp "go-clean-api/cmd/infra/integrations/amqp"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_User_Create_Producer(t *testing.T) {
	t.Run("succeffully", func(t *testing.T) {
		// make mock
		mac := new(mockamqpnotification.MockAmqpClient)
		dto := portsservice.Dto{
			Name:  "test",
			Event: "test",
		}

		config := amqp.ConfigAmqpClient{
			Exchange:    "notification.dx",
			Routing_key: "notify.create",
		}

		result, _ := json.Marshal(&dto)

		mac.On("Publish", []byte(string(result)), config).Return(nil)
		//

		// test func
		testService := New(mac)
		err := testService.SendNotify(dto)
		//

		// asserts
		assert.Nil(t, err)
		mac.AssertCalled(t, "Publish", []byte(string(result)), config)
		//
	})
}
