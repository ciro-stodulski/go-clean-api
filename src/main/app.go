package app

import (
	"go-api/src/main/container"
	database "go-api/src/main/module/db/mysql"
	http_server "go-api/src/main/module/http/server"
	"go-api/src/main/module/work"
	"log"

	"github.com/joho/godotenv"
)

type Server struct {
	Container *container.Container
	engine    http_server.HttpServer
	db        database.Database
}

func (server *Server) Setup() *Server {
	server.Container = container.NewContainer(
		container.NewContainerConfig(server.db.Db),
	)
	work := work.New(server.Container)

	work.StartCrons()

	server.engine.New(server.Container)

	return server
}

func New() (server *Server, err error) {
	server = &Server{}
	InitEnvs()
	return
}

func InitEnvs() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}
