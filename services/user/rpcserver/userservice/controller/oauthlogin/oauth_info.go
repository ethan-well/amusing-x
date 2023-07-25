package oauthlogin

import (
	service "amusingx.fit/amusingx/protos/user/service/user/proto"
	"amusingx.fit/amusingx/services/user/oauth"
	"context"
	"github.com/ItsWewin/superfactory/aerror"
)

func HandlerGetOAuthInfo(ctx context.Context, in *service.OAuthInfoRequest) (*service.OAuthInfoResponse, aerror.Error) {

	if in == nil {
		return nil, aerror.NewError(nil, aerror.Code.CParamsErr, "requst info is invalid")
	}

	if err := in.Valid(); err != nil {
		return nil, err
	}

	sate, err := oauth.RandomState(ctx)
	if err != nil {
		return nil, err
	}

	provider, err2 := GetProvider(ctx, in.Provider, in.Service)
	if err2 != nil {
		return nil, err2
	}

	return &service.OAuthInfoResponse{
		Provider:    provider.Provider,
		OauthUrl:    provider.OAuthUrl,
		ClientId:    provider.ClientID,
		Scope:       provider.Scope,
		State:       sate,
		GrantType:   provider.GrantType,
		RedirectUrl: provider.RedirectUrl,
	}, nil
}
