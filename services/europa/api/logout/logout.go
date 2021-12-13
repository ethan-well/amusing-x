package logout

import (
	ganymedeservice "amusingx.fit/amusingx/protos/ganymede/service"
	"amusingx.fit/amusingx/services/europa/rpcclient/ganymede"
	"amusingx.fit/amusingx/services/europa/session"
	"context"
	"github.com/ItsWewin/superfactory/httputil/rest"
	"net/http"
)

func HandlerLogOut(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	sid, ok := session.GetSessionID(r)
	if !ok && sid == "" {
		rest.SucceedJsonResponse(w, "logout succeed")
		return
	}

	req := &ganymedeservice.LogoutRequest{SessionID: sid}

	resp, err := ganymede.RPCClient.Client.LogOut(ctx, req)
	if err != nil {
		rest.SucceedJsonResponse(w, "logout failed")
	}

	rest.SucceedJsonResponse(w, resp)
}
