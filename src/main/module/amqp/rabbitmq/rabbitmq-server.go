package rabbitmq

import (
	"go-api/src/main/container"
	amqp_server "go-api/src/main/module/amqp"
	"log"
	"time"

	"github.com/streadway/amqp"
)

type rabbitMq struct {
	container container.Container
}

func New(c container.Container) amqp_server.AmqpServer {
	return &rabbitMq{c}
}

func (rabbit_mq *rabbitMq) Start() {
	conn, err := amqp.Dial("amqp://admin:admin@localhost:5672/")

	if err != nil {
		log.Default().Print("Failed to open a RabbitMQ")
		log.Default().Print(err)
		defer conn.Close()
	}

	ch, err := conn.Channel()

	if err != nil {
		log.Default().Print("Failed to open a channel")
		log.Default().Print(err)
		defer ch.Close()
	}

}

func (rabbit_mq *rabbitMq) LoadConsumers() {

}

func (rabbit_mq *rabbitMq) StartConsumers() {

}

func (rabbit_mq *rabbitMq) Reconnect() {
	time.Sleep(10 * time.Second)
	rabbit_mq.Start()
}
