package delivery

import (
	"fmt"
	"log"
	"roomate/config"

	"github.com/gin-gonic/gin"
)

type Server struct {
	engine *gin.Engine
	host   string
}

func (s *Server) Run() {
	if err := s.engine.Run(s.host); err != nil {
		log.Fatal("server can't run")
	}
}

func NewServer() *Server {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatal(err)
	}

	engine := gin.Default()
	host := fmt.Sprintf(":%s", cfg.ApiPort)
	return &Server{
		engine: engine,
		host:   host,
	}
}
