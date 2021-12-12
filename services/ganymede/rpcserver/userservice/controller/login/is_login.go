package login

import (
	ganymedeservice "amusingx.fit/amusingx/protos/ganymede/service"
	"amusingx.fit/amusingx/services/ganymede/rpcserver/userservice/session"
	"context"
	"github.com/ItsWewin/superfactory/xerror"
)

func HandlerIsLogin(ctx context.Context, req *ganymedeservice.IsLoginRequest) (*ganymedeservice.IsLoginResponse, error) {
	if len(req.SessionID) == 0 {
		return nil, xerror.NewError(nil, xerror.Code.CParamsError, "session is is invalid")
	}

	resp := &ganymedeservice.IsLoginResponse{Login: false}

	model := session.New()
	id, err := model.GetUserID(ctx, req.SessionID)
	if err != nil || id <= 0 {
		return resp, xerror.NewErrorf(nil, err.Code, err.Message)
	}

	resp.Login = true
	return resp, nil
}
