package amqp

import (
	"encoding/json"
	consumer_type "go-api/src/interface/amqp/consumers"
	ports_amqp "go-api/src/interface/amqp/ports"
	"go-api/src/main/container"
	"go-api/src/main/modules"
	amqp_helper "go-api/src/main/modules/amqp/helper"

	"log"
	"time"

	"github.com/streadway/amqp"
)

type rabbitMq struct {
	container *container.Container
	channel   *amqp.Channel
}

func New(container *container.Container) modules.Module {
	return &rabbitMq{container: container}
}

func (rabbit_mq *rabbitMq) RunGo() bool {
	return true
}

func (rabbit_mq *rabbitMq) Stop() {}

func (rabbit_mq *rabbitMq) Start() error {
	conn, err_connection := amqp.Dial(
		amqp_helper.GetConnection(),
	)

	failOnError(err_connection, "Failed to connect to RabbitMQ")

	ch, err := conn.Channel()

	rabbit_mq.NeedToReconnect(err, "Failed to open a channel")

	rabbit_mq.channel = ch

	constumers := rabbit_mq.LoadConsumers(rabbit_mq.container)

	for i := 0; i < len(constumers); i++ {
		go rabbit_mq.StartConsumers(constumers, i)
	}

	return err
}

func (rabbit_mq *rabbitMq) StartConsumers(constumers []consumer_type.Comsumer, position int) {
	queue, err := rabbit_mq.channel.QueueDeclarePassive(
		constumers[position].GetQueue(), // name
		false,                           // durable
		false,                           // delete when unused
		false,                           // exclusive
		false,                           // no-wait
		nil,                             // arguments
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
		schema := constumers[position].GetSchema()
		err := json.Unmarshal(msg.Body, &schema)

		if err != nil {
			log.Printf("Error decoding JSON: %s", err)
			if err := msg.Ack(false); err != nil {
				log.Println("unable to acknowledge the message, dropped", err)
			}
			rabbit_mq.NeedToReconnect(err, "ack message")
		} else {
			err_msg_consumer := constumers[position].MessageHandler(ports_amqp.Message{
				Body: schema,
			})

			if err_msg_consumer != nil {
				err_consumer := constumers[position].OnConsumerError(err_msg_consumer)
				if err := msg.Ack(false); err != nil {
					log.Println("unable to acknowledge the message, dropped", err)
				}
				rabbit_mq.NeedToReconnect(err_consumer, "ack message")
			}
		}
	}
}

func (rabbit_mq *rabbitMq) NeedToReconnect(err error, msg string) {
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
