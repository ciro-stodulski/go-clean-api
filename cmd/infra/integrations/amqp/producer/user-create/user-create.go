package usercreateproducer

import (
	"encoding/json"
	"go-api/cmd/core/ports"
	typesclient "go-api/cmd/infra/integrations/amqp/client/types"
)

func (ucp *userCreateProducer) CreateUser(dto ports.CreateDto) error {
	config := typesclient.ConfigAmqpClient{
		Exchange:    ucp.exchange,
		Routing_key: ucp.routing_key,
	}

	btresult, _ := json.Marshal(&dto)

	err := ucp.clientAmqp.Publish(
		[]byte(string(btresult)),
		config,
	)

	return err
}
