package amusingriskcontrol

import (
	"amusingx.fit/amusingx/mysqlstruct/amusingriskcontrol"
	"context"
	"database/sql"
	"github.com/ItsWewin/superfactory/xerror"
)

func QueryAccountRiskStrategy(ctx context.Context, strategyType string) (*amusingriskcontrol.AccountRiskStrategy, *xerror.Error) {
	query := `
		SELECT id, strategy, strategy_type, strategy_value, strategy_value_type 
		FROM account_risk_strategy
		WHERE strategy_type = ?
	`
	var accountRiskStrategy amusingriskcontrol.AccountRiskStrategy
	err := AmusingRiskDB.GetContext(ctx, &accountRiskStrategy, query, strategyType)
	switch {
	case err == sql.ErrNoRows:
		return nil, nil
	case err != nil:
		return nil, xerror.NewError(err, xerror.Code.SSqlExecuteErr, "Sql execute failed. ")
	}

	return &accountRiskStrategy, nil
}
