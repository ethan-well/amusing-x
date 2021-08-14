package join

import (
	"amusingx.fit/amusingx/protos/amusingxuserserv/userservice"
	"amusingx.fit/amusingx/services/amusinguserserv/rpcserver/userservice/handler/joinapp"
	"context"
	"github.com/ItsWewin/superfactory/xerror"
)

func HandleJoin(ctx context.Context, joinRequest *userservice.JoinRequest) (*userservice.JoinResponse, *xerror.Error) {
	err := getAndValidParams(joinRequest)
	if err != nil {
		return nil, xerror.NewError(err, err.Code, err.Message)
	}

	u, err := joinapp.CreateUser(ctx, joinRequest)
	if err != nil {
		return nil, xerror.NewError(err, err.Code, err.Message)
	}

	resp := &userservice.JoinResponse{
		Nickname: u.Nickname,
		Phone:    u.Phone,
		AreaCode: u.AreaCode,
	}

	return resp, nil
}

func getAndValidParams(joinRequest *userservice.JoinRequest) *xerror.Error {
	xErr := joinRequest.Valid()
	if xErr != nil {
		return xErr
	}

	return nil
}
