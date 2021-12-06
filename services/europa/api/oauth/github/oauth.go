package github

import (
	"amusingx.fit/amusingx/apistruct/europa"
	"context"
	"encoding/json"
	"github.com/ItsWewin/superfactory/httputil/rest"
	"github.com/ItsWewin/superfactory/logger"
	"github.com/ItsWewin/superfactory/xerror"
	"net/http"
)

func HandlerOauthLogin(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	req, xErr := getAndValidParams(r)
	if xErr != nil {
		logger.Errorf("get and valid params failed, err: %s", xErr.Error())

		rest.FailJsonResponse(w, xerror.Code.CParamsError, xerror.Message.ParamsError)
		return
	}

}

func getAndValidParams(r http.Request) (*europa.OAuthLogin, *xerror.Error) {
	code := r.FormValue("code")
	provider := r.FormValue("provider")
	if len(code) == 0 {
		return nil, xerror.NewErrorf(nil, xerror.Code.CParamsError, "params 'code' is expected")
	}

	if len(provider) == 0 {
		return nil, xerror.NewErrorf(nil, xerror.Code.CParamsError, "params 'code' is expected")
	}

	return &europa.OAuthLogin{
		Code:     code,
		Provider: provider,
	}, nil
}
