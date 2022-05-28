package app

import (
	"go-api/src/main/container"
	rabbitmq "go-api/src/main/module/amqp/rabbitmq/server"
	database "go-api/src/main/module/db/mysql"
	grpc_server "go-api/src/main/module/grpc"
	http_server "go-api/src/main/module/http/server"
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
	server.Container = container.NewContainer(
		container.NewContainerConfig(server.db.Db),
	)

	// descomentar depois que concluir tarefa para o grpc
	//work.New(server.Container).StartCrons()

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
