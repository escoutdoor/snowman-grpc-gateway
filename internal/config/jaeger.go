package config

import (
	"fmt"
	"net"
	"os"
)

const (
	jaegerHostEnvName = "JAEGER_HOST"
	jaegerPortEnvName = "JAEGER_PORT"
)

type jaegerConfig struct {
	host string
	port string
}

func NewJaegerConfig() (JaegerConfig, error) {
	host := os.Getenv(jaegerHostEnvName)
	if host == "" {
		return nil, fmt.Errorf("jaeger host is not defined or empty")
	}

	port := os.Getenv(jaegerPortEnvName)
	if port == "" {
		return nil, fmt.Errorf("jaeger port is not defined or empty")
	}

	return &jaegerConfig{
		host: host,
		port: port,
	}, nil
}

func (c *jaegerConfig) Addr() string {
	return net.JoinHostPort(c.host, c.port)
}
