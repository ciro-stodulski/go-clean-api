package amqp

import (
	"encoding/json"
	"go-api/cmd/infra/adapters/rabbitmq"
	consumer_type "go-api/cmd/interface/amqp/consumers"
	ports_amqp "go-api/cmd/interface/amqp/ports"
	"go-api/cmd/main/container"
	"go-api/cmd/main/modules"

	"log"
	"time"

	"github.com/streadway/amqp"
)

type amqpModule struct {
	container *container.Container
	channel   *amqp.Channel
}

func New(container *container.Container) modules.Module {
	return &amqpModule{container: container}
}

func (am *amqpModule) RunGo() bool {
	return true
}

func (am *amqpModule) Stop() {}

func (am *amqpModule) Start() error {
	am.channel = rabbitmq.GetChanel()

	constumers := am.LoadConsumers(am.container)

	for i := 0; i < len(constumers); i++ {
		go am.StartConsumers(constumers, i)
	}

	return nil
}

func (am *amqpModule) StartConsumers(constumers []consumer_type.Comsumer, position int) {
	queue, err := am.channel.QueueDeclarePassive(
		constumers[position].GetQueue(), // name
		false,                           // durable
		false,                           // delete when unused
		false,                           // exclusive
		false,                           // no-wait
		nil,                             // arguments
	)

	failOnError(err, "Failed to declare a queue")

	msgs, err := am.channel.Consume(
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
			am.NeedToReconnect(err, "ack message")
		} else {
			err_msg_consumer := constumers[position].MessageHandler(ports_amqp.Message{
				Body: schema,
			})

			if err_msg_consumer != nil {
				err_consumer := constumers[position].OnConsumerError(err_msg_consumer)
				if err := msg.Ack(false); err != nil {
					log.Println("unable to acknowledge the message, dropped", err)
				}
				am.NeedToReconnect(err_consumer, "ack message")
			}
		}
	}
}

func (am *amqpModule) NeedToReconnect(err error, msg string) {
	if err != nil {
		log.Default().Printf("%s: %s", msg, err)
		time.Sleep(2 * time.Second)
		am.Start()
	}
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}
