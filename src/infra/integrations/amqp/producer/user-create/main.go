package user_create

import (
	"go-api/src/core/ports"
	amqp "go-api/src/infra/integrations/amqp/client"
)

type userCreate struct {
	clientAmqp  amqp.IAmqpClient
	exchange    string
	routing_key string
}

func NewProdocer(clientAmqp amqp.IAmqpClient) ports.UserProducer {
	return &userCreate{
		clientAmqp:  clientAmqp,
		exchange:    "user.dx",
		routing_key: "user.create",
	}
}
