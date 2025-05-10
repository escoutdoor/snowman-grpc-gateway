package config

import (
	"fmt"
	"net"
	"os"
)

const (
	prometheusServerHostEnvName = "PROMETHEUS_SERVER_HOST"
	prometheusServerPortEnvName = "PROMETHEUS_SERVER_PORT"
)

type prometheusServerConfig struct {
	host, port string
}

func NewPrometheusServerConfig() (PrometheusServerConfig, error) {
	host := os.Getenv(prometheusServerHostEnvName)
	if host == "" {
		return nil, fmt.Errorf("prometheus server host is not defined or empty")
	}

	port := os.Getenv(prometheusServerPortEnvName)
	if port == "" {
		return nil, fmt.Errorf("prometheus server port is not defined or empty")
	}

	return &prometheusServerConfig{
		host: host,
		port: port,
	}, nil
}

func (c *prometheusServerConfig) Addr() string {
	return net.JoinHostPort(c.host, c.port)
}
