package login

import (
	"amusingx.fit/amusingx/apistruct/amusinguserserv"
	"amusingx.fit/amusingx/services/amusinguserserv/modle"
	"amusingx.fit/amusingx/xerror"
	"context"
	"github.com/ItsWewin/superfactory/httputil"
	"github.com/ItsWewin/superfactory/httputil/rest"
	"net/http"
)

func HandlerLogin(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	_, err := getAndValidParams(r)
	if err != nil {
		rest.FailJsonResponse(w, xerror.Code.ParamsError, xerror.Message.ParamsError)

		return
	}

	_, err = modle.FindUserByID(ctx, 1)
	if err != nil {
		w.Write([]byte(err.Error()))

		return
	}

	return
}

func getAndValidParams(r *http.Request) (*amusinguserserv.LoginRequest, error) {
	login := &amusinguserserv.LoginRequest{}

	err := httputil.DecodeJsonBody(r, login)
	return login, err
}
