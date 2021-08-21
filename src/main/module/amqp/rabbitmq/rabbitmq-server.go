package rabbitmq

import (
	"encoding/json"
	"go-api/src/main/container"
	amqp_server "go-api/src/main/module/amqp"
	comsumer "go-api/src/presentation/amqp/comsumers"
	"log"
	"time"

	"github.com/streadway/amqp"
)

type RabbitMq struct {
	container  *container.Container
	connection *amqp.Connection
	channel    *amqp.Channel
}

func (rabbit_mq *RabbitMq) New(container *container.Container) amqp_server.AmqpServer {
	return &RabbitMq{container: container}
}

func (rabbit_mq *RabbitMq) Start() {
	conn, err_connection := amqp.Dial("amqp://admin:admin@localhost:5672/")

	rabbit_mq.NeedToReconnect(err_connection, "Failed to connect to RabbitMQ")
	defer conn.Close()

	rabbit_mq.connection = conn
	ch, err := conn.Channel()

	rabbit_mq.NeedToReconnect(err, "Failed to open a channel")
	defer ch.Close()

	rabbit_mq.channel = ch
	log.Default().Print("RabbitMq: Connection host and channel with succeffully")

	rabbit_mq.StartConsumers(rabbit_mq.container)
}

func (rabbit_mq *RabbitMq) StartConsumers(container *container.Container) {
	for _, consumer := range rabbit_mq.LoadConsumers(container) {
		queue, err := rabbit_mq.channel.QueueDeclare(
			consumer.GetQueue(), // name
			false,               // durable
			false,               // delete when unused
			false,               // exclusive
			false,               // no-wait
			nil,                 // arguments
		)
		failOnError(err, "Failed to declare a queue")

		rabbit_mq.channel.Qos(1, 0, false)

		msgs, err := rabbit_mq.channel.Consume(
			queue.Name, // queue
			"",         // consumer
			true,       // auto-ack
			false,      // exclusive
			false,      // no-local
			false,      // no-wait
			nil,        // args
		)
		failOnError(err, "Failed to register a consumer")

		log.Default().Print("RabbitMq: Started queue " + queue.Name + " to consume")
		for msg := range msgs {
			schema := consumer.GetSchema()
			err := json.Unmarshal(msg.Body, &schema)

			if err != nil {
				log.Printf("Error decoding JSON: %s", err)
			} else {
				err_msg_consumer := consumer.MessageHandler(comsumer.Message{
					Body: schema,
				})

				if err_msg_consumer != nil {
					consumer.OnConsumerError(err_msg_consumer)
					msg.Nack(true, true)
				}
			}
		}
	}
}

func (rabbit_mq *RabbitMq) NeedToReconnect(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		time.Sleep(10 * time.Second)
		rabbit_mq.Start()
	}
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}
