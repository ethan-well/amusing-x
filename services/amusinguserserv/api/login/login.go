package login

import (
	"amusingx.fit/amusingx/apistruct/amusinguserserv"
	"amusingx.fit/amusingx/services/amusinguserserv/modle"
	"amusingx.fit/amusingx/xerror"
	"context"
	"github.com/ItsWewin/superfactory/httputil"
	"github.com/ItsWewin/superfactory/httputil/rest"
	"github.com/ItsWewin/superfactory/logger"
	"net/http"
)

func HandlerLogin(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	_, err := getAndValidParams(r)
	if err != nil {
		logger.Errorf("get and valid params failed, err: %s", err.Error())

		rest.FailJsonResponse(w, xerror.Code.CParamsError, xerror.Message.ParamsError)
		return
	}

	_, err = modle.FindUserByID(ctx, 1)
	if err != nil {
		w.Write([]byte(err.Error()))

		return
	}

	return
}

func getAndValidParams(r *http.Request) (*amusinguserserv.LoginRequest, *xerror.Error) {
	login := &amusinguserserv.LoginRequest{}

	err := httputil.DecodeJsonBody(r, login)
	if err != nil {
		return nil, xerror.NewError(err, xerror.Code.SUnexpectedErr, "Unexpect error. ")
	}

	xErr := login.Valid()
	if err != nil {
		return nil, xErr
	}

	return login, nil
}
