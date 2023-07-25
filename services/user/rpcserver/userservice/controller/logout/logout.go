package logout

import (
	service "amusingx.fit/amusingx/protos/user/service/user/proto"
	"amusingx.fit/amusingx/services/user/rpcserver/userservice/session"
	"context"
	"github.com/ItsWewin/superfactory/aerror"
)

func HandlerIsLogOut(ctx context.Context, req *service.LogoutRequest) (*service.LogoutResponse, error) {
	if len(req.SessionID) == 0 {
		return nil, aerror.NewError(nil, aerror.Code.CParamsErr, "session is is invalid")
	}

	resp := &service.LogoutResponse{Logout: false}

	model := session.New()
	err := model.Delete(ctx, req.SessionID)
	if err != nil {
		return resp, aerror.NewErrorf(nil, err.Code(), err.Message())
	}

	resp.Logout = true

	return resp, nil
}
