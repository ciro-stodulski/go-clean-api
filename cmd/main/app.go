package app

import (
	"go-clean-api/cmd/main/container"
	"go-clean-api/cmd/main/modules"
	"go-clean-api/cmd/main/modules/grpc"
	"go-clean-api/cmd/main/modules/http"
	"go-clean-api/cmd/main/modules/work"
	"go-clean-api/cmd/shared/env"
)

type App struct {
	modules []modules.Module
}

func (app *App) start() error {
	var err error

	for i := 0; i < len(app.modules); i++ {
		if app.modules[i].RunGo() {
			go app.modules[i].Start()
		} else {
			err = app.modules[i].Start()
		}
	}

	return err
}

func New() error {
	env.Load()

	c := container.New()

	app := &App{
		modules: []modules.Module{
			grpc.New(c),
			work.New(c),
			http.New(c),
		},
	}

	err := app.start()

	return err
}
