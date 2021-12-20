package charonserver

import (
	"amusingx.fit/amusingx/protos/charons/service"
	"amusingx.fit/amusingx/services/charon/conf"
	"amusingx.fit/amusingx/services/charon/rpcserver/controller/category"
	"context"
)

type CharonServer struct {
	charonservice.UnimplementedCharonServServer
}

// 	Pong(context.Context, charonservice. .BlankParams) (*PongResponse, error)
func (s *CharonServer) Pong(ctx context.Context, in *charonservice.BlankParams) (*charonservice.PongResponse, error) {
	return &charonservice.PongResponse{ServerName: conf.Conf.Server.Name}, nil
}

func (s *CharonServer) Create(ctx context.Context, in *charonservice.CategoryCreateRequest) (*charonservice.CategoryCreateResponse, error) {
	if err := in.Valid(); err != nil {
		return nil, err
	}

	return category.HandlerCreateCategory(ctx, in)
}
