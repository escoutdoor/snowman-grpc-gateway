package converter

import (
	"github.com/escoutdoor/snowman-grpc-gateway/internal/model"
	pb "github.com/escoutdoor/snowman-grpc-gateway/pkg/snowman/v1"
)

func ToPbFromSnowman(in *model.Snowman) *pb.Snowman {
	return &pb.Snowman{
		Id:     in.ID,
		Name:   in.Name,
		Height: in.Height,
		Width:  in.Width,
	}
}

func ToPbFromSnowmen(in []*model.Snowman) []*pb.Snowman {
	snowmen := make([]*pb.Snowman, len(in))
	for i, sm := range in {
		snowmen[i] = ToPbFromSnowman(sm)
	}

	return snowmen
}
