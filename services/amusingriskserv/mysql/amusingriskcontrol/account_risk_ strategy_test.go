package amusingriskcontrol

import (
	"amusingx.fit/amusingx/services/amusingriskserv/conf"
	"context"
	"github.com/ItsWewin/superfactory/logger"
	"testing"
)

func TestQueryAccountRiskStrategy(t *testing.T) {
	conf.Mock()
	Mock()
	ctx := context.Background()
	strategyType := "verification_code"
	strategy, err := QueryAccountRiskStrategy(ctx, strategyType)
	if err != nil {
		t.Fatalf("Some err: %s", err)
	}

	t.Logf("strategy: %s", logger.ToJson(strategy))
}
