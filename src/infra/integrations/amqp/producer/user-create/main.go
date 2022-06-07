package usercreateproducer

import (
	"go-api/src/core/ports"
	amqp "go-api/src/infra/integrations/amqp/client"
)

type userCreateProducer struct {
	clientAmqp  amqp.IAmqpClient
	exchange    string
	routing_key string
}

func New(amqpc amqp.IAmqpClient) ports.UserProducer {
	return &userCreateProducer{
		clientAmqp:  amqpc,
		exchange:    "user.dx",
		routing_key: "user.create",
	}
}
