package cosumer

import ports_amqp "go-api/cmd/interface/amqp/ports"

type (
	ConsumeConfig struct {
		Queue  string
		Schema interface{}
	}
	Comsumer interface {
		GetConfig() ConsumeConfig
		MessageHandler(ports_amqp.Message) error
		OnConsumerError(error) error
	}
)
