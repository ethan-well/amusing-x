package login

import (
	service "amusingx.fit/amusingx/protos/user/service/user/proto"
	"amusingx.fit/amusingx/regexp"
	"amusingx.fit/amusingx/services/user/rpcserver/userservice/app/loginapp"
	"context"
	"github.com/ItsWewin/superfactory/aerror"
	"github.com/ItsWewin/superfactory/logger"
)

func HandlerLogin(ctx context.Context, req *service.LoginRequest) (string, aerror.Error) {
	if err := ParamsValid(req); err != nil {
		return "", aerror.NewError(nil, err.Code(), err.Message())
	}

	sessionID, err := Login(ctx, req)
	if err != nil {
		return "", aerror.NewError(nil, err.Code(), err.Message())
	}

	return sessionID, nil
}

func ParamsValid(x *service.LoginRequest) aerror.Error {
	if x.Type != 0 && x.Type != 1 {
		return aerror.NewError(nil, aerror.Code.CParamsErr, "Type is invalid. ")
	}

	if err := regexp.PhoneNumberValid(x.GetPhone()); err != nil {
		return err
	}

	if err := regexp.AreaCodeValid(x.AreaCode); err != nil {
		return err
	}

	if x.Type == 0 {
		if err := regexp.PasswordValid(x.Password); err != nil {
			return err
		}
	} else {
		if err := regexp.VerificationCodeValid(x.VerificationCode); err != nil {
			return err
		}
	}

	return nil
}

func Login(ctx context.Context, loginRequest *service.LoginRequest) (string, aerror.Error) {
	loginDomain := loginapp.NewDomain()

	logger.Infof("loginRequest: %s", logger.ToJson(loginRequest))

	err := loginDomain.SetLoginRequestInfo(loginRequest)
	if err != nil {
		return "", err
	}

	err = loginDomain.SetUserModelInfo(ctx)
	if err != nil {
		return "", err
	}

	err = loginDomain.LoginAuthentication(ctx)
	if err != nil {
		return "", err
	}

	err = loginDomain.SetSession(ctx)
	if err != nil {
		return "", err
	}

	return loginDomain.SessionID, nil
}
