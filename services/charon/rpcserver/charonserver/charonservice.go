package charonserver

import (
	"amusingx.fit/amusingx/protos/charons/service"
	"amusingx.fit/amusingx/services/charon/conf"
	"context"
)

type CharonServer struct {
	service.UnimplementedCharonServServer
}

// 	Pong(context.Context, *BlankParams) (*PongResponse, error)
func (s *CharonServer) Pong(ctx context.Context, in *service.BlankParams) (*service.PongResponse, error) {
	return &service.PongResponse{ServerName: conf.Conf.Server.Name}, nil
}
