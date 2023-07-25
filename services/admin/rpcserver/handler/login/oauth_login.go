package login

import (
	panguservice "amusingx.fit/amusingx/protos/pangu/service/pangu/proto"
	ganymedeservice "amusingx.fit/amusingx/protos/user/service/user/proto"
	"amusingx.fit/amusingx/rpcclient/user"
	"context"
	"github.com/ItsWewin/superfactory/aerror"
)

func HandlerOauthLogin(ctx context.Context, req *panguservice.OAuthLoginRequest) (*panguservice.OAuthLoginResponse, aerror.Error) {
	rpcReq := &ganymedeservice.OAuthLoginRequest{
		Provider: req.Provider,
		Code:     req.Code,
		Service:  req.Service,
	}

	resp, err := user.Client.OAuthLogin(ctx, rpcReq)
	if err != nil || !resp.Result || resp.LoginInfo == nil {
		return nil, aerror.NewErrorf(err, aerror.Code.BUnexpectedData, "login failed")
	}

	return &panguservice.OAuthLoginResponse{
		Succeed: true,
		Result: &panguservice.LoginInfo{
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
