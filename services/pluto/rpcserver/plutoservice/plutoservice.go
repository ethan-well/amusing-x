package plutoservice

import (
	"amusingx.fit/amusingx/protos/plutoice/plutoservice"
	"amusingx.fit/amusingx/services/pluto/conf"
	"context"
)

type PlutoService struct {
	plutoservice.UnimplementedAmusingxPlutoServiceServer
}

func (s *PlutoService) Pong(context.Context, *plutoservice.BlankParams) (*plutoservice.PongResponse, error) {
	return &plutoservice.PongResponse{ServerName: conf.Conf.ServerName}, nil
}
