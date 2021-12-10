package oauthlogin

import (
	"amusingx.fit/amusingx/apistruct/github"
	"amusingx.fit/amusingx/mysqlstruct/ganymede"
	"amusingx.fit/amusingx/protos/ganymede/service"
	"amusingx.fit/amusingx/services/ganymede/conf"
	"amusingx.fit/amusingx/services/ganymede/mysql/ganymededb/model"
	"amusingx.fit/amusingx/services/ganymede/oauth"
	"amusingx.fit/amusingx/services/ganymede/rpcserver/userservice/session"
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
	loginInfo, err := oauthLogin(ctx, req)
	if err != nil {
		return &ganymedeservice.OAuthLoginResponse{Result: false}, err
	}

	return &ganymedeservice.OAuthLoginResponse{Result: true, LoginInfo: loginInfo}, nil
}

func getAndValidRequest(req *ganymedeservice.OAuthLoginRequest) *xerror.Error {
	xErr := req.Valid()
	if xErr != nil {
		return xErr
	}

	return nil
}

func oauthLogin(ctx context.Context, req *ganymedeservice.OAuthLoginRequest) (*ganymedeservice.LoginInfo, *xerror.Error) {

	var loginInfo = &ganymedeservice.LoginInfo{}

	clientID, clientSecret, redirectUrl, accessTokenUrl, userProfileUrl, err := getOauthConf(req.Provider)
	if err != nil {
		return loginInfo, err
	}

	logger.Infof("clientID: %s, clientSecret: %s, redirectUrl: %s, accessTokenUrl: %s, userProfileUrl: %s",
		clientID, clientSecret, redirectUrl, accessTokenUrl, userProfileUrl)

	oAuth, err := oauth.NewOAuth(req.Provider, clientID, clientSecret, redirectUrl)
	if err != nil {
		return loginInfo, err
	}

	token, err := oAuth.GetAccessToken(accessTokenUrl, req.Code)
	if err != nil {
		return loginInfo, err
	}

	logger.Infof("token: %s", logger.ToJson(token))

	userProfile, err := oAuth.GetUserProfile(userProfileUrl, token.AccessToken)
	if err != nil {
		return loginInfo, err
	}
	logger.Infof("userProfile: %s", logger.ToJson(userProfile))

	user, err := saveUserInfo(ctx, req.Provider, req.Code, token, userProfile)
	if err != nil {
		return loginInfo, err
	}

	sessionID, err := setSession(ctx, user.ID)
	if err != nil {
		return loginInfo, err
	}

	loginInfo.SessionId = sessionID
	loginInfo.UserInfo = &ganymedeservice.UserInfo{
		Id:    user.ID,
		Name:  user.Name,
		Login: user.Login,
	}

	return loginInfo, nil
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

func saveUserInfo(ctx context.Context, provider, code string, token *github.AccessTokenResponse, profile *github.UserProfile) (*ganymede.User, *xerror.Error) {
	tx, err := model.GanymedeDB.Beginx()
	if err != nil {
		return nil, xerror.NewError(err, xerror.Code.SSqlExecuteErr, "bing tx failed")
	}
	defer tx.Rollback()

	oauth := &ganymede.OauthInfo{
		Provider:    provider,
		OuterID:     profile.ID,
		Login:       profile.Login,
		AvatarUrl:   profile.AvatarUrl,
		Email:       profile.Email,
		AccessToken: token.AccessToken,
		Code:        code,
	}

	// 保存或者更新 oauth info
	oauth, xErr := model.InsertOrUpdateOAuthInfo(ctx, tx, oauth)
	if xErr != nil {
		return nil, xErr
	}

	// 加锁查询 oauth info 是不是已经关联了 user
	authInfo, xErr := model.QueryOAuthInfoByProviderAndLogin(ctx, tx, provider, profile.Login)
	if err != nil {
		return nil, xErr
	}

	var user *ganymede.User
	// 已经关联了用户
	if authInfo != nil && authInfo.UseID != 0 {
		user, xErr = model.QueryUserByID(ctx, tx, authInfo.UseID)
		if err != nil {
			return nil, xErr
		}
	}

	// 用户不存在
	if user == nil {
		user = &ganymede.User{
			Name:  profile.Name,
			Login: provider + profile.Login,
		}

		user, xErr = model.InsertUser(ctx, tx, user)
		if xErr != nil {
			return nil, xErr
		}
	}

	// auth info 关联用户
	if authInfo == nil || authInfo.UseID != user.ID {
		xErr = model.UpdateOAuthUserID(ctx, tx, user.ID, provider, profile.Login)
		if xErr != nil {
			return nil, xErr
		}
	}

	err = tx.Commit()
	if err != nil {
		return nil, xerror.NewErrorf(err, xerror.Code.SSqlExecuteErr, "commit failed")
	}

	return user, nil
}

func setSession(ctx context.Context, userID int64) (string, *xerror.Error) {
	info := &session.Info{UserID: userID}

	model := session.New()
	return model.SetSession(ctx, info)
}
