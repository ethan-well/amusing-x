package oauthlogin

import (
	ganymedeservice "amusingx.fit/amusingx/protos/ganymede/service"
	"amusingx.fit/amusingx/services/ganymede/conf"
	"amusingx.fit/amusingx/services/ganymede/oauth"
	"amusingx.fit/amusingx/services/ganymede/oauth/oauthstruct"
	"context"
	"github.com/ItsWewin/superfactory/xerror"
	"google.golang.org/grpc"
)

func HandlerGetOAuthInfo(ctx context.Context, in *ganymedeservice.OAuthInfoRequest, opts ...grpc.CallOption) (*ganymedeservice.OAuthInfoResponse, error) {
	if in == nil {
		return nil, xerror.NewError(nil, xerror.Code.CParamsError, "requst info is invalid")
	}

	if err := in.Valid(); err != nil {
		return nil, err
	}

	sate, err := oauth.RandomState(ctx)
	if err != nil {
		return nil, err
	}

	switch in.Provider {
	case oauthstruct.ProviderGitHub:
		provider := conf.Conf.OAuth.Github
		return &ganymedeservice.OAuthInfoResponse{
			Provider:    provider.Provider,
			OauthUrl:    provider.OAuthUrl,
			ClientId:    provider.ClientID,
			Scope:       provider.Scope,
			State:       sate,
			GrantType:   provider.GrantType,
			RedirectUrl: provider.RedirectUrl,
		}, nil
	case oauthstruct.ProviderWeChat:
		provider := conf.Conf.OAuth.WeChat
		return &ganymedeservice.OAuthInfoResponse{
			Provider:    provider.Provider,
			OauthUrl:    provider.OAuthUrl,
			ClientId:    provider.ClientID,
			Scope:       provider.Scope,
			State:       sate,
			GrantType:   provider.GrantType,
			RedirectUrl: provider.RedirectUrl,
		}, nil
	default:
		return nil, xerror.NewErrorf(err, xerror.Code.CParamsError, "provider '%s' is invalid", in.Provider)
	}
}
