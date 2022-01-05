package logout

import (
	ganymedeservice "amusingx.fit/amusingx/protos/ganymede/service/ganymede/proto"
	"amusingx.fit/amusingx/services/ganymede/rpcserver/userservice/session"
	"context"
	"github.com/ItsWewin/superfactory/aerror"
)

func HandlerIsLogOut(ctx context.Context, req *ganymedeservice.LogoutRequest) (*ganymedeservice.LogoutResponse, error) {
	if len(req.SessionID) == 0 {
		return nil, aerror.NewError(nil, aerror.Code.CParamsError, "session is is invalid")
	}

	resp := &ganymedeservice.LogoutResponse{Logout: false}

	model := session.New()
	err := model.Delete(ctx, req.SessionID)
	if err != nil {
		return resp, aerror.NewErrorf(nil, err.Code(), err.Message())
	}

	resp.Logout = true

	return resp, nil
}
