package login

import (
	ganymedeservice "amusingx.fit/amusingx/protos/ganymede/service/ganymede/proto"
	panguservice "amusingx.fit/amusingx/protos/pangu/service/pangu/proto"
	"amusingx.fit/amusingx/rpcclient/ganymede"
	"context"
	"github.com/ItsWewin/superfactory/aerror"
)

func HandlerOauthLogin(ctx context.Context, req *panguservice.OAuthLoginRequest) (*panguservice.OAuthLoginResponse, aerror.Error) {
	rpcReq := &ganymedeservice.OAuthLoginRequest{
		Provider: req.Provider,
		Code:     req.Code,
		Service:  req.Service,
	}

	resp, err := ganymede.Client.OAuthLogin(ctx, rpcReq)
	if err != nil || !resp.Result || resp.LoginInfo == nil {
		return nil, aerror.NewErrorf(err, aerror.Code.BUnexpectedData, "login failed")
	}

	return &panguservice.OAuthLoginResponse{
		Result: true,
		LoginInfo: &panguservice.LoginInfo{
			SessionId: resp.LoginInfo.SessionId,
			UserInfo: &panguservice.UserInfo{
				Id:    resp.LoginInfo.UserInfo.Id,
				Name:  resp.LoginInfo.UserInfo.Name,
				Login: resp.LoginInfo.UserInfo.Login,
				Email: resp.LoginInfo.UserInfo.Email,
				Roles: resp.LoginInfo.UserInfo.Roles,
			},
			SessionInfo: &panguservice.SessionInfo{
				SessionID: resp.LoginInfo.SessionInfo.SessionID,
				MaxAge:    resp.LoginInfo.SessionInfo.MaxAge,
			},
		},
	}, nil
}
