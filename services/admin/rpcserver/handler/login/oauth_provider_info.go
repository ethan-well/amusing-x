package login

import (
	panguservice "amusingx.fit/amusingx/protos/pangu/service/pangu/proto"
	ganymedeservice "amusingx.fit/amusingx/protos/user/service/user/proto"
	"amusingx.fit/amusingx/rpcclient/user"
	"context"
	"github.com/ItsWewin/superfactory/aerror"
	"net/url"
)

func HandlerOauthProviderInfo(ctx context.Context, req *panguservice.OauthProviderInfoRequest) (*panguservice.OAuthProviderInfoResponse, aerror.Error) {
	rpcReq := &ganymedeservice.OAuthInfoRequest{
		Provider: req.Provider,
		Service:  req.Service,
	}

	resp, err := user.Client.OAuthInfo(ctx, rpcReq)
	if err != nil || resp == nil {
		return nil, aerror.NewErrorf(err, aerror.Code.BUnexpectedData, "login failed")
	}

	u, err := url.Parse(resp.OauthUrl)
	if err != nil {
		return nil, aerror.NewError(nil, aerror.Code.BUnexpectedData, err.Error())
	}

	_, err = url.Parse(resp.RedirectUrl)
	if err != nil {
		return nil, aerror.NewError(nil, aerror.Code.BUnexpectedData, err.Error())
	}

	q := u.Query()
	q.Set("client_id", resp.ClientId)
	q.Set("redirect_uri", resp.RedirectUrl)
	q.Set("scope", resp.Scope)
	q.Set("state", resp.State)
	q.Set("grant_type", resp.GrantType)
	u.RawQuery = q.Encode()

	providerInfo := &panguservice.OAuthProviderInfo{
		Provider:     resp.Provider,
		OauthUrl:     resp.OauthUrl,
		ClientID:     resp.ClientId,
		Scope:        resp.Scope,
		State:        resp.State,
		GrantType:    resp.GrantType,
		RedirectUrl:  resp.RedirectUrl,
		CompletePath: u.String(),
	}

	return &panguservice.OAuthProviderInfoResponse{
		Succeed: true,
		Result:  providerInfo,
	}, nil
}
