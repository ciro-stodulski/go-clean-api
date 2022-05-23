package user_create

import (
	"encoding/json"
	"fmt"
	types_client "go-api/src/main/module/amqp/rabbitmq/client/types"
	create_dto "go-api/src/presentation/http/controllers/v1/users/create/dto"
)

func (userCreate *userCreate) CreateUser(dto create_dto.CreateDto) error {
	config := types_client.ConfigAmqpClient{
		Exchange:    userCreate.exchange,
		Routing_key: userCreate.routing_key,
	}

	btResult, _ := json.Marshal(&dto)

	err := userCreate.clientAmqp.Publish(
		[]byte(string(btResult)),
		config,
	)
	fmt.Println(err)
	return err
}
