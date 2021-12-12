package oauth

import (
	"amusingx.fit/amusingx/services/ganymede/oauth/github"
	"amusingx.fit/amusingx/services/ganymede/oauth/oauthstruct"
	wechat2 "amusingx.fit/amusingx/services/ganymede/oauth/wechat"
	"context"
	"github.com/ItsWewin/superfactory/xerror"
	"strings"
)

type OAuth interface {
	GetAccessToken(accessTokenUrl, code string) (*oauthstruct.AccessToken, *xerror.Error)
	GetUserProfile(url string, accessToken string) (*oauthstruct.UserProfile, *xerror.Error)
}

func NewOAuth(provider, clientID, clientSecret, redirectUrl, grantType string) (OAuth, *xerror.Error) {
	switch strings.ToLower(provider) {
	case strings.ToLower(oauthstruct.ProviderGitHub):
		return github.New(clientID, clientSecret, redirectUrl), nil
	case strings.ToLower(oauthstruct.ProviderWeChat):
		return wechat2.New(clientID, clientSecret, grantType), nil
	default:
		return nil, xerror.NewErrorf(nil, xerror.Code.CParamsError, "unknown oauth provider")
	}
}

func RandomState(ctx context.Context) (string, *xerror.Error) {
	return StoreIns.RandomStateCode(ctx)
}

func StateCodeValid(ctx context.Context, stateCode string) *xerror.Error {
	return StoreIns.ValidSate(ctx, stateCode)
}
