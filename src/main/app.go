package app

import (
	"go-api/src/main/container"
	database "go-api/src/main/module/db"
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

	gin.SetMode(gin.ReleaseMode)

	server.engine.Use(gin.Recovery())

	http.SetupRoutes(server.engine, server.Container)
	return server
}

func (server *Server) Start() error {
	port := os.Getenv("HOST_PORT")

	log.Default().Print("server started with succeffully")

	return server.engine.Run("localhost:" + port)
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
