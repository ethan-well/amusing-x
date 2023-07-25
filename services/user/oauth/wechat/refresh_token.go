package wechat

import (
	"amusingx.fit/amusingx/apistruct/wechat"
	"amusingx.fit/amusingx/services/user/oauth/oauthstruct"
	"github.com/ItsWewin/superfactory/aerror"
	"github.com/ItsWewin/superfactory/httputil"
	"github.com/ItsWewin/superfactory/httputil/rest"
	"net/url"
)

//AccessTokenResponse

func (o *OAuth) RefreshToken(refreshTokenUrl, refreshToken string) (*oauthstruct.AccessToken, aerror.Error) {
	var params = url.Values{}
	params.Set("appid", o.ClientID)
	params.Set("grant_type", o.GrantType)
	params.Set("refresh_token", refreshToken)

	u, err := url.Parse(refreshTokenUrl)
	if err != nil {
		return nil, aerror.NewErrorf(nil, aerror.Code.BUnexpectedData, "accessTokenUrl format err")
	}
	u.RawQuery = params.Encode()

	opts := func(opts *rest.Options) {
		opts.Header = map[string]string{
			httputil.HeaderContent: httputil.JsonHeaderContent,
			httputil.HeaderAccept:  httputil.JsonHeaderAccept,
		}
	}

	var dest wechat.AccessTokenResponse
	xErr := rest.Get(u.String(), &dest, opts)
	if xErr != nil {
		return nil, xErr
	}

	if dest.IsError() {
		return nil, aerror.NewErrorf(nil, aerror.Code.OtherNetworkErr, dest.Errmsg)
	}

	return &oauthstruct.AccessToken{
		AccessToken:  dest.AccessToken,
		Scope:        dest.Scope,
		ExpiresIn:    dest.ExpiresIn,
		RefreshToken: dest.RefreshToken,
		OpenID:       dest.Openid,
		UnionID:      dest.Unionid,
	}, nil
}
