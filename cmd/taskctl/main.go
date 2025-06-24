package main

import (
	"os"
	"taskctl/internal/config"
	"taskctl/internal/storage/httpserver"

	"github.com/rs/zerolog"
)

var logger zerolog.Logger

func init() {
	logger = zerolog.New(os.Stderr).With().Timestamp().Logger()
}

func main() {
	if err := run(); err != nil {
		logger.Fatal().Err(err)
	}
}

func run() error {
	cfg, err := config.Read("config.json")
	if err != nil {
		return err
	}

	srv := httpserver.NewServer(cfg.Server)
	if err := srv.Start(); err != nil {
		return err
	}

	return nil
}
