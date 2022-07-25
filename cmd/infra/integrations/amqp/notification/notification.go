package notificationproducer

import (
	"encoding/json"
	amqpclient "go-api/cmd/infra/integrations/amqp"
)

type (
	Dto struct {
		Name  string `json:"name"`
		Event string `json:"event"`
	}

	NotificationProducer interface {
		SendNotify(dto Dto) error
	}

	userCreateProducer struct {
		clientAmqp  amqpclient.AmqpClient
		exchange    string
		routing_key string
	}
)

func New(amqpc amqpclient.AmqpClient) NotificationProducer {
	return &userCreateProducer{
		clientAmqp:  amqpc,
		exchange:    "notification.dx",
		routing_key: "notify.create",
	}
}

func (ucp *userCreateProducer) SendNotify(dto Dto) error {
	config := amqpclient.ConfigAmqpClient{
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
