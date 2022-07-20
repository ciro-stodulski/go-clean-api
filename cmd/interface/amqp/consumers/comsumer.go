package comsumer

import ports_amqp "go-api/cmd/interface/amqp/ports"

type Comsumer interface {
	MessageHandler(ports_amqp.Message) error
	OnConsumerError(error) error
	GetQueue() string
	GetSchema() interface{}
}
