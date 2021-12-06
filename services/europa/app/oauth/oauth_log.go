package oauth

import (
	"amusingx.fit/amusingx/services/europa/conf"
	"amusingx.fit/amusingx/services/europa/oauth"
	"github.com/ItsWewin/superfactory/xerror"
)

type LoginDomain struct {
	Code     string
	Provider string
}

func NewOauthLoginDomain(provider, code string) (*LoginDomain, *xerror.Error) {
	clientID, clientSecret, redirectUrl, accessTokenUrl, userProfileUrl, err := getOauthConf(provider)
	if err != nil {
		return nil, err
	}

	oAuth, err := oauth.NewOAuth(provider, clientID, clientSecret, redirectUrl)
	if err != nil {
		return nil, err
	}

	token, err := oAuth.GetAccessToken(accessTokenUrl, code)
	if err != nil {
		return nil, err
	}

	userProfile, err := oAuth.GetUserProfile(userProfileUrl, token.AccessToken)
	if err != nil {
		return nil, err
	}
}

func getOauthConf(provider string) (clientID, clientSecret, redirectUrl, accessTokenUrl, userProfileUrl string, err *xerror.Error) {
	switch provider {
	case oauth.ProviderGitHub:
		p := conf.Conf.OAuth.Github
		clientID = p.ClientID
		clientSecret = p.ClientSecret
		redirectUrl = p.RedirectUrl
		accessTokenUrl = p.AccessTokenUrl
		userProfileUrl = p.UserProfileUrl

		return
	default:
		err = xerror.NewError(nil, xerror.Code.CParamsError, "provider is invalid")
		return
	}
}
