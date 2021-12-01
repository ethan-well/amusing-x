package rpcclient

import (
	"amusingx.fit/amusingx/protos/callisto/service"
	"amusingx.fit/amusingx/services/ganymede/conf"
	"amusingx.fit/amusingx/services/ganymede/rpcclient/amusingxriskrpcserver"
	"context"
	"github.com/ItsWewin/superfactory/logger"
	"testing"
)

func TestInitRPCClient(t *testing.T) {
	conf.Conf.RiskServiceRPCAddress = "localhost:11002"
	xErr := InitRPCClient()
	if xErr == nil {
		t.Fatalf("some error: %s", xErr)
	}

	req := &riskservice.LoginRiskRequest{
		UserId:       0,
		StrategyType: "",
		Phone:        "",
		Action:       "",
	}

	result, err := amusingxriskrpcserver.RiskServerRPCClient.RiskServerRPCClient.LoginRiskControl(context.Background(), req)
	if err != nil {
		t.Fatalf("some error: %s", err)
	}

	t.Logf("result: %s", logger.ToJson(result))
}
