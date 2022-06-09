package app

import (
	"go-api/src/main/container"
	rabbitmq "go-api/src/main/module/amqp/server"
	database "go-api/src/main/module/db/mysql"
	grpc_server "go-api/src/main/module/grpc"
	http_server "go-api/src/main/module/http"
	"go-api/src/main/module/work"
	"log"

	"github.com/joho/godotenv"
)

type Server struct {
	Container *container.Container
	http      http_server.HttpServer
	db        database.Database
	amqp      rabbitmq.RabbitMq
	grpc      grpc_server.GRPCServer
}

func (server *Server) Setup() *Server {
	server.Container = container.New(
		container.NewConfig(server.db.Db),
	)

	work.New(server.Container).StartCrons()

	go server.amqp.New(server.Container).Start()
	go server.grpc.New(server.Container).Start()

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
