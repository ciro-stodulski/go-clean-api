package notificationproducer

import (
	"encoding/json"
	domaindto "go-clean-api/cmd/domain/dto"
	domainnotificationproducer "go-clean-api/cmd/domain/integration/amqp"
	amqpclient "go-clean-api/cmd/infra/integration/amqp"
)

type (
	userCreateProducer struct {
		clientAmqp  amqpclient.AmqpClient
		exchange    string
		routing_key string
	}
)

func New(amqpc amqpclient.AmqpClient) domainnotificationproducer.NotificationProducer {
	return &userCreateProducer{
		clientAmqp:  amqpc,
		exchange:    "notification.dx",
		routing_key: "notify.create",
	}
}

func (ucp *userCreateProducer) SendNotify(dto domaindto.Event) error {
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
