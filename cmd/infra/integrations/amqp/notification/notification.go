package notificationproducer

import (
	"encoding/json"
	portsservice "go-clean-api/cmd/core/ports"
	amqpclient "go-clean-api/cmd/infra/integrations/amqp"
)

type (
	NotificationProducer interface {
		SendNotify(dto portsservice.Dto) error
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

func (ucp *userCreateProducer) SendNotify(dto portsservice.Dto) error {
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
