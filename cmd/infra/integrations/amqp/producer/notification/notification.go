package notificationproducer

import (
	"encoding/json"
	amqp "go-api/cmd/infra/integrations/amqp/client"
	typesclient "go-api/cmd/infra/integrations/amqp/client/types"
)

type userCreateProducer struct {
	clientAmqp  amqp.AmqpClient
	exchange    string
	routing_key string
}

func New(amqpc amqp.AmqpClient) NotificationProducer {
	return &userCreateProducer{
		clientAmqp:  amqpc,
		exchange:    "notification.dx",
		routing_key: "notify.create",
	}
}

func (ucp *userCreateProducer) SendNotify(dto Dto) error {
	config := typesclient.ConfigAmqpClient{
		Exchange:    ucp.exchange,
		Routing_key: ucp.routing_key,
	}

	btresult, _ := json.Marshal(&dto)

	err := ucp.clientAmqp.Publish(
		[]byte(string(btresult)),
		config,
	)

	return err
}
