package interceptor

import (
	"context"
	"time"

	"github.com/escoutdoor/snowman-grpc-gateway/internal/metrics"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
)

func MetricsUnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		metrics.IncRequestCounter()

		start := time.Now()
		resp, err := handler(ctx, req)
		end := time.Since(start).Seconds()

		code := status.Code(err)
		metrics.IncResponseCounter(code.String(), info.FullMethod)
		metrics.HistogramResponseTimeObserve(code.String(), end)
		return resp, err
	}
}
