package oauthlogin

import (
	ganymedeservice "amusingx.fit/amusingx/protos/ganymede/service/ganymede/proto"
	"amusingx.fit/amusingx/services/ganymede/oauth"
	"context"
	"github.com/ItsWewin/superfactory/aerror"
)

func HandlerGetOAuthInfo(ctx context.Context, in *ganymedeservice.OAuthInfoRequest) (*ganymedeservice.OAuthInfoResponse, error) {
	if in == nil {
		return nil, aerror.NewError(nil, aerror.Code.CParamsError, "requst info is invalid")
	}

	if err := in.Valid(); err != nil {
		return nil, err
	}

	sate, err := oauth.RandomState(ctx)
	if err != nil {
		return nil, err
	}

	provider, err := GetProvider(ctx, in.Provider, in.Service)
	if err != nil {
		return nil, err
	}

	return &ganymedeservice.OAuthInfoResponse{
		Provider:    provider.Provider,
		OauthUrl:    provider.OAuthUrl,
		ClientId:    provider.ClientID,
		Scope:       provider.Scope,
		State:       sate,
		GrantType:   provider.GrantType,
		RedirectUrl: provider.RedirectUrl,
	}, nil
}
