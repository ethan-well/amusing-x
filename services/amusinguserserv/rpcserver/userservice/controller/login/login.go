package login

import (
	"amusingx.fit/amusingx/protos/amusingxuserserv/userservice"
	"amusingx.fit/amusingx/services/amusinguserserv/regexp"
	"amusingx.fit/amusingx/services/amusinguserserv/rpcserver/userservice/handler/loginapp"
	"context"
	"github.com/ItsWewin/superfactory/xerror"
)

func HandlerLogin(ctx context.Context, req *userservice.LoginRequest) (string, *xerror.Error) {
	if err := ParamsValid(req); err != nil {
		return "", xerror.NewError(nil, err.Code, err.Message)
	}

	sessionID, err := Login(ctx, req)
	if err != nil {
		return "", xerror.NewError(nil, err.Code, err.Message)
	}

	return sessionID, nil
}

func ParamsValid(x *userservice.LoginRequest) *xerror.Error {
	if x.Type != 0 && x.Type != 1 {
		return xerror.NewError(nil, xerror.Code.CParamsError, "Type is invalid. ")
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

func Login(ctx context.Context, loginRequest *userservice.LoginRequest) (string, *xerror.Error) {
	loginDomain := loginapp.NewDomain()

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
