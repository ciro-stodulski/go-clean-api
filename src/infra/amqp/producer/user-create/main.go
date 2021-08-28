package user_create

import (
	"go-api/src/infra/amqp/producer"
	amqp "go-api/src/main/module/amqp/rabbitmq/client"
)

type userCreate struct {
	clientAmqp  amqp.IAmqpClient
	exchange    string
	routing_key string
}

func NewProdocer(clientAmqp amqp.IAmqpClient) producer.Producer {
	return &userCreate{
		clientAmqp: clientAmqp,
	}
}
