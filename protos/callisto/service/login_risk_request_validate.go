package riskservice

import "github.com/ItsWewin/superfactory/aerror"

var StrategyType = []string{"verification_code"}

func (r *LoginRiskRequest) ValidStrategyType() bool {
	for _, st := range StrategyType {
		if st == r.StrategyType {
			return true
		}
	}

	return false
}

// 策略控制值增加
func (r *LoginRiskRequest) IsValueAddAction() bool {
	return r.Action == "value_add"
}

// 策略控制值验证
func (r *LoginRiskRequest) IsValueVerifyAction() bool {
	return r.Action == "value_verify"
}

func (r *LoginRiskRequest) IsVerificationCodeStrategy() bool {
	return r.StrategyType == "verification_code"
}

func (r *LoginRiskRequest) VerificationCodeStrategyRequestValid() aerror.Error {
	if len(r.Phone) == 0 {
		return aerror.NewError(nil, aerror.Code.CParamsError, "params 'phone' is expected")
	}

	if !r.IsValueAddAction() && !r.IsValueVerifyAction() {
		return aerror.NewError(nil, aerror.Code.CParamsError, "params 'action' is expected")
	}

	return nil
}

func (r *LoginRiskRequest) Valid() aerror.Error {
	if len(r.StrategyType) == 0 {
		return aerror.NewError(nil, aerror.Code.CUnexpectRequestDate, "strategy_type is expected")
	}

	if !r.ValidStrategyType() {
		return aerror.NewError(nil, aerror.Code.CParamsError, "strategy_type is invalid")
	}

	if r.IsVerificationCodeStrategy() {
		return r.VerificationCodeStrategyRequestValid()
	}

	return nil
}
