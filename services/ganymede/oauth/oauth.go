package oauth

import (
	github2 "amusingx.fit/amusingx/apistruct/github"
	"amusingx.fit/amusingx/services/ganymede/oauth/github"
	"github.com/ItsWewin/superfactory/xerror"
	"strings"
)

const (
	AccessTokenExpireNever = -1
)

type AccessToken struct {
	AccessToken  string
	Scope        string
	TokenType    string
	ExpiresIn    int64 // -1 表示永不过期
	RefreshToken string
	OpenID       string
	UnionID      string
}

type UserProfile struct {
	Login       string // 地方方网站唯一标识
	OuterUserID int64  // 用户外部 id
	AvatarUrl   string
	Name        string
	Company     string
	Blog        string
	Email       string
	Location    string
}

type OAuth interface {
	GetAccessToken(accessTokenUrl, code string) (*AccessToken, *xerror.Error)
	GetUserProfile(url string, accessToken string) (*UserProfile, *xerror.Error)
}

func NewOAuth(provider, clientID, clientSecret, redirectUrl string) (OAuth, *xerror.Error) {
	switch strings.ToLower(provider) {
	case strings.ToLower(github2.ProviderGitHub):
		return github.New(clientID, clientSecret, redirectUrl), nil
	default:
		return nil, xerror.NewErrorf(nil, xerror.Code.CParamsError, "unknown oauth provider")
	}
}
