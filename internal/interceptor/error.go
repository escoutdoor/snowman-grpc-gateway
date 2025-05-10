package interceptor

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func ErrorUnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		resp, err = handler(ctx, req)

		_, ok := status.FromError(err)
		if ok {
			return resp, err
		}

		switch {
		default:
			err = status.Error(codes.Internal, err.Error())
		}

		return resp, err
	}
}
