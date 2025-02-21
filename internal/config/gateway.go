package config

import (
	"fmt"
	"net"
	"os"
)

const (
	gatewayServerPortEnvName = "GATEWAY_SERVER_PORT"
	gatewayServerHostEnvName = "GATEWAY_SERVER_HOST"
)

type gatewayServerConfig struct {
	host string
	port string
}

func NewGatewayServerConfig() (GatewayServerConfig, error) {
	host := os.Getenv(gatewayServerHostEnvName)
	if host == "" {
		return nil, fmt.Errorf("gateway server host is not defined or empty")
	}

	port := os.Getenv(gatewayServerPortEnvName)
	if port == "" {
		return nil, fmt.Errorf("gateway server port is not defined or empty")
	}

	return &gatewayServerConfig{
		host: host,
		port: port,
	}, nil
}

func (c *gatewayServerConfig) Addr() string {
	return net.JoinHostPort(c.host, c.port)
}
