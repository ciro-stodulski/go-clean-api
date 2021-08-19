package comsumer

type Comsumer interface {
	MessageHandler()
	OnConsumerError()
}
