package charonserver

import (
	"amusingx.fit/amusingx/protos/amusingcharonservice/charonservice"
	"amusingx.fit/amusingx/services/amusingcharonserv/conf"
	"context"
)

type CharonServer struct {
	charonservice.UnimplementedCharonServServer
}

// 	Pong(context.Context, *BlankParams) (*PongResponse, error)
func (s *CharonServer) Pong(ctx context.Context, in *charonservice.BlankParams) (*charonservice.PongResponse, error) {
	return &charonservice.PongResponse{ServerName: conf.Conf.ServerName}, nil
}
