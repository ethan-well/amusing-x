package callistoservice

import (
	"amusingx.fit/amusingx/protos/callisto/service"
	"amusingx.fit/amusingx/services/risk/rpcserver/callistoservice/controller/verificationcoderisk"
	"context"
	"github.com/ItsWewin/superfactory/aerror"
)

type LoginRiskService struct {
	riskservice.UnimplementedRiskServiceServer
}

func (s *LoginRiskService) LoginRiskControl(ctx context.Context, req *riskservice.LoginRiskRequest) (*riskservice.LoginRiskReply, error) {
	if req == nil {
		return nil, aerror.NewError(nil, aerror.Code.CUnexpectRequestDate, "req is nil")
	}

	err := verificationcoderisk.LoginRiskControl(ctx, req)
	if err != nil {
		return nil, aerror.NewError(err, err.Code(), err.Message())
	}

	return &riskservice.LoginRiskReply{Result: true}, nil
}

func (s *LoginRiskService) Pong(ctx context.Context, params *riskservice.BlankParams) (*riskservice.PongReply, error) {
	return &riskservice.PongReply{Result: "Risk service. "}, nil
}
