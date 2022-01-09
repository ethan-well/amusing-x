package login

import (
	ganymedeservice "amusingx.fit/amusingx/protos/ganymede/service/ganymede/proto"
	panguservice "amusingx.fit/amusingx/protos/pangu/service/pangu/proto"
	"amusingx.fit/amusingx/rpcclient/ganymede"
	"context"
	"github.com/ItsWewin/superfactory/aerror"
)

func HandlerLogout(ctx context.Context, sessionID string) (*panguservice.LogoutResponse, aerror.Error) {
	resp, err := ganymede.Client.LogOut(ctx, &ganymedeservice.LogoutRequest{SessionID: sessionID})
	if err != nil {
		return &panguservice.LogoutResponse{Succeed: false}, aerror.NewErrorf(err, aerror.Code.SUnexpectedErr, "logout failed")
	}

	return &panguservice.LogoutResponse{Succeed: resp.Logout}, nil
}
