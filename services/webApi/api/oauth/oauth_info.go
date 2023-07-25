package oauth

import (
	"amusingx.fit/amusingx/services/webApi/app/oauth"
	"context"
	"github.com/ItsWewin/superfactory/aerror"
	"github.com/ItsWewin/superfactory/httputil/rest"
	"github.com/ItsWewin/superfactory/logger"
	"net/http"
)

func HandlerGetOAuthInfo(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	provider, service, err := getAndValidParams(ctx, r)
	if err != nil {
		rest.FailJsonResponse(w, err.Code(), err.Message())
		return
	}

	resp, err := oauth.GetOAuthInfo(ctx, provider, service)
	if err != nil {
		logger.Errorf("get oauth info failed: %s", err)
		rest.FailJsonResponse(w, err.Code(), err.Message())
	}

	rest.SucceedJsonResponse(w, resp)
}

func getAndValidParams(ctx context.Context, r *http.Request) (string, string, aerror.Error) {
	provider := r.FormValue("provider")
	if len(provider) == 0 {
		return "", "", aerror.NewErrorf(nil, aerror.Code.CParamsErr, "params 'provider' is invalid")
	}

	service := r.FormValue("service")
	if len(service) == 0 {
		return "", "", aerror.NewErrorf(nil, aerror.Code.CParamsErr, "params 'service' is invalid")
	}

	return provider, service, nil
}
