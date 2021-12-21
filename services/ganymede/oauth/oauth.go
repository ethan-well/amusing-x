package oauth

import (
	"amusingx.fit/amusingx/services/ganymede/oauth/github"
	"amusingx.fit/amusingx/services/ganymede/oauth/oauthstruct"
	wechat2 "amusingx.fit/amusingx/services/ganymede/oauth/wechat"
	"context"
	"github.com/ItsWewin/superfactory/aerror"
	"strings"
)

type OAuth interface {
	GetAccessToken(accessTokenUrl, code string) (*oauthstruct.AccessToken, aerror.Error)
	GetUserProfile(url string, accessToken string) (*oauthstruct.UserProfile, aerror.Error)
}

func NewOAuth(provider, clientID, clientSecret, redirectUrl, grantType string) (OAuth, aerror.Error) {
	switch strings.ToLower(provider) {
	case strings.ToLower(oauthstruct.ProviderGitHub):
		return github.New(clientID, clientSecret, redirectUrl), nil
	case strings.ToLower(oauthstruct.ProviderWeChat):
		return wechat2.New(clientID, clientSecret, grantType), nil
	default:
		return nil, aerror.NewErrorf(nil, aerror.Code.CParamsError, "unknown oauth provider")
	}
}

// oauth info 中的 state 信息存储
func RandomState(ctx context.Context) (string, aerror.Error) {
	return StoreIns.RandomStateCode(ctx)
}

// 登陆验证 state 是否有效，防止 csrf
func StateCodeValid(ctx context.Context, stateCode string) aerror.Error {
	return StoreIns.ValidSate(ctx, stateCode)
}
