package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func Load(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return fmt.Errorf("file does not exist: %w", err)
	}

	err := godotenv.Load(path)
	if err != nil {
		return fmt.Errorf("load config: %w", err)
	}

	return nil
}

type GrpcServerConfig interface {
	Addr() string
}

type GatewayServerConfig interface {
	Addr() string
}

type SwaggerServerConfig interface {
	Addr() string
	Filepath() string
}

type PrometheusServerConfig interface {
	Addr() string
}

type JaegerConfig interface {
	Addr() string
}
