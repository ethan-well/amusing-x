package oauth

import (
	github2 "amusingx.fit/amusingx/apistruct/github"
	"amusingx.fit/amusingx/services/ganymede/oauth/github"
	"github.com/ItsWewin/superfactory/xerror"
)

type OAuth interface {
	GetAccessToken(accessTokenUrl, code string) (*github2.AccessTokenResponse, *xerror.Error)
	GetUserProfile(url string, accessToken string) (*github2.UserProfile, *xerror.Error)
}

func NewOAuth(provider, clientID, clientSecret, redirectUrl string) (OAuth, *xerror.Error) {
	switch provider {
	case github2.ProviderGitHub:
		return github.New(clientID, clientSecret, redirectUrl), nil
	default:
		return nil, xerror.NewErrorf(nil, xerror.Code.CParamsError, "unknown oauth provider")
	}
}
