package amqpclient

import typesclient "go-api/cmd/infra/integrations/amqp/client/types"

type AmqpClient interface {
	Publish(b []byte, c typesclient.ConfigAmqpClient) error
}
