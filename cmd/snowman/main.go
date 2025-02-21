package main

import (
	"context"
	"flag"

	"github.com/escoutdoor/snowman-grpc-gateway/internal/app"
	"github.com/escoutdoor/snowman-grpc-gateway/internal/logger"
	"go.uber.org/zap/zapcore"
)

var configPath string

func main() {
	ctx := context.Background()
	flag.StringVar(&configPath, "config_path", ".env", "path to config file")
	flag.Parse()

	logger.SetLevel(zapcore.DebugLevel)

	logger.Info(ctx, "init application")
	a, err := app.New(ctx, configPath)
	if err != nil {
		logger.Fatalf(ctx, "failed to init application: %s", err)
	}

	logger.Info(ctx, "running application..")
	err = a.Run()
	if err != nil {
		logger.Fatalf(ctx, "failed to run application: %s", err)
	}
}
