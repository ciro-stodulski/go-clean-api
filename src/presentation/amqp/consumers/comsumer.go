package comsumer

import ports_amqp "go-api/src/presentation/amqp/ports"

type Comsumer interface {
	MessageHandler(ports_amqp.Message) error
	OnConsumerError(error) error
	GetQueue() string
	GetSchema() interface{}
}
