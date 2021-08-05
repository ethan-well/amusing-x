package amusingriskcontrolserv

import (
	"github.com/ItsWewin/superfactory/xerror"
)

type LoginRiskControlRequest struct {
	UserId       int64  `json:"user_id"`
	StrategyType string `json:"strategy_type"`
	Phone        string `json:"phone"`
	Action       string `json:"action"`
}

var StrategyType = []string{"verification_code"}

func (r *LoginRiskControlRequest) ValidStrategyType() bool {
	for _, st := range StrategyType {
		if st == r.StrategyType {
			return true
		}
	}

	return false
}

// 策略控制值增加
func (r *LoginRiskControlRequest) IsValueAddAction() bool {
	return r.Action == "value_add"
}

// 策略控制值验证
func (r *LoginRiskControlRequest) IsValueVerifyAction() bool {
	return r.Action == "value_verify"
}

func (r *LoginRiskControlRequest) IsVerificationCodeStrategy() bool {
	return r.StrategyType == "verification_code"
}

func (r *LoginRiskControlRequest) VerificationCodeStrategyRequestValid() *xerror.Error {
	if len(r.Phone) == 0 {
		return xerror.NewError(nil, xerror.Code.CParamsError, "params 'phone' is expected")
	}

	if !r.IsValueAddAction() && !r.IsValueVerifyAction() {
		return xerror.NewError(nil, xerror.Code.CParamsError, "params 'action' is expected")
	}

	return nil
}

func (r *LoginRiskControlRequest) Valid() *xerror.Error {
	if len(r.StrategyType) == 0 {
		return xerror.NewError(nil, xerror.Code.CUnexpectRequestDate, "strategy_type is expected")
	}

	if !r.ValidStrategyType() {
		return xerror.NewError(nil, xerror.Code.CParamsError, "strategy_type is invalid")
	}

	if r.IsVerificationCodeStrategy() {
		return r.VerificationCodeStrategyRequestValid()
	}

	return nil
}
