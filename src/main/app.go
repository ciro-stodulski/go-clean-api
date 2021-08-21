package app

import (
	"go-api/src/main/container"
	"go-api/src/main/module/amqp/rabbitmq"
	database "go-api/src/main/module/db/mysql"
	http_server "go-api/src/main/module/http/server"
	"go-api/src/main/module/work"
	"log"

	"github.com/joho/godotenv"
)

type Server struct {
	Container *container.Container
	http      http_server.HttpServer
	db        database.Database
	amqp      rabbitmq.RabbitMq
}

func (server *Server) Setup() *Server {
	server.Container = container.NewContainer(
		container.NewContainerConfig(server.db.Db),
	)

	work.New(server.Container).StartCrons()

	go server.amqp.Start(server.Container)

	server.http.New(server.Container)

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
