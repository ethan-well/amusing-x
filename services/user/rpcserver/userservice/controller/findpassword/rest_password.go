package password

import (
	service "amusingx.fit/amusingx/protos/user/service/user/proto"
	"amusingx.fit/amusingx/services/user/rpcserver/userservice/app/passwordapp"
	"context"
	"github.com/ItsWewin/superfactory/aerror"
)

func HandlerResetPassword(ctx context.Context,
	req *service.ResetPasswordRequest) (*service.ResetPasswordResponse, aerror.Error) {

	err := getAndValidRequest(req)
	if err != nil {
		return nil, err
	}

	err = resetPassword(ctx, req)
	if err != nil {
		return nil, err
	}

	return &service.ResetPasswordResponse{Result: "密码重置成功"}, nil
}

func resetPassword(ctx context.Context, req *service.ResetPasswordRequest) aerror.Error {
	domain := passwordapp.NewDomain()

	err := domain.SetResetPasswordInfo(req)
	if err != nil {
		return aerror.NewError(err, err.Code(), "SetResetPasswordInfo failed. ")
	}

	err = domain.SetUserModelInfo(ctx)
	if err != nil {
		return aerror.NewError(err, err.Code(), "SetUserModelInfo failed. ")
	}

	err = domain.ResetPassword(ctx, domain.ResetPasswordInfo.Password)
	if err != nil {
		return aerror.NewError(err, err.Code(), "ResetPassword failed. ")
	}

	return nil
}

func getAndValidRequest(request *service.ResetPasswordRequest) aerror.Error {
	xErr := request.Valid()
	if xErr != nil {
		return xErr
	}

	return nil
}
