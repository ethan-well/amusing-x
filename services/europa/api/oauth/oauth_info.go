package oauth

import (
	"amusingx.fit/amusingx/services/europa/app/oauth"
	"context"
	"github.com/ItsWewin/superfactory/httputil/rest"
	"github.com/ItsWewin/superfactory/logger"
	"github.com/ItsWewin/superfactory/xerror"
	"net/http"
)

func HandlerGetOAuthInfo(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	provider, err := getAndValidParams(ctx, r)
	if err != nil {
		rest.FailJsonResponse(w, err.Code, err.Message)
		return
	}

	resp, err := oauth.GetOAuthInfo(ctx, provider)
	if err != nil {
		logger.Errorf("get oauth info failed: %s", err)
		rest.FailJsonResponse(w, err.Code, err.Message)
	}

	rest.SucceedJsonResponse(w, resp)
}

func getAndValidParams(ctx context.Context, r *http.Request) (string, *xerror.Error) {
	provider := r.FormValue("provider")
	if len(provider) == 0 {
		return "", xerror.NewErrorf(nil, xerror.Code.CParamsError, "params 'provider' is invalid")
	}

	return provider, nil
}
