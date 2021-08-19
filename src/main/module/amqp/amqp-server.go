package amqp_server

type AmqpServer interface {
	Start()
	StartConsumers()
	Reconnect()
	LoadConsumers()
}
