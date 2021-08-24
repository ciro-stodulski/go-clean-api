package rabbitmq

import (
	"encoding/json"
	"go-api/src/main/container"
	amqp_server "go-api/src/main/module/amqp"
	consumer_type "go-api/src/presentation/amqp/consumers"
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
	constumers := rabbit_mq.LoadConsumers(rabbit_mq.container)

	for i := 0; i < len(constumers); i++ {
		go rabbit_mq.StartConsumers(constumers, i)
	}
}

func (rabbit_mq *RabbitMq) StartConsumers(constumers []consumer_type.Comsumer, i int) {
	conn, err_connection := amqp.Dial("amqp://admin:admin@localhost:5672/")

	rabbit_mq.NeedToReconnect(err_connection, "Failed to connect to RabbitMQ")
	defer conn.Close()

	rabbit_mq.connection = conn
	ch, err := conn.Channel()

	rabbit_mq.NeedToReconnect(err, "Failed to open a channel")
	defer ch.Close()
	rabbit_mq.channel = ch

	queue, err := rabbit_mq.channel.QueueDeclare(
		constumers[i].GetQueue(), // name
		false,                    // durable
		false,                    // delete when unused
		false,                    // exclusive
		false,                    // no-wait
		nil,                      // arguments
	)
	failOnError(err, "Failed to declare a queue")

	msgs, err := rabbit_mq.channel.Consume(
		queue.Name, // queue
		queue.Name, // consumer
		true,       // auto-ack
		false,      // exclusive
		false,      // no-local
		false,      // no-wait
		nil,        // args
	)
	failOnError(err, "Failed to register a consumer")

	log.Default().Print("RabbitMq: Started queue " + queue.Name + " to consume")

	for msg := range msgs {
		schema := constumers[i].GetSchema()
		err := json.Unmarshal(msg.Body, &schema)

		if err != nil {
			log.Printf("Error decoding JSON: %s", err)
			if err := msg.Ack(false); err != nil {
				log.Println("unable to acknowledge the message, dropped", err)
			}
			rabbit_mq.NeedToReconnect(err, "ack message")
		} else {
			err_msg_consumer := constumers[i].MessageHandler(consumer_type.Message{
				Body: schema,
			})

			if err_msg_consumer != nil {
				constumers[i].OnConsumerError(err_msg_consumer)
				if err := msg.Ack(false); err != nil {
					log.Println("unable to acknowledge the message, dropped", err)
				}
				rabbit_mq.NeedToReconnect(err, "ack message")
			}
		}
	}
}

func (rabbit_mq *RabbitMq) NeedToReconnect(err error, msg string) {
	if err != nil {
		log.Default().Printf("%s: %s", msg, err)
		time.Sleep(2 * time.Second)
		rabbit_mq.Start()
	}
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}
