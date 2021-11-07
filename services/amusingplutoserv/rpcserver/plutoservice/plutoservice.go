package plutoservice

import (
	"amusingx.fit/amusingx/protos/amusingplutoservice/plutoservice"
	"amusingx.fit/amusingx/services/amusingplutoserv/conf"
	"context"
)

type PlutoService struct {
	plutoservice.UnimplementedAmusingxPlutoServiceServer
}

func (s *PlutoService) Pong(context.Context, *plutoservice.BlankParams) (*plutoservice.PongResponse, error) {
	return &plutoservice.PongResponse{ServerName: conf.Conf.ServerName}, nil
}
