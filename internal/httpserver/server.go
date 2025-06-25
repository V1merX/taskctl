package httpserver

import (
	"fmt"
	"taskctl/internal/config"
	"taskctl/internal/storage"
)

type Server struct {
	settings *config.HTTPServer
	db       *storage.Storage
}

func NewServer(db *storage.Storage, srvOptions config.HTTPServer) *Server {
	return &Server{db: db, settings: &srvOptions}
}

func (s *Server) Start() error {
	const op = "internal.httpserver.Server.Start"

	router := s.setupRoutes()

	if err := router.Run(fmt.Sprintf("%s:%d", s.settings.Host, s.settings.Port)); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}
