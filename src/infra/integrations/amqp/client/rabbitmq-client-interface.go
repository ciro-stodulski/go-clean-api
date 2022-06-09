package amqpclient

import typesclient "go-api/src/infra/integrations/amqp/client/types"

type AmqpClient interface {
	Publish(b []byte, c typesclient.ConfigAmqpClient) error
}
