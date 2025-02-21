package config

import (
	"fmt"
	"net"
	"os"
)

const (
	grpcServerPortEnvName = "GRPC_SERVER_PORT"
	grpcServerHostEnvName = "GRPC_SERVER_HOST"
)

type grpcServerConfig struct {
	host string
	port string
}

func NewGrpcServerConfig() (GrpcServerConfig, error) {
	host := os.Getenv(grpcServerHostEnvName)
	if host == "" {
		return nil, fmt.Errorf("grpc server host is not defined or empty")
	}

	port := os.Getenv(grpcServerPortEnvName)
	if port == "" {
		return nil, fmt.Errorf("grpc server port is not defined or empty")
	}

	return &grpcServerConfig{
		host: host,
		port: port,
	}, nil
}

func (c *grpcServerConfig) Addr() string {
	return net.JoinHostPort(c.host, c.port)
}
