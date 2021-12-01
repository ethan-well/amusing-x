package password

import (
	"amusingx.fit/amusingx/apistruct/amusinguserserv"
	"amusingx.fit/amusingx/services/europa/app/passwordapp"
	"context"
	"github.com/ItsWewin/superfactory/httputil"
	"github.com/ItsWewin/superfactory/httputil/rest"
	"github.com/ItsWewin/superfactory/logger"
	"github.com/ItsWewin/superfactory/xerror"
	"net/http"
)

func HandlerResetPassword(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	req, err := getAndValidParams(r)
	if err != nil {
		logger.Errorf("params is invalid: %s", err.Error())

		rest.FailJsonResponse(w, xerror.Code.CUnexpectRequestDate, err.Message)
		return
	}

	err = resetPassword(ctx, req)
	if err != nil {
		logger.Errorf("reset password failed: %s", err.Error())

		rest.FailJsonResponse(w, err.Code, err.Message)
		return
	}

	rest.SucceedJsonResponse(w, "密码重置成功")
}

func resetPassword(ctx context.Context, req *amusinguserserv.ResetPasswordRequest) *xerror.Error {
	domain := passwordapp.NewDomain()

	err := domain.SetResetPasswordInfo(req)
	if err != nil {
		return xerror.NewError(err, err.Code, "SetResetPasswordInfo failed. ")
	}

	err = domain.SetUserModelInfo(ctx)
	if err != nil {
		return xerror.NewError(err, err.Code, "SetUserModelInfo failed. ")
	}

	err = domain.ResetPassword(ctx, domain.ResetPasswordInfo.Password)
	if err != nil {
		return xerror.NewError(err, err.Code, "ResetPassword failed. ")
	}

	return nil
}

func getAndValidParams(r *http.Request) (*amusinguserserv.ResetPasswordRequest, *xerror.Error) {
	resetPasswordRequest := amusinguserserv.ResetPasswordRequest{}

	err := httputil.DecodeJsonBody(r, &resetPasswordRequest)
	if err != nil {
		return nil, xerror.NewError(err, xerror.Code.CParamsError, "decode request body failed")
	}

	xErr := resetPasswordRequest.Valid()
	if xErr != nil {
		return nil, xErr
	}

	return &resetPasswordRequest, nil
}
