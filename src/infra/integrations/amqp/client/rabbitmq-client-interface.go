package amqpclient

import typesclient "go-api/src/infra/integrations/amqp/client/types"

type IAmqpClient interface {
	Publish(b []byte, c typesclient.ConfigAmqpClient) error
}
