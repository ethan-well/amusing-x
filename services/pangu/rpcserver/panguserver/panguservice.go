package panguserver

import (
	panguservice "amusingx.fit/amusingx/protos/pangu/service/pangu"
	"amusingx.fit/amusingx/services/pangu/conf"
	"amusingx.fit/amusingx/services/pangu/rpcserver/handler/category"
	"amusingx.fit/amusingx/services/pangu/rpcserver/handler/login"
	"context"
)

type PanguServer struct {
	panguservice.UnimplementedPanGuServiceServer
}

// 	Pong(context.Context, *BlankParams) (*PongResponse, error)
func (s *PanguServer) Pong(ctx context.Context, in *panguservice.BlankParams) (*panguservice.PongResponse, error) {
	return &panguservice.PongResponse{ServerName: conf.Conf.Server.Name}, nil
}

func (s *PanguServer) CategoryCreate(ctx context.Context, in *panguservice.CategoryCreateRequest) (*panguservice.CategoryCreateResponse, error) {
	return category.HandlerCreateCategoryCreate(ctx, in)
}

func (s *PanguServer) CategoryList(ctx context.Context, in *panguservice.CategoryListRequest) (*panguservice.CategoryListResponse, error) {
	return category.HandlerCategoryList(ctx, in)
}

func (s *PanguServer) CategoryDelete(ctx context.Context, in *panguservice.CategoryDeleteRequest) (*panguservice.CategoryDeleteResponse, error) {
	return category.HandlerCategoryDelete(ctx, in)
}

func (s *PanguServer) Category(ctx context.Context, in *panguservice.CategoryRequest) (*panguservice.CategoryResponse, error) {
	return category.HandlerCategory(ctx, in)
}

func (s *PanguServer) CategoryUpdate(ctx context.Context, in *panguservice.CategoryUpdateRequest) (*panguservice.CategoryUpdateResponse, error) {
	return category.HandlerCategoryUpdate(ctx, in)
}

func (s *PanguServer) Login(ctx context.Context, in *panguservice.OAuthLoginRequest) (*panguservice.OAuthLoginResponse, error) {
	return login.HandlerLogin(ctx, in)
}
