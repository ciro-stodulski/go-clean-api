package amqp_helper

import (
	"go-api/src/shared/env"
)

func GetConnection() string {
	return env.Env().RabbitMqProtocol + "://" +
		env.Env().RabbitMqUser +
		":" +
		env.Env().RabbitMqPassword +
		"@" + env.Env().RabbitMqHost +
		":" + env.Env().RabbitMqPort +
		"/" + env.Env().RabbitMqVHost
}
