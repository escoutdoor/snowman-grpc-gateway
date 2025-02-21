package interceptor

import (
	"context"

	"github.com/escoutdoor/snowman-grpc-gateway/internal/logger"
	"google.golang.org/grpc"
)

func LoggerUnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {
		logCtx := logger.ToContext(ctx,
			logger.FromContext(ctx).With(
				"operation", info.FullMethod,
				"component", "interceptor",
			),
		)

		logger.Debug(logCtx, "receive request")
		resp, err := handler(ctx, req)
		logger.Debug(logCtx, "handle request")

		if err != nil {
			logger.Error(logCtx, "handle error", "error", err)
		}

		return resp, err
	}
}
