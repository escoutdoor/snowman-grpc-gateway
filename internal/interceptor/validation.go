package interceptor

import (
	"context"

	"github.com/bufbuild/protovalidate-go"
	"github.com/escoutdoor/snowman-grpc-gateway/internal/grpcutil"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
)

func ValidationUnaryServerInterceptor(validator protovalidate.Validator) grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {
		msg, ok := req.(proto.Message)
		if !ok {
			return nil, status.Errorf(codes.Internal, "unsupported message type: %T", msg)
		}

		err := validator.Validate(msg)
		if err != nil {
			// return nil, status.Error(codes.InvalidArgument, err.Error())
			return nil, grpcutil.ProtoValidationError(err)
		}

		return handler(ctx, req)
	}
}
