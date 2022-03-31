package panguserver

import (
	"amusingx.fit/amusingx/protos/comm/response"
	panguservice "amusingx.fit/amusingx/protos/pangu/service/pangu/proto"
	"amusingx.fit/amusingx/services/comm"
	"amusingx.fit/amusingx/services/pangu/rpcserver/handler/category"
	"context"
	"github.com/ItsWewin/superfactory/logger"
)

func (s *PanguServer) CategoryCreate(ctx context.Context, in *panguservice.CategoryCreateRequest) (*panguservice.CategoryCreateResponse, error) {
	return category.HandlerCreateCategoryCreate(ctx, in)
}

func (s *PanguServer) CategoryList(ctx context.Context, in *panguservice.CategoryListRequest) (*panguservice.CategoryListResponse, error) {
	resp, err := category.HandlerCategoryList(ctx, in)
	if err != nil {
		logger.Errorf("CategoryList err: %s", err)
	}

	return resp, err
}

func (s *PanguServer) CategoriesDelete(ctx context.Context, in *panguservice.CategoriesDeleteRequest) (*response.CommResponse, error) {
	response, err := category.HandlerCategoriesDelete(ctx, in)
	return comm.Response(response, err), nil
}

func (s *PanguServer) CategoryDelete(ctx context.Context, in *panguservice.CategoryDeleteRequest) (*response.CommResponse, error) {
	response, err := category.HandlerCategoryDelete(ctx, in)
	return comm.Response(response, err), nil
}

func (s *PanguServer) Category(ctx context.Context, in *panguservice.CategoryRequest) (*panguservice.CategoryResponse, error) {
	return category.HandlerCategory(ctx, in)
}

func (s *PanguServer) CategoryUpdate(ctx context.Context, in *panguservice.CategoryUpdateRequest) (*panguservice.CategoryUpdateResponse, error) {
	return category.HandlerCategoryUpdate(ctx, in)
}
