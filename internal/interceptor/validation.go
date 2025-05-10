package interceptor

import (
	"context"
	"fmt"
	"log"

	"github.com/bufbuild/protovalidate-go"
	"github.com/escoutdoor/snowman-grpc-gateway/internal/grpcutil"

	"github.com/escoutdoor/snowman-grpc-gateway/pkg/tracing"
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
		log.Println("called validation")
		if !ok {
			err := fmt.Errorf("unsupported message type: %T", msg)
			tracing.RecordError(ctx, err)
			return nil, status.Errorf(codes.Internal, err.Error())
		}

		err := validator.Validate(msg)
		if err != nil {
			tracing.RecordError(ctx, err)
			return nil, grpcutil.ProtoValidationError(err)
		}

		return handler(ctx, req)
	}
}
