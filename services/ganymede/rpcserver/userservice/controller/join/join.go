package join

import (
	"amusingx.fit/amusingx/protos/ganymede/service"
	"amusingx.fit/amusingx/services/ganymede/rpcserver/userservice/app/joinapp"
	"context"
	"github.com/ItsWewin/superfactory/aerror"
)

func HandleJoin(ctx context.Context, joinRequest *ganymedeservice.JoinRequest) (*ganymedeservice.JoinResponse, aerror.Error) {
	err := getAndValidParams(joinRequest)
	if err != nil {
		return nil, aerror.NewError(err, err.Code(), err.Message())
	}

	u, err := joinapp.CreateUser(ctx, joinRequest)
	if err != nil {
		return nil, aerror.NewError(err, err.Code(), err.Message())
	}

	resp := &ganymedeservice.JoinResponse{
		Nickname: u.Nickname,
		Phone:    u.Phone,
		AreaCode: u.AreaCode,
	}

	return resp, nil
}

func getAndValidParams(joinRequest *ganymedeservice.JoinRequest) aerror.Error {
	xErr := joinRequest.Valid()
	if xErr != nil {
		return xErr
	}

	return nil
}
