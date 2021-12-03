package rpcclient

import (
	"amusingx.fit/amusingx/protos/callisto/service"
	"amusingx.fit/amusingx/protos/ganymede/service"
	"amusingx.fit/amusingx/services/europa/conf"
	"amusingx.fit/amusingx/services/europa/rpcclient/riskrpcserver"
	"amusingx.fit/amusingx/services/europa/rpcclient/userrpcserver"
	"context"
	"github.com/ItsWewin/superfactory/logger"
	"testing"
)

func TestInitRiskServerRPCClient(t *testing.T) {
	conf.ConfIns.RiskServiceRPCAddress = "localhost:11002"
	xErr := InitRiskServerRPCClient()
	if xErr == nil {
		t.Fatalf("some error: %s", xErr)
	}

	req := &riskservice.LoginRiskRequest{
		UserId:       0,
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
	conf.ConfIns.UserServiceRPCAddress = "localhost:11001"
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

	regex, err := userrpcserver.RPCClient.Client.Regexps(context.Background(), req)
	if err != nil {
		t.Fatalf("some error: %s", err)
	}

	t.Logf("regex: %s", regex)

	var loginReq = &userservice.LoginRequest{
		Type:             0,
		AreaCode:         "",
		Password:         "",
		VerificationCode: "",
		Phone:            "",
	}
	response, err := userrpcserver.RPCClient.Client.Login(context.Background(), loginReq)
	if err != nil {
		t.Logf("err: %s", err)
	}
	t.Logf("response: %s", logger.ToJson(response))

	joinReq := &userservice.JoinRequest{
		Nickname:         "wei,wei",
		Password:         "wewin123",
		AreaCode:         "87",
		Phone:            "18710565589",
		VerificationCode: "123",
	}
	joinResp, err := userrpcserver.RPCClient.Client.Join(context.Background(), joinReq)
	if err != nil {
		t.Logf("err: %s", err)
	}
	t.Logf("joinResp: %s", joinResp)

	code, err := userrpcserver.RPCClient.Client.CountryCodes(context.Background(), req)
	if err != nil {
		t.Logf("CountryCodes err: %s", err)
	}

	t.Logf("code: %s", code)

	vcReq := &userservice.VerificationCodeRequest{
		Phone:    "",
		Action:   "",
		AreaCode: "",
	}

	vsResult, err := userrpcserver.RPCClient.Client.GetVerificationCode(context.Background(), vcReq)
	if err != nil {
		t.Logf("GetVerificationCode err: %s", err)
	}
	t.Logf("vsResult: %s", vsResult)

	vccReq := &userservice.VerificationCodeRequest{
		Phone:    "",
		Action:   "",
		AreaCode: "",
	}
	vccResult, err := userrpcserver.RPCClient.Client.GetVerificationCode(context.Background(), vccReq)
	if err != nil {
		t.Logf("GetVerificationCode err: %s", err)
	}
	t.Logf("vccResult: %s", vccResult)

	rspReq := &userservice.ResetPasswordRequest{
		Password:         "",
		AreaCode:         "",
		Phone:            "",
		VerificationCode: "",
	}
	rspResult, err := userrpcserver.RPCClient.Client.ResetPassword(context.Background(), rspReq)
	if err != nil {
		t.Logf("ResetPassword err: %s", err)
	}
	t.Logf("rspResult: %s", rspResult)
}
