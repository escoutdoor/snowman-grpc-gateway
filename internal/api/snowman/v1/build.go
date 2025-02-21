package v1

import (
	"context"

	"github.com/escoutdoor/snowman-grpc-gateway/internal/model"
	pb "github.com/escoutdoor/snowman-grpc-gateway/pkg/snowman/v1"
	"github.com/google/uuid"
)

func (i *Implementation) Build(ctx context.Context, req *pb.BuildSnowmanRequest) (*pb.BuildSnowmanResponse, error) {
	snowman := &model.Snowman{
		ID:     uuid.New().String(),
		Name:   req.GetName(),
		Height: req.GetHeight(),
		Width:  req.GetWidth(),
	}

	i.mu.Lock()
	i.snowmen[snowman.ID] = snowman
	i.mu.Unlock()

	return &pb.BuildSnowmanResponse{Id: snowman.ID}, nil
}
