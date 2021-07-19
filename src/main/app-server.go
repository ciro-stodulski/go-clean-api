package app

import (
	"fmt"
	"go-api/src/main/config"
	"go-api/src/main/container"
	http "go-api/src/main/module/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type Server struct {
	Config    *config.Config
	Container *container.Container
	engine    *gin.Engine
	db        *gorm.DB
}

func New() (s *Server, err error) {
	s = &Server{}

	s.Config, err = config.InitConfig()
	return
}

func (s *Server) Setup() *Server {
	s.Container = container.NewContainer(
		container.NewContainerConfig(s.db),
	)

	s.engine = gin.New()

	s.engine.Use(gin.Recovery())

	http.SetupRoutes(s.engine, s.Container)
	return s
}

func (s *Server) Start() error {
	fmt.Println("server listening on localhost:", s.Config.Port)
	return s.engine.Run(fmt.Sprintf(":%d", s.Config.Port))
}

func (s *Server) CloseDB() {
	s.db.Close()
}
