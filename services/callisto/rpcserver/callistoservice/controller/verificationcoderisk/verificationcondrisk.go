package verificationcoderisk

import (
	"amusingx.fit/amusingx/protos/callisto/service"
	"amusingx.fit/amusingx/services/callisto/rpcserver/callistoservice/model/verificationcontrolmodel"
	"context"
	"github.com/ItsWewin/superfactory/aerror"
)

func LoginRiskControl(ctx context.Context, req *riskservice.LoginRiskRequest) aerror.Error {
	if req == nil {
		return aerror.NewError(nil, aerror.Code.CParamsError, "request is nil")
	}

	switch req.StrategyType {
	case `verification_code`:
		return verificationCodeRiskControl(ctx, req)
	}

	return aerror.NewError(nil, aerror.Code.CUnexpectRequestDate, "strategy_type is invalid")
}

func verificationCodeRiskControl(ctx context.Context, req *riskservice.LoginRiskRequest) aerror.Error {
	verificationCodeModel := verificationcontrolmodel.NewModel()

	err := verificationCodeModel.SetPhoneNumber(req.Phone)
	if err != nil {
		return err
	}

	err = verificationCodeModel.SetAccountRiskStrategyDB(ctx)
	if err != nil {
		return err
	}

	err = verificationCodeModel.SetVerificationCodeRickDB(ctx)
	if err != nil {
		return err
	}

	switch {
	case req.IsValueVerifyAction():
		canUsed, err := verificationCodeModel.CanUseVerificationCode()
		if err != nil {
			return err
		}

		if !canUsed {
			return aerror.NewError(nil, aerror.Code.BDataIsNotAllow, "不能使用验证码服务，可能是超过次数限制")
		}

		return nil
	case req.IsValueAddAction():
		return verificationCodeModel.VerificationUsedTimeAdd(ctx)
	default:
		return aerror.NewError(nil, aerror.Code.CUnexpectRequestDate, "非法封 action")
	}
}
