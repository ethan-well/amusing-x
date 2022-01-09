package panguserver

import (
	panguservice "amusingx.fit/amusingx/protos/pangu/service/pangu/proto"
	"amusingx.fit/amusingx/services/pangu/conf"
	"amusingx.fit/amusingx/services/pangu/rpcserver/getway"
	"amusingx.fit/amusingx/services/pangu/rpcserver/handler/category"
	"amusingx.fit/amusingx/services/pangu/rpcserver/handler/login"
	"context"
	"github.com/ItsWewin/superfactory/aerror"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
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

func (s *PanguServer) OauthLogin(ctx context.Context, in *panguservice.OAuthLoginRequest) (*panguservice.OAuthLoginResponse, error) {
	login, err := login.HandlerOauthLogin(ctx, in)
	if err != nil || login == nil || login.Result.UserInfo == nil || login.Result.SessionInfo == nil {
		return nil, aerror.NewError(err, aerror.Code.BUnexpectedData, "login failed")
	}

	header := metadata.Pairs(
		getway.SessionIdKey, login.Result.SessionInfo.SessionID,
	)
	grpc.SetHeader(ctx, header)

	return login, nil
}

func (s *PanguServer) OauthProviderInfo(ctx context.Context, in *panguservice.OauthProviderInfoRequest) (*panguservice.OAuthProviderInfoResponse, error) {
	return login.HandlerOauthProviderInfo(ctx, in)
}

func (s *PanguServer) Logout(ctx context.Context, in *panguservice.LogoutRequest) (*panguservice.LogoutResponse, error) {
	result, err := login.HandlerLogout(ctx, in)
	if err != nil {
		return nil, err
	}

	if result.Logout {
		header := metadata.Pairs(
			getway.DeleteSessionKey, "true",
		)
		grpc.SetHeader(ctx, header)
	}

	return result, nil
}
