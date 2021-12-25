package panguserver

import (
	"amusingx.fit/amusingx/protos/pangu/service"
	"amusingx.fit/amusingx/services/pangu/conf"
	"amusingx.fit/amusingx/services/pangu/rpcserver/handler/category"
	"context"
	"github.com/ItsWewin/superfactory/logger"
)

type PanguServer struct {
	panguservice.UnimplementedPanGuServiceServer
}

// 	Pong(context.Context, *BlankParams) (*PongResponse, error)
func (s *PanguServer) Pong(ctx context.Context, in *panguservice.BlankParams) (*panguservice.PongResponse, error) {
	return &panguservice.PongResponse{ServerName: conf.Conf.Server.Name}, nil
}

func (s *PanguServer) CategoryCreate(ctx context.Context, in *panguservice.CategoryCreateRequest) (*panguservice.CategoryCreateResponse, error) {
	logger.Errorf("info: %s", logger.ToJson(in))
	return category.HandlerCreateCategoryCreate(ctx, in)
}
