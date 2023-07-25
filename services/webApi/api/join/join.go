package join

import (
	"amusingx.fit/amusingx/apistruct/apiService"
	"amusingx.fit/amusingx/services/webApi/app/joinapp"
	"context"
	"github.com/ItsWewin/superfactory/aerror"
	"github.com/ItsWewin/superfactory/httputil"
	"github.com/ItsWewin/superfactory/httputil/rest"
	"github.com/ItsWewin/superfactory/logger"
	"net/http"
)

func HandleJoin(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	params, err := getAndValidParams(r)
	if err != nil {
		logger.Errorf("params is invalid: %s", err.Error())

		rest.FailJsonResponse(w, aerror.Code.CreateUserErr, err.Message())
		return
	}

	u, err := joinapp.CreateUser(ctx, params)
	if err != nil {
		logger.Errorf("create user failed, params: %s, err: %s", logger.ToJson(params), err.Error())
		rest.FailJsonResponse(w, aerror.Code.CreateUserErr, err.Message())
		return
	}

	resp := apiService.JoinResponse{
		Nickname: u.Nickname,
		Phone:    u.Phone,
	}

	rest.SucceedJsonResponse(w, resp)
}

func getAndValidParams(r *http.Request) (*apiService.JoinRequest, aerror.Error) {
	joinRequest := apiService.JoinRequest{}

	err := httputil.DecodeJsonBody(r, &joinRequest)
	if err != nil {
		return nil, aerror.NewError(err, aerror.Code.CParamsErr, "decode request body failed")
	}

	xErr := joinRequest.Valid()
	if xErr != nil {
		return nil, xErr
	}

	return &joinRequest, nil
}
