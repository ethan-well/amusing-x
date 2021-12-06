package oauth

import (
	"amusingx.fit/amusingx/services/europa/oauth/github"
	"github.com/ItsWewin/superfactory/xerror"
)

const (
	ProviderGitHub = "GitHub"
)

type AccessTokenResponse struct {
	AccessToken string `json:"access_token"`
	Scope       string `json:"scope"`
	TokenType   string `json:"token_type"`
}

type UserProfile struct {
	Login     string `json:"login"`
	ID        int64  `json:"id"`
	AvatarUrl string `json:"avatar_url"`
	Name      string `json:"name"`
	Company   string `json:"company"`
	Blog      string `json:"blog"`
	Email     string `json:"email"`
	Location  string `json:"location"`
}

type OAuth interface {
	GetAccessToken(accessTokenUrl, code string) (*AccessTokenResponse, *xerror.Error)
	GetUserProfile(url string, accessToken string) (*UserProfile, *xerror.Error)
}

func NewOAuth(provider, clientID, clientSecret, redirectUrl string) (OAuth, *xerror.Error) {
	switch provider {
	case ProviderGitHub:
		return github.New(clientID, clientSecret, redirectUrl), nil
	default:
		return nil, xerror.NewErrorf(nil, xerror.Code.CParamsError, "unknown oauth provider")
	}
}
