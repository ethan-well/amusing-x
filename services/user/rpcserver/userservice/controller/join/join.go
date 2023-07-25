package join

import (
	service "amusingx.fit/amusingx/protos/user/service/user/proto"
	"amusingx.fit/amusingx/services/user/rpcserver/userservice/app/joinapp"
	"context"
	"github.com/ItsWewin/superfactory/aerror"
)

func HandleJoin(ctx context.Context, joinRequest *service.JoinRequest) (*service.JoinResponse, aerror.Error) {
	err := getAndValidParams(joinRequest)
	if err != nil {
		return nil, aerror.NewError(err, err.Code(), err.Message())
	}

	u, err := joinapp.CreateUser(ctx, joinRequest)
	if err != nil {
		return nil, aerror.NewError(err, err.Code(), err.Message())
	}

	resp := &service.JoinResponse{
		Nickname: u.Nickname,
		Phone:    u.Phone,
		AreaCode: u.AreaCode,
	}

	return resp, nil
}

func getAndValidParams(joinRequest *service.JoinRequest) aerror.Error {
	xErr := joinRequest.Valid()
	if xErr != nil {
		return xErr
	}

	return nil
}
