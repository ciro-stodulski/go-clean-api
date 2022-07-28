package env

import (
	"github.com/thrcorrea/envloader"
)

type Environment struct {
	HostPort                     string `env:"HOST_PORT"`
	HostHttp                     string `env:"HOST_HTTP"`
	GrpcClientUrl                string `env:"GRPC_CLIENT_URL"`
	JsonPlaceOlderIntegrationUrl string `env:"JSON_PLACE_OLDER_INTEGRATION_URL"`
	DBDrive                      string `env:"DB_DRIVE"`
	DBHost                       string `env:"DB_HOST"`
	DBSchema                     string `env:"DB_SCHEMA"`
	DBPort                       string `env:"DB_PORT"`
	DBUser                       string `env:"DB_USER"`
	RedisPort                    string `env:"REDIS_PORT"`
	GrpcServerPort               string `env:"GRPC_SERVER_PORT"`
	RedisHost                    string `env:"REDIS_HOST"`
	DBPassword                   string `env:"DB_PASSWORD"`
	RabbitMqHost                 string `env:"RABBIT_MQ_HOST"`
	RabbitMqPort                 string `env:"RABBIT_MQ_PORT"`
	RabbitMqUser                 string `env:"RABBIT_MQ_USERNAME"`
	RabbitMqPassword             string `env:"RABBIT_MQ_PASSWORD"`
	RabbitMqVHost                string `env:"RABBIT_MQ_VHOST,optional"`
	RabbitMqProtocol             string `env:"RABBIT_MQ_PROTOCOL,optional,default=amqp"`
}

var env Environment

func Env() *Environment {
	return &env
}

func Load() {
	err := envloader.Load(&env, ".env")
	if err != nil {
		panic(err)
	}
}
