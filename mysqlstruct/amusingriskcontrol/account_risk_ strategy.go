package amusingriskcontrol

type AccountRiskStrategy struct {
	ID                int64  `db:"id"`
	Strategy          string `db:"strategy"`
	StrategyValue     string `db:"strategy_value"`
	StrategyType      string `db:"strategy_type"`
	StrategyValueType string `db:"strategy_value_type"`
}
