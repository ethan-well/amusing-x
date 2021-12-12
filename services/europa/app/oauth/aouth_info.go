package oauth

import (
	"amusingx.fit/amusingx/apistruct/europa"
	ganymedeservice "amusingx.fit/amusingx/protos/ganymede/service"
	"amusingx.fit/amusingx/services/europa/rpcclient/ganymede"
	"context"
	"github.com/ItsWewin/superfactory/xerror"
)

func GetOAuthInfo(ctx context.Context, provider string) (*europa.OAuthInfo, *xerror.Error) {
	in := &ganymedeservice.OAuthInfoRequest{Provider: provider}

	resp, err := ganymede.RPCClient.Client.OAuthInfo(ctx, in)
	if err != nil {
		return nil, xerror.NewErrorf(nil, xerror.Code.SUnexpectedErr, err.Error())
	}

	return &europa.OAuthInfo{
		Provider:    resp.Provider,
		OauthUrl:    resp.OauthUrl,
		ClientID:    resp.ClientId,
		Scope:       resp.Scope,
		State:       resp.State,
		GrantType:   resp.GrantType,
		RedirectUrl: resp.RedirectUrl,
	}, nil
}
