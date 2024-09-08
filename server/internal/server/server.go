package server

import (
	"github.com/Sergio-Na/argus/server/config"
	"github.com/Sergio-Na/argus/server/internal/handler"
	"github.com/gin-gonic/gin"
)

type Server struct {
	router *gin.Engine
	config *config.Config
}

func New(cfg *config.Config) (*Server, error) {
	r := gin.Default()
	s := &Server{
		router: r,
		config: cfg,
	}

	s.registerRoutes()
	return s, nil
}

func (s *Server) registerRoutes() {
	s.router.GET("/", handler.Home)
}

func (s *Server) Run() error {
	return s.router.Run(s.config.ServerAddress)
}
