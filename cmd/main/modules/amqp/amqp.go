package amqp

import (
	"encoding/json"
	rabbitmqadapter "go-clean-api/cmd/infra/adapters/rabbitmq"
	consumer_type "go-clean-api/cmd/interface/amqp/consumers"
	ports_amqp "go-clean-api/cmd/interface/amqp/ports"
	"go-clean-api/cmd/main/container"
	"go-clean-api/cmd/main/modules"

	"log"
	"time"

	"github.com/isayme/go-amqp-reconnect/rabbitmq"
)

type amqpModule struct {
	container *container.Container
	channel   *rabbitmq.Channel
}

func New(container *container.Container) modules.Module {
	return &amqpModule{container: container}
}

func (am *amqpModule) RunGo() bool {
	return true
}

func (am *amqpModule) Stop() {}

func (am *amqpModule) Start() error {
	am.channel = rabbitmqadapter.GetChanel()

	constumers := am.LoadConsumers(am.container)

	for i := 0; i < len(constumers); i++ {
		go am.StartConsumers(constumers, i)
	}

	return nil
}

func (am *amqpModule) StartConsumers(constumers []consumer_type.Comsumer, position int) {

	queue, err := am.channel.QueueDeclarePassive(
		constumers[position].GetConfig().Queue, // name
		true,                                   // durable
		false,                                  // delete when unused
		false,                                  // exclusive
		false,                                  // no-wait
		nil,                                    // arguments
	)

	failOnError(err, "Failed to declare a queue")

	msgs, err := am.channel.Consume(
		queue.Name, // queue
		queue.Name, // consumer
		false,      // auto-ack
		false,      // exclusive
		false,      // no-local
		false,      // no-wait
		nil,        // args
	)
	failOnError(err, "Failed to register a consumer")

	log.Default().Print("RabbitMq: Started queue " + queue.Name + " to consume")

	for msg := range msgs {
		schema := constumers[position].GetConfig().Schema
		err := json.Unmarshal(msg.Body, &schema)

		if err != nil {
			log.Printf("Error decoding JSON: %s", err)

			msg.Ack(true)
		} else {
			err_msg_consumer := constumers[position].MessageHandler(ports_amqp.Message{
				Body: schema,
			})

			if err_msg_consumer != nil {
				shoudl_ack := constumers[position].OnConsumerError(err_msg_consumer)

				if shoudl_ack.Requeue {
					msg.Nack(shoudl_ack.Multiple, true)
				} else {
					msg.Ack(shoudl_ack.Multiple)
				}
			} else {
				msg.Ack(true)
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
