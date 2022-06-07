package amqp_client

import types_client "go-api/src/infra/integrations/amqp/client/types"

type IAmqpClient interface {
	Publish(body []byte, config types_client.ConfigAmqpClient) error
}
