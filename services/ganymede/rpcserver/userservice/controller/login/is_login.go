package login

import (
	ganymedeservice "amusingx.fit/amusingx/protos/ganymede/service/ganymede/proto"
	model2 "amusingx.fit/amusingx/services/ganymede/rpcserver/userservice/model"
	"amusingx.fit/amusingx/services/ganymede/rpcserver/userservice/session"
	"context"
	"github.com/ItsWewin/superfactory/aerror"
)

func HandlerIsLogin(ctx context.Context, req *ganymedeservice.IsLoginRequest) (*ganymedeservice.IsLoginResponse, error) {
	if len(req.SessionID) == 0 {
		return nil, aerror.NewError(nil, aerror.Code.CParamsError, "session is is invalid")
	}

	resp := &ganymedeservice.IsLoginResponse{Login: false}

	model := session.New()
	id, err := model.GetUserID(ctx, req.SessionID)
	if err != nil || id <= 0 {
		return resp, aerror.NewErrorf(nil, err.Code(), err.Message())
	}

	user, err := model2.NewUserModel().GetUserInfoByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if user == nil {
		resp.Login = false
		return resp, nil
	}

	resp.UserInfo = &ganymedeservice.UserInfo{
		Id:    id,
		Name:  user.Name,
		Login: user.Login,
		Email: "",
	}
	resp.Login = true

	return resp, nil
}
