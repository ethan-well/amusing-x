package charonserver

import (
	charonservice "amusingx.fit/amusingx/protos/charon/service/charon/proto"
	"amusingx.fit/amusingx/services/product/conf"
	"context"
)

type CharonServer struct {
	charonservice.UnimplementedCharonServServer
}

// 	Pong(context.Context, charonservice. .BlankParams) (*PongResponse, error)
func (s *CharonServer) Pong(ctx context.Context, in *charonservice.BlankParams) (*charonservice.PongResponse, error) {
	return &charonservice.PongResponse{ServerName: conf.Conf.Server.Name}, nil
}
