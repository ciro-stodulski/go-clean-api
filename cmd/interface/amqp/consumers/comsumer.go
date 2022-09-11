package cosumer

import ports_amqp "go-clean-api/cmd/interface/amqp/ports"

type (
	ConsumeConfig struct {
		Queue  string
		Schema interface{}
	}

	AckConfig struct {
		Multiple bool
		Requeue  bool
	}
	Comsumer interface {
		GetConfig() ConsumeConfig
		MessageHandler(ports_amqp.Message) error
		OnConsumerError(error) AckConfig
	}
)
