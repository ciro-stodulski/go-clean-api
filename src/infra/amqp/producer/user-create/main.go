package user_create

import "go-api/src/infra/amqp/producer"

type userCreate struct {
	channel interface{}
}

func NewProdocer() producer.Producer {
	return &userCreate{}
}
