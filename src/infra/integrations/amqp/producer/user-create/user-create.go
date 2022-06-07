package usercreateproducer

import (
	"encoding/json"
	typesclient "go-api/src/infra/integrations/amqp/client/types"
	create_dto "go-api/src/presentation/http/controllers/v1/users/create/dto"
)

func (ucp *userCreateProducer) CreateUser(dto create_dto.CreateDto) error {
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
