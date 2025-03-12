package app

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/escoutdoor/snowman-grpc-gateway/internal/config"
	"github.com/escoutdoor/snowman-grpc-gateway/internal/interceptor"
	"github.com/escoutdoor/snowman-grpc-gateway/internal/logger"
	pb "github.com/escoutdoor/snowman-grpc-gateway/pkg/snowman/v1"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/rs/cors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/reflection"
)

const ()

type App struct {
	grpcServer    *grpc.Server
	gatewayServer *http.Server
	swaggerServer *http.Server

	serviceProvider *serviceProvider
	configPath      string
}

func New(ctx context.Context, configPath string) (*App, error) {
	a := &App{configPath: configPath}

	if err := a.initDeps(ctx); err != nil {
		return nil, err
	}

	return a, nil
}

func (a *App) Run() error {
	wg := &sync.WaitGroup{}
	wg.Add(3)

	go func() {
		defer wg.Done()
		if err := a.runGrpcServer(); err != nil {
			logger.Logger().Fatalf("failed to run grpc server: %s", err)
		}
	}()

	go func() {
		defer wg.Done()
		if err := a.runGatewayServer(); err != nil {
			logger.Logger().Fatalf("failed to run gateway server: %s", err)
		}
	}()

	go func() {
		defer wg.Done()
		if err := a.runSwaggerServer(); err != nil {
			logger.Logger().Fatalf("failed to run swagger server: %s", err)
		}
	}()

	wg.Wait()
	return nil
}

func (a *App) initDeps(ctx context.Context) error {
	deps := []func(ctx context.Context) error{
		a.initServiceProvider,
		a.initConfig,
		a.initGrpcServer,
		a.initGatewayServer,
		a.initSwaggerServer,
	}

	for _, fn := range deps {
		err := fn(ctx)
		if err != nil {
			return err
		}
	}

	return nil
}

func (a *App) initServiceProvider(_ context.Context) error {
	a.serviceProvider = newServiceProvider()
	return nil
}

func (a *App) initConfig(_ context.Context) error {
	err := config.Load(a.configPath)
	if err != nil {
		return fmt.Errorf("failed to load config: %w", err)
	}

	return nil
}

func (a *App) initGrpcServer(_ context.Context) error {
	grpcServer := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			interceptor.LoggerUnaryServerInterceptor(),
			interceptor.ValidationUnaryServerInterceptor(a.serviceProvider.Validator()),
		),
		grpc.Creds(credentials.NewTLS(a.serviceProvider.TLSConfig())),
	)

	reflection.Register(grpcServer)
	pb.RegisterSnowmanServiceV1Server(grpcServer, a.serviceProvider.SnowmanImplementation())

	a.grpcServer = grpcServer
	return nil
}

func (a *App) initGatewayServer(ctx context.Context) error {
	mux := runtime.NewServeMux()

	caCert, err := os.ReadFile(caCertFilePath)
	if err != nil {
		return fmt.Errorf("failed to load ca certificate: %s", err)
	}

	clientCert, err := tls.LoadX509KeyPair(clientCertFilePath, clientKeyFilePath)
	if err != nil {
		return fmt.Errorf("failed to load client certificate and key: %s", err)
	}

	certPool := x509.NewCertPool()
	if !certPool.AppendCertsFromPEM(caCert) {
		return fmt.Errorf("failed to append ca certificate to certificate pool: %s", err)
	}

	tlsConfig := &tls.Config{
		ServerName:   "localhost",
		Certificates: []tls.Certificate{clientCert},
		RootCAs:      certPool,
	}

	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(credentials.NewTLS(tlsConfig)),
	}

	err = pb.RegisterSnowmanServiceV1HandlerFromEndpoint(ctx, mux, a.serviceProvider.GrpcServerConfig().Addr(), opts)
	if err != nil {
		return fmt.Errorf("failed to register snowman service handler from endpoint: %w", err)
	}

	cors := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Content-Type", "Content-Length", "Authorization"},
		AllowCredentials: true,
	})

	server := &http.Server{
		Addr:              a.serviceProvider.GatewayServerConfig().Addr(),
		Handler:           cors.Handler(mux),
		ReadHeaderTimeout: time.Second * 2,
	}

	a.gatewayServer = server
	return nil
}

func (a *App) initSwaggerServer(_ context.Context) error {
	original, err := os.ReadFile(a.serviceProvider.SwaggerServerConfig().Filepath())
	if err != nil {
		return fmt.Errorf("missing swagger.json: %w", err)
	}

	patched, err := injectHost(original, a.serviceProvider.GatewayServerConfig().Addr())
	if err != nil {
		return err
	}

	mux := http.NewServeMux()
	mux.Handle("/swagger.json", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if err != nil {
			w.Write([]byte(err.Error()))
			w.WriteHeader(500)
			return
		}

		w.Header().Set("Content-type", "application/json")
		w.WriteHeader(200)
		w.Write(patched)
	}))

	docsServer := http.FileServer(http.Dir("./swagger/dist"))

	mux.Handle("/docs/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/docs" || strings.HasPrefix(r.URL.Path, "/docs/") {
			r.URL.Path = strings.TrimPrefix(r.URL.Path, "/docs")
			docsServer.ServeHTTP(w, r)
		} else {
			w.WriteHeader(404)
		}
	}))

	server := &http.Server{
		Addr:              a.serviceProvider.SwaggerServerConfig().Addr(),
		Handler:           mux,
		ReadHeaderTimeout: time.Second * 2,
	}

	a.swaggerServer = server
	return nil
}

func (a *App) runGrpcServer() error {
	logger.Logger().Info("running grpc server: ", a.serviceProvider.GrpcServerConfig().Addr())

	ln, err := net.Listen("tcp", a.serviceProvider.GrpcServerConfig().Addr())
	if err != nil {
		return fmt.Errorf("listen tcp: %w", err)
	}

	err = a.grpcServer.Serve(ln)
	if err != nil {
		return fmt.Errorf("serve grpc server: %w", err)
	}

	return nil
}

func (a *App) runGatewayServer() error {
	logger.Logger().Info("running grpc gateway server: ", a.serviceProvider.GatewayServerConfig().Addr())

	err := a.gatewayServer.ListenAndServeTLS(serverCertFilePath, serverKeyFilePath)
	if err != nil {
		return fmt.Errorf("listen and serve: %w", err)
	}

	return nil
}

func (a *App) runSwaggerServer() error {
	logger.Logger().Info("running swagger server: ", a.serviceProvider.SwaggerServerConfig().Addr())

	err := a.swaggerServer.ListenAndServe()
	if err != nil {
		return fmt.Errorf("listen and serve: %w", err)
	}

	return nil
}

func injectHost(swaggerBytes []byte, host string) ([]byte, error) {
	parsedSwagger := map[string]interface{}{}
	err := json.Unmarshal(swaggerBytes, &parsedSwagger)
	if err != nil {
		return nil, fmt.Errorf("unmarshal error: %w", err)
	}
	parsedSwagger["host"] = host

	resultBytes, err := json.Marshal(parsedSwagger)
	if err != nil {
		return nil, fmt.Errorf("unexpected json marshal error: %w", err)
	}

	return resultBytes, nil
}
