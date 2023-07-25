package oauthlogin

import (
	"amusingx.fit/amusingx/mysqlstruct/ganymede"
	service "amusingx.fit/amusingx/protos/user/service/user/proto"
	"amusingx.fit/amusingx/services/user/conf"
	"amusingx.fit/amusingx/services/user/mysql/ganymededb/model"
	"amusingx.fit/amusingx/services/user/oauth"
	"amusingx.fit/amusingx/services/user/oauth/oauthstruct"
	model2 "amusingx.fit/amusingx/services/user/rpcserver/userservice/model"
	"amusingx.fit/amusingx/services/user/rpcserver/userservice/session"
	session2 "amusingx.fit/amusingx/services/user/session"
	"context"
	"github.com/ItsWewin/superfactory/aerror"
)

func HandlerOAuthLogin(ctx context.Context, req *service.OAuthLoginRequest) (*service.OAuthLoginResponse, aerror.Error) {
	err := getAndValidRequest(req)
	if err != nil {
		return nil, err
	}

	loginInfo, err := oauthLogin(ctx, req)
	if err != nil {
		return &service.OAuthLoginResponse{Result: false}, err
	}

	if err := setRoles(ctx, loginInfo, req.Service); err != nil {
		return &service.OAuthLoginResponse{Result: false}, err
	}

	return &service.OAuthLoginResponse{Result: true, LoginInfo: loginInfo}, nil
}

func getAndValidRequest(req *service.OAuthLoginRequest) aerror.Error {
	xErr := req.Valid()
	if xErr != nil {
		return xErr
	}

	return nil
}

func oauthLogin(ctx context.Context, req *service.OAuthLoginRequest) (*service.LoginInfo, aerror.Error) {

	var loginInfo = &service.LoginInfo{}

	clientID, clientSecret, redirectUrl, accessTokenUrl, _, userProfileUrl, grantType, err := getOauthConf(ctx, req.Provider, req.Service)
	if err != nil {
		return loginInfo, err
	}

	oAuth, err := oauth.NewOAuth(req.Provider, clientID, clientSecret, redirectUrl, grantType)
	if err != nil {
		return loginInfo, err
	}

	token, err := oAuth.GetAccessToken(accessTokenUrl, req.Code)
	if err != nil {
		return loginInfo, err
	}

	userProfile, err := oAuth.GetUserProfile(userProfileUrl, token.AccessToken)
	if err != nil {
		return loginInfo, err
	}

	user, err := saveUserInfo(ctx, req.Provider, req.Code, token, userProfile)
	if err != nil {
		return loginInfo, err
	}

	sessionID, err := setSession(ctx, user.ID)
	if err != nil {
		return loginInfo, err
	}

	loginInfo.SessionInfo = &service.SessionInfo{
		SessionID: sessionID,
		MaxAge:    int64(session2.MaxAge),
	}

	loginInfo.SessionId = sessionID
	loginInfo.UserInfo = &service.UserInfo{
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

func getOauthConf(ctx context.Context, provider, service string) (clientID, clientSecret, redirectUrl, accessTokenUrl, refreshTokenUrl, userProfileUrl, grantType string, err aerror.Error) {
	var p *conf.OAuthProviderInfo
	p, err = GetProvider(ctx, provider, service)
	if err != nil {
		return
	}
	clientID = p.ClientID
	clientSecret = p.ClientSecret
	redirectUrl = p.RedirectUrl
	accessTokenUrl = p.AccessTokenUrl
	userProfileUrl = p.UserProfileUrl
	refreshTokenUrl = p.RefreshTokenUrl
	grantType = p.GrantType

	return
}

func saveUserInfo(ctx context.Context, provider, code string, token *oauthstruct.AccessToken, profile *oauthstruct.UserProfile) (*ganymede.User, aerror.Error) {
	tx, err := model.UserDB.Beginx()
	if err != nil {
		return nil, aerror.NewError(err, aerror.Code.SSqlExecuteErr, "bing tx failed")
	}
	defer tx.Rollback()

	oauth := &ganymede.OauthInfo{
		Provider:    provider,
		OuterID:     profile.OuterUserID,
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
		return nil, aerror.NewErrorf(err, aerror.Code.SSqlExecuteErr, "commit failed")
	}

	return user, nil
}

func setSession(ctx context.Context, userID int64) (string, aerror.Error) {
	info := &session.Info{UserID: userID}

	model := session.New()
	return model.SetSession(ctx, info)
}

func setRoles(ctx context.Context, info *service.LoginInfo, service string) aerror.Error {
	model := model2.NewUserModel()
	roles, err := model.GetUserRoles(ctx, info.UserInfo.Id, service)
	if err != nil {
		return aerror.NewErrorf(nil, err.Code(), "set roles failed")
	}
	info.UserInfo.Roles = roles

	return nil
}
