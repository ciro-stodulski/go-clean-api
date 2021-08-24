package comsumer

type Message struct {
	Body interface{}
}

type Comsumer interface {
	MessageHandler(Message) error
	OnConsumerError(error) error
	GetQueue() string
	GetSchema() interface{}
}
