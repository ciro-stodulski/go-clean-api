package app

import (
	"go-api/src/main/container"
	database "go-api/src/main/module/db"
	http_server "go-api/src/main/module/http"
	"log"

	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
)

type Server struct {
	Container *container.Container
	engine    http_server.HttpServer
	db        *gorm.DB
}

func (server *Server) Setup() *Server {
	server.Container = container.NewContainer(
		container.NewContainerConfig(server.db),
	)

	server.engine.New(server.Container)

	return server
}

func (server *Server) CloseDB() {
	server.db.Close()
}

func (s *Server) ConnectToDabase() error {
	db, err := createDatabase()

	if err != nil {
		return err
	}

	s.db = db

	log.Default().Print("connection db with succeffully")

	database.LoadMigrationByRepositores(db)

	return nil
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

func createDatabase() (db *gorm.DB, err error) {
	dbConfig, err := database.NewDbConfig()

	if err == nil && dbConfig != nil {
		db, err = database.GetDatabase(dbConfig)
	}

	return db, err
}
