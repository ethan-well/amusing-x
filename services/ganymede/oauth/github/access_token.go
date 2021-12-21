package github

import (
	"amusingx.fit/amusingx/apistruct/github"
	"amusingx.fit/amusingx/services/ganymede/oauth/oauthstruct"
	"github.com/ItsWewin/superfactory/aerror"
	"github.com/ItsWewin/superfactory/httputil"
	"github.com/ItsWewin/superfactory/httputil/rest"
)

type OAuth struct {
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	Code         string `json:"code"`
	RedirectUrl  string `json:"redirect_url"`
}

func New(clientID, clientSecret, redirectUrl string) *OAuth {
	return &OAuth{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		RedirectUrl:  redirectUrl,
	}
}

func (c *OAuth) GetAccessToken(accessTokenUrl, code string) (*oauthstruct.AccessToken, aerror.Error) {
	req := github.AccessTokenRequest{
		ClientID:     c.ClientID,
		ClientSecret: c.ClientSecret,
		Code:         code,
		RedirectUrl:  c.RedirectUrl,
	}

	opts := func(opts *rest.Options) {
		opts.Header = map[string]string{
			httputil.HeaderContent: httputil.JsonHeaderContent,
			httputil.HeaderAccept:  httputil.JsonHeaderAccept,
		}
	}

	var dest github.AccessTokenResponse
	_, err := rest.Post(accessTokenUrl, req, &dest, opts)
	if err != nil {
		return nil, aerror.NewErrorf(err, err.Code(), err.Message())
	}

	if dest.IsError() {
		return nil, aerror.NewErrorf(nil, aerror.Code.OtherNetworkError, dest.ErrorDescription)
	}

	return &oauthstruct.AccessToken{
		AccessToken:  dest.AccessToken,
		Scope:        dest.Scope,
		TokenType:    dest.TokenType,
		ExpiresIn:    oauthstruct.AccessTokenExpireNever,
		RefreshToken: "",
		OpenID:       "",
		UnionID:      "",
	}, nil
}
