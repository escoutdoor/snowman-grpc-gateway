package v1

import (
	"context"

	"github.com/escoutdoor/snowman-grpc-gateway/internal/converter"
	"github.com/escoutdoor/snowman-grpc-gateway/internal/model"
	pb "github.com/escoutdoor/snowman-grpc-gateway/pkg/snowman/v1"
)

func (i *Implementation) List(ctx context.Context, req *pb.ListSnowmenRequest) (*pb.ListSnowmenResponse, error) {
	i.mu.RLock()
	defer i.mu.RUnlock()

	snowmen := make([]*model.Snowman, 0, len(i.snowmen))
	for _, s := range i.snowmen {
		snowmen = append(snowmen, s)
	}

	return &pb.ListSnowmenResponse{Snowmen: converter.ToPbFromSnowmen(snowmen)}, nil
}
