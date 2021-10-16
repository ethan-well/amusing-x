package login

import (
	"amusingx.fit/amusingx/protos/amusingxuserserv/userservice"
	"amusingx.fit/amusingx/services/amusingwebapiserv/rpcclient/userrpcserver"
	"context"
	"github.com/ItsWewin/superfactory/httputil"
	"github.com/ItsWewin/superfactory/httputil/rest"
	"github.com/ItsWewin/superfactory/logger"
	"github.com/ItsWewin/superfactory/xerror"
	"net/http"
)

func HandlerLogin(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	req, xErr := getAndValidParams(r)
	if xErr != nil {
		logger.Errorf("get and valid params failed, err: %s", xErr.Error())

		rest.FailJsonResponse(w, xerror.Code.CParamsError, xerror.Message.ParamsError)
		return
	}

	logger.Infof("HandlerLogin getAndValidParams: %s", logger.ToJson(req))
	response, err := userrpcserver.RPCClient.Client.Login(ctx, req)
	if err != nil {
		logger.Errorf("Login failed, err: %s", err)

		rest.FailJsonResponse(w, xerror.Code.CParamsError, xerror.Message.ParamsError)
		return
	}

	rest.SucceedJsonResponse(w, response)

	return
}

func getAndValidParams(r *http.Request) (*userservice.LoginRequest, *xerror.Error) {
	login := &userservice.LoginRequest{}

	err := httputil.DecodeJsonBody(r, login)
	if err != nil {
		return nil, xerror.NewError(err, xerror.Code.SUnexpectedErr, "Unexpect error. ")
	}

	logger.Infof("getAndValidParams login: %s", logger.ToJson(login))

	xErr := login.Valid()
	if err != nil {
		return nil, xErr
	}

	return login, nil
}
