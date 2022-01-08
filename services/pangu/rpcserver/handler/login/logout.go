package login

import (
	ganymedeservice "amusingx.fit/amusingx/protos/ganymede/service/ganymede/proto"
	panguservice "amusingx.fit/amusingx/protos/pangu/service/pangu/proto"
	"amusingx.fit/amusingx/rpcclient/ganymede"
	"context"
	"github.com/ItsWewin/superfactory/aerror"
)

func HandlerLogout(ctx context.Context, req *panguservice.LogoutRequest) (*panguservice.LogoutResponse, aerror.Error) {
	resp, err := ganymede.Client.LogOut(ctx, &ganymedeservice.LogoutRequest{SessionID: req.SessionID})
	if err != nil {
		return &panguservice.LogoutResponse{Logout: false}, aerror.NewErrorf(err, aerror.Code.SUnexpectedErr, "logout failed")
	}

	return &panguservice.LogoutResponse{Logout: resp.Logout}, nil
}
