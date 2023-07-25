package github

import (
	"amusingx.fit/amusingx/apistruct/apiService"
	"amusingx.fit/amusingx/services/webApi/app/oauth"
	"context"
	"github.com/ItsWewin/superfactory/aerror"
	"github.com/ItsWewin/superfactory/httputil"
	"github.com/ItsWewin/superfactory/httputil/rest"
	"github.com/ItsWewin/superfactory/logger"
	"net/http"
)

func HandlerOauthLogin(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	req, err := getAndValidParams(r)
	if err != nil {
		logger.Errorf("get and valid params failed, err: %s", err.Error())

		rest.FailJsonResponse(w, aerror.Code.CParamsErr, err.Message())
		return
	}

	resp, err := oauth.Login(ctx, w, req.Provider, req.Code, req.Service)
	if err != nil {
		logger.Errorf("oAuth login failed: %s", err)
		rest.FailJsonResponse(w, err.Code(), err.Error())
		return
	}

	rest.SucceedJsonResponse(w, resp.LoginInfo.UserInfo)
}

func getAndValidParams(r *http.Request) (*apiService.OAuthLogin, aerror.Error) {
	var req apiService.OAuthLogin

	err := httputil.DecodeJsonBody(r, &req)
	if err != nil {
		return nil, aerror.NewErrorf(err, aerror.Code.CParamsErr, err.Error())
	}

	xErr := req.Valid()
	if xErr != nil {
		return nil, xErr
	}

	return &req, nil
}
