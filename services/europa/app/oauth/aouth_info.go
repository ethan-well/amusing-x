package oauth

import (
	"amusingx.fit/amusingx/apistruct/europa"
	ganymedeservice "amusingx.fit/amusingx/protos/ganymede/service"
	"amusingx.fit/amusingx/services/europa/rpcclient/ganymede"
	"context"
	"github.com/ItsWewin/superfactory/xerror"
	"net/url"
)

func GetOAuthInfo(ctx context.Context, provider string) (*europa.OAuthInfo, *xerror.Error) {
	in := &ganymedeservice.OAuthInfoRequest{Provider: provider}

	resp, err := ganymede.RPCClient.Client.OAuthInfo(ctx, in)
	if err != nil {
		return nil, xerror.NewErrorf(nil, xerror.Code.SUnexpectedErr, err.Error())
	}

	u, err := url.Parse(resp.OauthUrl)
	if err != nil {
		return nil, xerror.NewError(nil, xerror.Code.BUnexpectedData, err.Error())
	}

	_, err = url.Parse(resp.RedirectUrl)
	if err != nil {
		return nil, xerror.NewError(nil, xerror.Code.BUnexpectedData, err.Error())
	}

	q := u.Query()
	q.Set("client_id", resp.ClientId)
	q.Set("redirect_uri", resp.RedirectUrl)
	q.Set("scope", resp.Scope)
	q.Set("state", resp.State)
	q.Set("grant_type", resp.GrantType)
	u.RawQuery = q.Encode()

	return &europa.OAuthInfo{
		Provider:     resp.Provider,
		OauthUrl:     resp.OauthUrl,
		ClientID:     resp.ClientId,
		Scope:        resp.Scope,
		State:        resp.State,
		GrantType:    resp.GrantType,
		RedirectUrl:  resp.RedirectUrl,
		CompletePath: u.String(),
	}, nil
}
