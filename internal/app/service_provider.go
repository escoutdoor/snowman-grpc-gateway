package app

import (
	"crypto/tls"
	"crypto/x509"
	"os"

	"github.com/bufbuild/protovalidate-go"
	snowman_implementation "github.com/escoutdoor/snowman-grpc-gateway/internal/api/snowman/v1"
	"github.com/escoutdoor/snowman-grpc-gateway/internal/config"
	"github.com/escoutdoor/snowman-grpc-gateway/internal/logger"
)

type serviceProvider struct {
	grpcServerConfig    config.GrpcServerConfig
	gatewayServerConfig config.GatewayServerConfig
	swaggerServerConfig config.SwaggerServerConfig

	tlsConfig *tls.Config

	validator protovalidate.Validator

	snowmanImplementation *snowman_implementation.Implementation
}

func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

func (s *serviceProvider) GrpcServerConfig() config.GrpcServerConfig {
	if s.grpcServerConfig == nil {
		cfg, err := config.NewGrpcServerConfig()
		if err != nil {
			logger.Logger().Fatalf("failed to load grpc server config: %s", err)
		}

		s.grpcServerConfig = cfg
	}

	return s.grpcServerConfig
}

func (s *serviceProvider) GatewayServerConfig() config.GatewayServerConfig {
	if s.gatewayServerConfig == nil {
		cfg, err := config.NewGatewayServerConfig()
		if err != nil {
			logger.Logger().Fatalf("failed to load gateway server config: %s", err)
		}

		s.gatewayServerConfig = cfg
	}
	return s.gatewayServerConfig
}

func (s *serviceProvider) SwaggerServerConfig() config.SwaggerServerConfig {
	if s.swaggerServerConfig == nil {
		cfg, err := config.NewSwaggerServerConfig()
		if err != nil {
			logger.Logger().Fatalf("failed to load swagger server config: %s", err)
		}

		s.swaggerServerConfig = cfg
	}

	return s.swaggerServerConfig
}

func (s *serviceProvider) Validator() protovalidate.Validator {
	if s.validator == nil {
		validator, err := protovalidate.New()
		if err != nil {
			logger.Logger().Fatalf("failed to init protovalidate validator: %s", err)
		}

		s.validator = validator
	}

	return s.validator
}

func (s *serviceProvider) SnowmanImplementation() *snowman_implementation.Implementation {
	if s.snowmanImplementation == nil {
		s.snowmanImplementation = snowman_implementation.NewImplementation()
	}

	return s.snowmanImplementation
}

func (s *serviceProvider) TLSConfig() *tls.Config {
	if s.tlsConfig == nil {
		caCert, err := os.ReadFile(caCertFilePath)
		if err != nil {
			logger.Logger().Fatalf("failed to load ca certificate: %s", err)
		}

		serverCert, err := tls.LoadX509KeyPair(serverCertFilePath, serverKeyFilePath)
		if err != nil {
			logger.Logger().Fatalf("failed to load server certificate and key: %s", err)
		}

		certPool := x509.NewCertPool()
		if !certPool.AppendCertsFromPEM(caCert) {
			logger.Logger().Fatalf("failed to append ca certificate to certificate pool: %s", err)
		}

		tlsConfig := &tls.Config{
			Certificates: []tls.Certificate{serverCert},
			ClientCAs:    certPool,
			ClientAuth:   tls.RequireAndVerifyClientCert,
		}
		s.tlsConfig = tlsConfig
	}

	return s.tlsConfig
}
