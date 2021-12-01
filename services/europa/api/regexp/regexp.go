package regexp

import (
	"amusingx.fit/amusingx/protos/ganymede/service"
	"amusingx.fit/amusingx/services/europa/rpcclient/userrpcserver"
	"context"
	"github.com/ItsWewin/superfactory/httputil/rest"
	"github.com/ItsWewin/superfactory/logger"
	"github.com/ItsWewin/superfactory/xerror"
	"net/http"
)

func HandlerRegexp(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	req := &userservice.BlankParams{}

	regex, err := userrpcserver.RPCClient.Client.Regexps(ctx, req)
	if err != nil {
		logger.Errorf("rpc failed, userrpcserver.RPCClient.Client.Regexps: %s", regex)
		rest.FailJsonResponse(w, xerror.Code.OtherNetworkError, "网络错误，稍后重试")
	}

	rest.SucceedJsonResponse(w, regex)
}
