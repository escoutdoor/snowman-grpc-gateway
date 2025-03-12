package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"os"

	"github.com/escoutdoor/snowman-grpc-gateway/internal/logger"
	pb "github.com/escoutdoor/snowman-grpc-gateway/pkg/snowman/v1"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

const (
	caCertFilePath = "./certificate/ca.crt"
	certFilePath   = "./certificate/client.crt"
	keyFilePath    = "./certificate/client.key"
)

func main() {
	ctx := context.Background()
	logger.SetLevel(zap.DebugLevel)

	// Load the CA certificate
	caCert, err := os.ReadFile(caCertFilePath)
	if err != nil {
		logger.Fatalf(ctx, "failed to load ca certificate: %s", err)
	}

	clientCert, err := tls.LoadX509KeyPair(certFilePath, keyFilePath)
	if err != nil {
		logger.Fatalf(ctx, "failed to load client certificate and key: %s", err)
	}

	certPool := x509.NewCertPool()
	if !certPool.AppendCertsFromPEM(caCert) {
		logger.Fatalf(ctx, "failed to append ca certificate to certificate pool: %s", err)
	}

	// Create TLS configuration with the certificate of the server
	tlsConfig := &tls.Config{
		Certificates:       []tls.Certificate{clientCert},
		RootCAs:            certPool,
		InsecureSkipVerify: false,
	}
	creds := credentials.NewTLS(tlsConfig)

	conn, err := grpc.NewClient("localhost:9010", grpc.WithTransportCredentials(creds))
	if err != nil {
		logger.Fatalf(ctx, "failed to init grpc server connection: %s", err)
	}
	defer conn.Close()

	client := pb.NewSnowmanServiceV1Client(conn)

	req := &pb.BuildSnowmanRequest{
		Name:   "snowman-1",
		Height: 123,
		Width:  123,
	}

	resp, err := client.Build(ctx, req)
	if err != nil {
		logger.ErrorKV(ctx, "build snowman", "error", err)
	}

	logger.DebugKV(ctx, "snowman just got built", "id", resp.Id)

	listResp, err := client.List(ctx, &pb.ListSnowmenRequest{})
	if err != nil {
		logger.ErrorKV(ctx, "list snowmen", "error", err)
	}
	logger.DebugKV(ctx, "snowmen list request", "list", listResp.Snowmen)
}
