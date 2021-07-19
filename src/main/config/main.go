package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	// The port to bind the web app
	Port int

	// The request/response body limit of the api
	BodyLimit string
}

func InitConfig() (config *Config, err error) {
	config = &Config{
		Port:      viper.GetInt("api.port"),
		BodyLimit: viper.GetString("api.bodyLimit"),
	}

	return
}
