package verificationcoderisk

import (
	"amusingx.fit/amusingx/protos/amusingriskservice/riskservice"
	"amusingx.fit/amusingx/services/amusingriskserv/model/verificationcontrolmodel"
	"context"
	"github.com/ItsWewin/superfactory/xerror"
)

func LoginRiskControl(ctx context.Context, req *riskservice.LoginRiskRequest) *xerror.Error {
	if req == nil {
		return xerror.NewError(nil, xerror.Code.CParamsError, "request is nil")
	}

	switch req.StrategyType {
	case `verification_code`:
		return verificationCodeRiskControl(ctx, req)
	}

	return xerror.NewError(nil, xerror.Code.CUnexpectRequestDate, "strategy_type is invalid")
}

func verificationCodeRiskControl(ctx context.Context, req *riskservice.LoginRiskRequest) *xerror.Error {
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
			return xerror.NewError(nil, xerror.Code.BDataIsNotAllow, "不能使用验证码服务，可能是超过次数限制")
		}

		return nil
	case req.IsValueAddAction():
		return verificationCodeModel.VerificationUsedTimeAdd(ctx)
	default:
		return xerror.NewError(nil, xerror.Code.CUnexpectRequestDate, "非法封 action")
	}
}
