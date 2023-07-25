package password

import (
	"amusingx.fit/amusingx/apistruct/apiService"
	"amusingx.fit/amusingx/services/webApi/app/passwordapp"
	"context"
	"github.com/ItsWewin/superfactory/aerror"
	"github.com/ItsWewin/superfactory/httputil"
	"github.com/ItsWewin/superfactory/httputil/rest"
	"github.com/ItsWewin/superfactory/logger"
	"net/http"
)

func HandlerResetPassword(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	req, err := getAndValidParams(r)
	if err != nil {
		logger.Errorf("params is invalid: %s", err.Error())

		rest.FailJsonResponse(w, aerror.Code.CUnexpectRequestDate, err.Message())
		return
	}

	err = resetPassword(ctx, req)
	if err != nil {
		logger.Errorf("reset password failed: %s", err.Error())

		rest.FailJsonResponse(w, err.Code(), err.Message())
		return
	}

	rest.SucceedJsonResponse(w, "密码重置成功")
}

func resetPassword(ctx context.Context, req *apiService.ResetPasswordRequest) aerror.Error {
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

func getAndValidParams(r *http.Request) (*apiService.ResetPasswordRequest, aerror.Error) {
	resetPasswordRequest := apiService.ResetPasswordRequest{}

	err := httputil.DecodeJsonBody(r, &resetPasswordRequest)
	if err != nil {
		return nil, aerror.NewError(err, aerror.Code.CParamsErr, "decode request body failed")
	}

	xErr := resetPasswordRequest.Valid()
	if xErr != nil {
		return nil, xErr
	}

	return &resetPasswordRequest, nil
}
