package verificationcontrolmodel

import (
	"amusingx.fit/amusingx/mysqlstruct/amusingriskcontrol"
	amusingriskcontrol2 "amusingx.fit/amusingx/services/callisto/mysql/amusingriskcontrol"
	"context"
	"github.com/ItsWewin/superfactory/xerror"
	"strconv"
)

// user id
// id
// verification
//
type VerificationRiskControlModel struct {
	Phone                  string                                   `json:"phone"`
	VerificationCodeRickDB *amusingriskcontrol.VerificationCodeRisk `json:"verification_code_rick_db"`
	AccountRiskStrategyDB  *amusingriskcontrol.AccountRiskStrategy  `json:"account_risk_strategy_db"`
}

func NewModel() *VerificationRiskControlModel {
	return &VerificationRiskControlModel{}
}

func (m *VerificationRiskControlModel) SetPhoneNumber(phone string) *xerror.Error {
	if len(phone) == 0 {
		return xerror.NewError(nil, xerror.Code.CParamsError, "Phone is invalid. ")
	}
	m.Phone = phone

	return nil
}

func (m *VerificationRiskControlModel) SetAccountRiskStrategyDB(ctx context.Context) *xerror.Error {
	strategy, err := amusingriskcontrol2.QueryAccountRiskStrategy(ctx, "verification_code")
	if err != nil {
		return xerror.NewError(err, err.Code, err.Message)
	}

	m.AccountRiskStrategyDB = strategy

	return nil
}

func (m *VerificationRiskControlModel) SetVerificationCodeRickDB(ctx context.Context) *xerror.Error {
	risk, err := amusingriskcontrol2.QueryVerificationCodeRiskByPhone(ctx, m.Phone)
	if err != nil {
		return xerror.NewError(err, xerror.Code.SSqlExecuteErr, "Sql execute failed. ")
	}

	m.VerificationCodeRickDB = risk

	return nil
}

func (m *VerificationRiskControlModel) VerificationUsedTimeAdd(ctx context.Context) *xerror.Error {
	if m.VerificationCodeRickDB == nil {

		value, err := strconv.ParseInt(m.AccountRiskStrategyDB.StrategyValue, 10, 64)
		if err != nil {
			return xerror.NewError(err, xerror.Code.BUnexpectedData,
				"m.AccountRiskStrategyDB.StrategyValue conv to int64 failed")
		}

		m.VerificationCodeRickDB = &amusingriskcontrol.VerificationCodeRisk{
			Phone:     m.Phone,
			UsedCount: 0,
			MaxCount:  value,
		}
	}

	m.VerificationCodeRickDB.UsedCount = m.VerificationCodeRickDB.UsedCount + 1

	err := amusingriskcontrol2.VerificationCodeRiskInsertOrUpdate(ctx, m.VerificationCodeRickDB)
	if err != nil {
		return xerror.NewError(err, err.Code, "VerificationUsedTimeAdd failed")
	}

	return err
}

func (m *VerificationRiskControlModel) CanUseVerificationCode() (bool, *xerror.Error) {
	if m.VerificationCodeRickDB == nil {
		return true, nil
	}

	return m.VerificationCodeRickDB.MaxCount > m.VerificationCodeRickDB.UsedCount, nil
}
