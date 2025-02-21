package config

import (
	"fmt"
	"net"
	"os"
)

const (
	swaggerServerPortEnvName = "SWAGGER_SERVER_PORT"
	swaggerServerHostEnvName = "SWAGGER_SERVER_HOST"
	swaggerFilepathEnvName   = "SWAGGER_FILEPATH"
)

type swaggerServerConfig struct {
	host     string
	port     string
	filepath string
}

func NewSwaggerServerConfig() (SwaggerServerConfig, error) {
	host := os.Getenv(swaggerServerHostEnvName)
	if host == "" {
		return nil, fmt.Errorf("swagger server host is not defined or empty")
	}

	port := os.Getenv(swaggerServerPortEnvName)
	if port == "" {
		return nil, fmt.Errorf("swagger server port is not defined or empty")
	}

	filepath := os.Getenv(swaggerFilepathEnvName)
	if filepath == "" {
		return nil, fmt.Errorf("swagger file path is not defined or empty")
	}
	if _, err := os.Stat(filepath); os.IsNotExist(err) {
		return nil, fmt.Errorf("swagger file does not exist")
	}

	return &swaggerServerConfig{
		host:     host,
		port:     port,
		filepath: filepath,
	}, nil
}

func (c *swaggerServerConfig) Addr() string {
	return net.JoinHostPort(c.host, c.port)
}

func (c *swaggerServerConfig) Filepath() string {
	return c.filepath
}
