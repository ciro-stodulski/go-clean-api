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

type Consumer struct {
	channel *amqp.Channel
	deliveries <-chan amqp.Delivery
	handler func(amqp.Delivery)
	done chan error
	session Session
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

	rabbit_mq.StartConsumers()
}

func (rabbit_mq *RabbitMq) StartConsumers() {

	constumers := rabbit_mq.LoadConsumers(rabbit_mq.container)
	err := rabbit_mq.channel.Qos(
		len(constumers), // prefetch count
		0,               // prefetch size
		false,           // global
	)
	failOnError(err, "Failed to set QoS")

	for _, consumer := range constumers {
		queue, err := rabbit_mq.channel.QueueDeclare(
			consumer.GetQueue(), // name
			false,               // durable
			false,               // delete when unused
			false,               // exclusive
			false,               // no-wait
			nil,                 // arguments
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
			schema := consumer.GetSchema()
			err := json.Unmarshal(msg.Body, &schema)

			if err != nil {
				log.Printf("Error decoding JSON: %s", err)
				if err := msg.Ack(false); err != nil {
					// TODO: Should DLX the message
					log.Println("unable to acknowledge the message, dropped", err)
				}
				rabbit_mq.NeedToReconnect(err, "ack message")
			} else {
				err_msg_consumer := consumer.MessageHandler(consumer_type.Message{
					Body: schema,
				})

				if err_msg_consumer != nil {
					consumer.OnConsumerError(err_msg_consumer)
					if err := msg.Ack(false); err != nil {
						// TODO: Should DLX the message
						log.Println("unable to acknowledge the message, dropped", err)
					}
					rabbit_mq.NeedToReconnect(err, "ack message")
				}
			}
		}
	}
}

func 

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
