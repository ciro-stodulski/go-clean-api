package app

import (
	"go-api/src/main/container"
	"go-api/src/main/modules/amqp"
	"go-api/src/main/modules/grpc"
	"go-api/src/main/modules/http"
	"go-api/src/main/modules/work"
	"go-api/src/shared/env"
)

type App struct {
}

func (server *App) start() error {
	c := container.New()

	go amqp.New(c).Start()
	go grpc.New(c).Start()
	work.New(c).Start()
	http.New(c).Start()

	return nil
}

func New() error {
	env.Load()

	server := &App{}

	err := server.start()

	return err
}
