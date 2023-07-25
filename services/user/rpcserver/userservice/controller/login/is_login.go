package login

import (
	service "amusingx.fit/amusingx/protos/user/service/user/proto"
	model2 "amusingx.fit/amusingx/services/user/rpcserver/userservice/model"
	"amusingx.fit/amusingx/services/user/rpcserver/userservice/session"
	"context"
	"github.com/ItsWewin/superfactory/aerror"
	"github.com/ItsWewin/superfactory/logger"
)

func HandlerIsLogin(ctx context.Context, req *service.IsLoginRequest) (*service.IsLoginResponse, error) {
	if len(req.SessionID) == 0 {
		return nil, aerror.NewError(nil, aerror.Code.CParamsErr, "session is is invalid")
	}

	resp := &service.IsLoginResponse{Login: false}

	model := session.New()
	id, err := model.GetUserID(ctx, req.SessionID)
	if err != nil || id <= 0 {
		return resp, aerror.NewErrorf(nil, err.Code(), err.Message())
	}

	logger.Infof("user id: %v", id)

	user, err := model2.NewUserModel().GetUserInfoByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if user == nil {
		resp.Login = false
		return resp, nil
	}

	resp.UserInfo = &service.UserInfo{
		Id:    id,
		Name:  user.Name,
		Login: user.Login,
		Email: user.Login,
	}
	resp.Login = true

	return resp, nil
}
