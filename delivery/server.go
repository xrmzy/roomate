package delivery

import (
	"log"
	"roomate/config"

	"github.com/gin-gonic/gin"
)

type Server struct {
	engine *gin.Engine
	infra  config.InfraConfig
	cfg    config.Config
}

func (s *Server) Run() {
	if err := s.engine.Run(s.cfg.DB_HOST); err != nil {
		log.Fatal("server can't run")
	}
}

func NewServer() *Server {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatal(err)
	}

	engine := gin.Default()
	return &Server{
		engine: engine,
		cfg:    cfg,
	}
}
