package main

import (
	"os"

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

	return nil
}
