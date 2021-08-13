package rpcclient

import (
	"amusingx.fit/amusingx/protos/amusingriskcontrolserv/loginrisk/loginrisk"
	"amusingx.fit/amusingx/protos/amusingxuserserv/userservice"
	"amusingx.fit/amusingx/services/amusingwebapiserv/conf"
	"amusingx.fit/amusingx/services/amusingwebapiserv/rpcclient/riskrpcserver"
	"amusingx.fit/amusingx/services/amusingwebapiserv/rpcclient/userrpcserver"
	"context"
	"github.com/ItsWewin/superfactory/logger"
	"testing"
)

func TestInitRiskServerRPCClient(t *testing.T) {
	conf.Conf.RiskServiceRPCAddress = "localhost:11002"
	xErr := InitRiskServerRPCClient()
	if xErr == nil {
		t.Fatalf("some error: %s", xErr)
	}

	req := &loginrisk.LoginRiskRequest{
		UserID:       0,
		StrategyType: "",
		Phone:        "",
		Action:       "",
	}

	result, err := riskrpcserver.RPCClient.Client.LoginRiskControl(context.Background(), req)
	if err != nil {
		t.Fatalf("some error: %s", err)
	}

	t.Logf("result: %s", logger.ToJson(result))
}

func TestInitUserServerRPCClient(t *testing.T) {
	conf.Conf.UserServiceRPCAddress = "localhost:11001"
	xErr := InitUserServerRPCClient()
	if xErr != nil {
		t.Fatalf("some error: %s", xErr)
	}

	req := &userservice.BlankParams{}

	result, err := userrpcserver.RPCClient.Client.Pong(context.Background(), req)
	if err != nil {
		t.Fatalf("some error: %s", err)
	}

	t.Logf("result: %s", logger.ToJson(result.ServerName))
}
