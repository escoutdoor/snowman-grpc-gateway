package v1

import (
	"sync"

	"github.com/escoutdoor/snowman-grpc-gateway/internal/model"
	pb "github.com/escoutdoor/snowman-grpc-gateway/pkg/snowman/v1"
)

type Implementation struct {
	pb.UnimplementedSnowmanServiceV1Server

	snowmen map[string]*model.Snowman
	mu      sync.RWMutex
}

func NewImplementation() *Implementation {
	return &Implementation{
		snowmen: make(map[string]*model.Snowman),
	}
}
