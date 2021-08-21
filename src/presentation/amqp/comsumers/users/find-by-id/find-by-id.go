package find_by_id

import (
	"fmt"
	comsumer "go-api/src/presentation/amqp/comsumers"

	"github.com/mitchellh/mapstructure"
)

func (findByIdConsumer *findByIdConsumer) MessageHandler(msg comsumer.Message) error {
	dto := FindByIdDto{}
	mapstructure.Decode(msg.Body, &dto)

	fmt.Println(findByIdConsumer)

	//	findByIdConsumer.container.ListUsersUseCase.ListUsers()

	return nil
}

func (findByIdConsumer *findByIdConsumer) OnConsumerError(err error) {

}
