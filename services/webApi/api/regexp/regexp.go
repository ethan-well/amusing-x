package regexp

import (
	ganymedeservice "amusingx.fit/amusingx/protos/user/service/user/proto"
	"amusingx.fit/amusingx/services/webApi/rpcclient/user"
	"context"
	"github.com/ItsWewin/superfactory/aerror"
	"github.com/ItsWewin/superfactory/httputil/rest"
	"github.com/ItsWewin/superfactory/logger"
	"net/http"
)

func HandlerRegexp(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	req := &ganymedeservice.BlankParams{}

	regex, err := user.RPCClient.Client.Regexps(ctx, req)
	if err != nil {
		logger.Errorf("rpc failed, userrpcserver.RPCClient.Client.Regexps: %s", regex)
		rest.FailJsonResponse(w, aerror.Code.OtherNetworkErr, "网络错误，稍后重试")
	}

	rest.SucceedJsonResponse(w, regex)
}
