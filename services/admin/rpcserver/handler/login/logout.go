package login

import (
	panguservice "amusingx.fit/amusingx/protos/pangu/service/pangu/proto"
	ganymedeservice "amusingx.fit/amusingx/protos/user/service/user/proto"
	"amusingx.fit/amusingx/rpcclient/user"
	"context"
	"github.com/ItsWewin/superfactory/aerror"
)

func HandlerLogout(ctx context.Context, sessionID string) (*panguservice.LogoutResponse, aerror.Error) {
	resp, err := user.Client.LogOut(ctx, &ganymedeservice.LogoutRequest{SessionID: sessionID})
	if err != nil {
		return &panguservice.LogoutResponse{Succeed: false}, aerror.NewErrorf(err, aerror.Code.SUnexpectedErr, "logout failed")
	}

	return &panguservice.LogoutResponse{Succeed: resp.Logout}, nil
}

func HandlerIsLogin(ctx context.Context, sessionID string) (*ganymedeservice.IsLoginResponse, aerror.Error) {
	resp, err := user.Client.IsLogin(ctx, &ganymedeservice.IsLoginRequest{SessionID: sessionID})
	if err != nil {
		return &ganymedeservice.IsLoginResponse{Login: false}, aerror.NewErrorf(err, aerror.Code.SUnexpectedErr, "use not login")
	}

	return resp, nil
}
