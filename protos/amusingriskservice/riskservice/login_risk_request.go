package riskservice

import (
	"github.com/ItsWewin/superfactory/xerror"
)

var StrategyType = []string{"verification_code"}

func (x *LoginRiskRequest) ValidStrategyType() bool {
	for _, st := range StrategyType {
		if st == x.StrategyType {
			return true
		}
	}

	return false
}

// 策略控制值增加
func (x *LoginRiskRequest) IsValueAddAction() bool {
	return x.Action == "value_add"
}

// 策略控制值验证
func (x *LoginRiskRequest) IsValueVerifyAction() bool {
	return x.Action == "value_verify"
}

func (x *LoginRiskRequest) IsVerificationCodeStrategy() bool {
	return x.StrategyType == "verification_code"
}

func (x *LoginRiskRequest) VerificationCodeStrategyRequestValid() *xerror.Error {
	if len(x.Phone) == 0 {
		return xerror.NewError(nil, xerror.Code.CParamsError, "params 'phone' is expected")
	}

	if !x.IsValueAddAction() && !x.IsValueVerifyAction() {
		return xerror.NewError(nil, xerror.Code.CParamsError, "params 'action' is expected")
	}

	return nil
}

func (x *LoginRiskRequest) Valid() *xerror.Error {
	if len(x.StrategyType) == 0 {
		return xerror.NewError(nil, xerror.Code.CUnexpectRequestDate, "strategy_type is expected")
	}

	if !x.ValidStrategyType() {
		return xerror.NewError(nil, xerror.Code.CParamsError, "strategy_type is invalid")
	}

	if x.IsVerificationCodeStrategy() {
		return x.VerificationCodeStrategyRequestValid()
	}

	return nil
}
