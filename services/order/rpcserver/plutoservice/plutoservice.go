package plutoservice

import (
	"amusingx.fit/amusingx/protos/pluto/service"
	"amusingx.fit/amusingx/services/order/conf"
	"context"
)

type PlutoService struct {
	plutoservice.UnimplementedAmusingxPlutoServiceServer
}

func (s *PlutoService) Pong(context.Context, *plutoservice.BlankParams) (*plutoservice.PongResponse, error) {
	return &plutoservice.PongResponse{ServerName: conf.Conf.Server.Name}, nil
}
