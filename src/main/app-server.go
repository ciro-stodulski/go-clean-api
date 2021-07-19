package app

import (
	"fmt"
	"go-api/src/main/container"
	http "go-api/src/main/module/http"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
)

type Server struct {
	Container *container.Container
	engine    *gin.Engine
	db        *gorm.DB
}

func (server *Server) Setup() *Server {
	server.Container = container.NewContainer(
		container.NewContainerConfig(server.db),
	)

	server.engine = gin.New()

	server.engine.Use(gin.Recovery())

	http.SetupRoutes(server.engine, server.Container)
	return server
}

func (server *Server) Start() error {
	port := os.Getenv("HOST_PORT")
	fmt.Println("server listening on localhost:", port)
	return server.engine.Run("localhost:" + port)
}

func (server *Server) CloseDB() {
	server.db.Close()
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
