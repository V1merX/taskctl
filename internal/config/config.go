package config

import (
	"encoding/json"
	"fmt"
	"os"
)

const (
	configDir = "configs/"
)

type Config struct {
	Server HTTPServer `json:"http-server"`
}

type HTTPServer struct {
	Host string `json:"host"`
	Port int8   `json:"port"`
}

func Read(filename string) (*Config, error) {
	const op = "internal.config.Read"

	var cfg Config

	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	err = json.Unmarshal(data, &cfg)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &cfg, nil
}
