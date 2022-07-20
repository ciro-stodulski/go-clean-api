package usercreateproducer

import (
	"go-api/cmd/core/ports"
	amqp "go-api/cmd/infra/integrations/amqp/client"
)

type userCreateProducer struct {
	clientAmqp  amqp.AmqpClient
	exchange    string
	routing_key string
}

func New(amqpc amqp.AmqpClient) ports.UserProducer {
	return &userCreateProducer{
		clientAmqp:  amqpc,
		exchange:    "user.dx",
		routing_key: "user.create",
	}
}
