package loginriskserver

import (
	"amusingx.fit/amusingx/protos/amusingriskcontrolserv/loginrisk/loginrisk"
	"amusingx.fit/amusingx/services/amusingriskserv/app/verificationcoderisk"
	"context"
	"github.com/ItsWewin/superfactory/xerror"
)

type LoginRiskServer struct {
	loginrisk.UnimplementedLoginRiskServer
}

func (s *LoginRiskServer) LoginRiskControl(ctx context.Context, req *loginrisk.LoginRiskRequest) (*loginrisk.LoginRiskReply, error) {
	if req == nil {
		return nil, xerror.NewError(nil, xerror.Code.CUnexpectRequestDate, "req is nil")
	}

	err := verificationcoderisk.LoginRiskControl(ctx, req)
	if err != nil {
		return nil, xerror.NewError(err, err.Code, err.Message)
	}

	return &loginrisk.LoginRiskReply{Result: true}, nil
}
