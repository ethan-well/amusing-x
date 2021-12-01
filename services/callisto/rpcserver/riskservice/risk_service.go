package riskservice

import (
	"amusingx.fit/amusingx/protos/callisto/service"
	"amusingx.fit/amusingx/services/callisto/rpcserver/riskservice/controller/verificationcoderisk"
	"context"
	"github.com/ItsWewin/superfactory/xerror"
)

type LoginRiskService struct {
	riskservice.UnimplementedRiskServiceServer
}

func (s *LoginRiskService) LoginRiskControl(ctx context.Context, req *riskservice.LoginRiskRequest) (*riskservice.LoginRiskReply, error) {
	if req == nil {
		return nil, xerror.NewError(nil, xerror.Code.CUnexpectRequestDate, "req is nil")
	}

	err := verificationcoderisk.LoginRiskControl(ctx, req)
	if err != nil {
		return nil, xerror.NewError(err, err.Code, err.Message)
	}

	return &riskservice.LoginRiskReply{Result: true}, nil
}

func (s *LoginRiskService) Pong(ctx context.Context, params *riskservice.BlankParams) (*riskservice.PongReply, error) {
	return &riskservice.PongReply{Result: "Risk service. "}, nil
}
