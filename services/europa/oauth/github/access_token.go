package github

import (
	"amusingx.fit/amusingx/services/europa/oauth"
	"github.com/ItsWewin/superfactory/httputil"
	"github.com/ItsWewin/superfactory/httputil/rest"
	"github.com/ItsWewin/superfactory/xerror"
)

type OAuth struct {
	AccessTokenUrl string `json:"access_token_url"`
	ClientID       string `json:"client_id"`
	ClientSecret   string `json:"client_secret"`
	Code           string `json:"code"`
	RedirectUrl    string `json:"redirect_url"`
}

type AccessTokenRequest struct {
	AccessTokenUrl string `json:"access_token_url"`
	ClientID       string `json:"client_id"`
	ClientSecret   string `json:"client_secret"`
	Code           string `json:"code"`
	RedirectUrl    string `json:"redirect_url"`
}

func New(clientID, clientSecret, redirectUrl string) *OAuth {
	return &OAuth{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		RedirectUrl:  redirectUrl,
	}
}

func (c *OAuth) GetAccessToken(accessTokenUrl, code string) (*oauth.AccessTokenResponse, *xerror.Error) {
	req := AccessTokenRequest{
		AccessTokenUrl: c.AccessTokenUrl,
		ClientID:       c.ClientID,
		ClientSecret:   c.ClientSecret,
		Code:           code,
		RedirectUrl:    c.RedirectUrl,
	}

	opts := func(opts *rest.Options) {
		opts.Header = map[string]string{
			httputil.HeaderContent: httputil.JsonHeaderContent,
			httputil.HeaderAccept:  httputil.JsonHeaderAccept,
		}
	}

	var resp oauth.AccessTokenResponse
	err := rest.Post(accessTokenUrl, req, &resp, opts)
	if err != nil {
		return nil, xerror.NewErrorf(err, err.Code, err.Message)
	}

	return &resp, nil
}
