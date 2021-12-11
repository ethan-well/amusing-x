package oauth

import (
	github2 "amusingx.fit/amusingx/apistruct/github"
	"amusingx.fit/amusingx/apistruct/wechat"
	"amusingx.fit/amusingx/services/ganymede/oauth/github"
	"amusingx.fit/amusingx/services/ganymede/oauth/oauthstruct"
	wechat2 "amusingx.fit/amusingx/services/ganymede/oauth/wechat"
	"github.com/ItsWewin/superfactory/xerror"
	"strings"
)

type OAuth interface {
	GetAccessToken(accessTokenUrl, code string) (*oauthstruct.AccessToken, *xerror.Error)
	GetUserProfile(url string, accessToken string) (*oauthstruct.UserProfile, *xerror.Error)
}

func NewOAuth(provider, clientID, clientSecret, redirectUrl, grantType string) (OAuth, *xerror.Error) {
	switch strings.ToLower(provider) {
	case strings.ToLower(github2.ProviderGitHub):
		return github.New(clientID, clientSecret, redirectUrl), nil
	case strings.ToLower(wechat.ProviderWeChat):
		return wechat2.New(clientID, clientSecret, grantType), nil
	default:
		return nil, xerror.NewErrorf(nil, xerror.Code.CParamsError, "unknown oauth provider")
	}
}
