package login

import (
	"amusingx.fit/amusingx/apistruct/amusinguserserv"
	"amusingx.fit/amusingx/services/amusinguserserv/app/loginapp"
	"amusingx.fit/amusingx/services/amusinguserserv/session"
	"context"
	"github.com/ItsWewin/superfactory/httputil"
	"github.com/ItsWewin/superfactory/httputil/rest"
	"github.com/ItsWewin/superfactory/logger"
	"github.com/ItsWewin/superfactory/xerror"
	"net/http"
	"time"
)

func HandlerLogin(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	req, err := getAndValidParams(r)
	if err != nil {
		logger.Errorf("get and valid params failed, err: %s", err.Error())

		rest.FailJsonResponse(w, xerror.Code.CParamsError, xerror.Message.ParamsError)
		return
	}

	err = Login(ctx, w, r, req)
	if err != nil {
		logger.Errorf("login failed, err: %s", err.Error())
		rest.FailJsonResponse(w, err.Code, err.Message)
		return
	}

	rest.SucceedJsonResponse(w, "login succeed")

	return
}

func Login(ctx context.Context, w http.ResponseWriter, r *http.Request, loginRequest *amusinguserserv.LoginRequest) *xerror.Error {
	loginDomain := loginapp.NewDomain()

	err := loginDomain.SetLoginRequestInfo(loginRequest)
	if err != nil {
		return err
	}

	err = loginDomain.SetUserModelInfo(ctx)
	if err != nil {
		return err
	}

	err = loginDomain.ValidPassword(ctx)
	if err != nil {
		return err
	}

	err = loginDomain.SetSession(ctx)
	if err != nil {
		return err
	}

	c1 := &http.Cookie{
		Name:     session.GlobalSessionManager.CookieName,
		Value:    loginDomain.SessionID,
		Path:     "/",
		Expires:  time.Now().Add(time.Duration(session.GlobalSessionManager.MaxLifeTime) * time.Second),
		MaxAge:   session.GlobalSessionManager.MaxLifeTime,
		HttpOnly: true,
	}

	http.SetCookie(w, c1)

	return nil
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
