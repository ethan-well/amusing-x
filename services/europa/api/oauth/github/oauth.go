package github

import (
	"amusingx.fit/amusingx/apistruct/europa"
	"amusingx.fit/amusingx/services/europa/app/oauth"
	"context"
	"github.com/ItsWewin/superfactory/httputil"
	"github.com/ItsWewin/superfactory/httputil/rest"
	"github.com/ItsWewin/superfactory/logger"
	"github.com/ItsWewin/superfactory/xerror"
	"net/http"
)

func HandlerOauthLogin(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	req, err := getAndValidParams(r)
	if err != nil {
		logger.Errorf("get and valid params failed, err: %s", err.Error())

		rest.FailJsonResponse(w, xerror.Code.CParamsError, err.Message)
		return
	}

	resp, err := oauth.Login(ctx, w, req.Provider, req.Code)
	if err != nil {
		logger.Infof("oAuth login failed: %s", err)
		rest.FailJsonResponse(w, err.Code, err.Error())
		return
	}

	rest.SucceedJsonResponse(w, resp.LoginInfo.UserInfo)
}

func getAndValidParams(r *http.Request) (*europa.OAuthLogin, *xerror.Error) {
	var req europa.OAuthLogin

	err := httputil.DecodeJsonBody(r, &req)
	if err != nil {
		return nil, xerror.NewErrorf(err, xerror.Code.CParamsError, err.Error())
	}

	xErr := req.Valid()
	if xErr != nil {
		return nil, xErr
	}

	return &req, nil
}
