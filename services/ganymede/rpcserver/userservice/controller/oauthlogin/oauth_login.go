package oauthlogin

import (
	"amusingx.fit/amusingx/apistruct/github"
	"amusingx.fit/amusingx/mysqlstruct/ganymede"
	"amusingx.fit/amusingx/protos/ganymede/service"
	"amusingx.fit/amusingx/services/ganymede/conf"
	"amusingx.fit/amusingx/services/ganymede/mysql/ganymededb/model"
	"amusingx.fit/amusingx/services/ganymede/oauth"
	"context"
	"github.com/ItsWewin/superfactory/logger"
	"github.com/ItsWewin/superfactory/xerror"
)

func HandlerOAuthLogin(ctx context.Context, req *ganymedeservice.OAuthLoginRequest) (*ganymedeservice.OAuthLoginResponse, *xerror.Error) {
	err := getAndValidRequest(req)
	if err != nil {
		return nil, err
	}

	logger.Infof("[HandlerOAuthLogin], req: %s", logger.ToJson(req))
	err = oauthLogin(ctx, req)
	if err != nil {
		return nil, err
	}

	return &ganymedeservice.OAuthLoginResponse{Result: true}, nil
}

func getAndValidRequest(req *ganymedeservice.OAuthLoginRequest) *xerror.Error {
	xErr := req.Valid()
	if xErr != nil {
		return xErr
	}

	return nil
}

func oauthLogin(ctx context.Context, req *ganymedeservice.OAuthLoginRequest) *xerror.Error {
	clientID, clientSecret, redirectUrl, accessTokenUrl, userProfileUrl, err := getOauthConf(req.Provider)
	if err != nil {
		return err
	}

	logger.Infof("clientID: %s, clientSecret: %s, redirectUrl: %s, accessTokenUrl: %s, userProfileUrl: %s",
		clientID, clientSecret, redirectUrl, accessTokenUrl, userProfileUrl)

	oAuth, err := oauth.NewOAuth(req.Provider, clientID, clientSecret, redirectUrl)
	if err != nil {
		return err
	}

	token, err := oAuth.GetAccessToken(accessTokenUrl, req.Code)
	if err != nil {
		return err
	}

	userProfile, err := oAuth.GetUserProfile(userProfileUrl, token.AccessToken)
	if err != nil {
		return err
	}

	err = saveUserInfo(ctx, req.Provider, req.Code, token, userProfile)
	if err != nil {
		return err
	}

	return nil
}

type LoginDomain struct {
	Code     string
	Provider string
}

func getOauthConf(provider string) (clientID, clientSecret, redirectUrl, accessTokenUrl, userProfileUrl string, err *xerror.Error) {
	switch provider {
	case github.ProviderGitHub:
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

func saveUserInfo(ctx context.Context, provider, code string, token *github.AccessTokenResponse, profile *github.UserProfile) *xerror.Error {
	tx, err := model.GanymedeDB.Beginx()
	if err != nil {
		return xerror.NewError(err, xerror.Code.SSqlExecuteErr, "bing tx failed")
	}
	defer tx.Rollback()

	user := &ganymede.User{
		Name:  profile.Name,
		Login: profile.Login,
	}
	user, xErr := model.InsertUser(ctx, tx, user)
	if xErr != nil {
		return xErr
	}

	oauth := &ganymede.OauthInfo{
		UseID:       user.ID,
		Provider:    provider,
		OuterID:     profile.ID,
		Login:       profile.Login,
		AvatarUrl:   profile.AvatarUrl,
		Email:       profile.Email,
		AccessToken: token.AccessToken,
		Code:        code,
	}
	oauth, xErr = model.InsertOAuthInfo(ctx, tx, oauth)
	if xErr != nil {
		return xErr
	}

	err = tx.Commit()
	if err != nil {
		return xerror.NewErrorf(err, xerror.Code.SSqlExecuteErr, "commit failed")
	}

	return nil
}
