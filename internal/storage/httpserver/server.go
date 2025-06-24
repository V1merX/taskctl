package httpserver

import (
	"fmt"
	"taskctl/internal/config"

	"github.com/gin-gonic/gin"
)

type Server struct {
	Host string
	Port int8
}

func NewServer(srvOptions config.HTTPServer) *Server {
	return &Server{
		Host: srvOptions.Host,
		Port: srvOptions.Port,
	}
}

func (s *Server) Start() error {
	const op = "internal.httpserver.Server.Start"

	router := gin.Default()

	// TODO: add routes
	// router.POST("/tasks", createTaskHandler)
	// router.GET("/tasks/:id", getTaskHandler)
	// router.DELETE("/tasks/:id", deleteTaskHandler)

	if err := router.Run(fmt.Sprintf("%s:%d", s.Host, s.Port)); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}
