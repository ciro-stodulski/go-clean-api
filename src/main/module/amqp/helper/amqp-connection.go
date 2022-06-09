package amqp_helper

import (
	"os"
)

func GetConnection() string {
	return os.Getenv("RABBIT_MQ_PROTOCOL") +
		"://" +
		os.Getenv("RABBIT_MQ_USERNAME") +
		":" +
		os.Getenv("RABBIT_MQ_PASSWORD") +
		"@" + os.Getenv("RABBIT_MQ_HOST") +
		":" + os.Getenv("RABBIT_MQ_PORT") +
		"/" + os.Getenv("RABBIT_MQ_VHOST")
}
