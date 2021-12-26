package charonserver

import (
	charonservice "amusingx.fit/amusingx/protos/charon/service/charon/proto"
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

func (s *CharonServer) Categories(ctx context.Context, in *charonservice.CategoryListRequest) (*charonservice.CategoryListResponse, error) {
	return category.HandlerCategoryList(ctx, in)
}

func (s *CharonServer) Delete(ctx context.Context, in *charonservice.CategoryDeleteRequest) (*charonservice.CategoryDeleteResponse, error) {
	return category.HandlerCategoryDelete(ctx, in)
}

func (s *CharonServer) Category(ctx context.Context, in *charonservice.CategoryRequest) (*charonservice.CategoryResponse, error) {
	return category.HandlerCategory(ctx, in)
}

func (s *CharonServer) Update(ctx context.Context, in *charonservice.CategoryUpdateRequest) (*charonservice.CategoryUpdateResponse, error) {
	return category.HandlerUpdate(ctx, in)
}
