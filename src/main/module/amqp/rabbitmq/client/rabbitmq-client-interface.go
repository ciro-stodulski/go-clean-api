package amqp_client

import types_client "go-api/src/main/module/amqp/rabbitmq/client/types"

type IAmqpClient interface {
	Publish(body []byte, config types_client.ConfigAmqpClient) error
}
