package login

import (
	ganymedeservice "amusingx.fit/amusingx/protos/user/service/user/proto"
	"amusingx.fit/amusingx/services/webApi/rpcclient/user"
	"context"
	"github.com/ItsWewin/superfactory/aerror"
	"github.com/ItsWewin/superfactory/httputil"
	"github.com/ItsWewin/superfactory/httputil/rest"
	"github.com/ItsWewin/superfactory/logger"
	"net/http"
)

func HandlerLogin(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	req, xErr := getAndValidParams(r)
	if xErr != nil {
		logger.Errorf("get and valid params failed, err: %s", xErr.Error())

		rest.FailJsonResponse(w, aerror.Code.CParamsErr, aerror.Message.ParamsError)
		return
	}

	logger.Infof("HandlerLogin getAndValidParams: %s", logger.ToJson(req))
	response, err := user.RPCClient.Client.Login(ctx, req)
	if err != nil {
		logger.Errorf("Login failed, err: %s", err)

		rest.FailJsonResponse(w, aerror.Code.CParamsErr, aerror.Message.ParamsError)
		return
	}

	rest.SucceedJsonResponse(w, response)

	return
}

func getAndValidParams(r *http.Request) (*ganymedeservice.LoginRequest, aerror.Error) {
	login := &ganymedeservice.LoginRequest{}

	err := httputil.DecodeJsonBody(r, login)
	if err != nil {
		return nil, aerror.NewError(err, aerror.Code.SUnexpectedErr, "Unexpect error. ")
	}

	xErr := login.Valid()
	if err != nil {
		return nil, xErr
	}

	return login, nil
}
