package rpcclient

import (
	"amusingx.fit/amusingx/protos/callisto/service"
	"amusingx.fit/amusingx/protos/ganymede/service/ganymede/proto"
	"amusingx.fit/amusingx/services/europa/conf"
	"amusingx.fit/amusingx/services/europa/rpcclient/callisto"
	"amusingx.fit/amusingx/services/europa/rpcclient/ganymede"
	"context"
	"errors"
	"github.com/ItsWewin/superfactory/aerror"
	"github.com/ItsWewin/superfactory/logger"
	"testing"
)

func TestInitCallistoRPCClient(t *testing.T) {
	conf.Mock()

	xErr := InitCallistoServerRPCClient()
	if xErr == nil {
		t.Fatalf("some error: %s", xErr)
	}

	req := &riskservice.LoginRiskRequest{
		UserId:       0,
		StrategyType: "",
		Phone:        "",
		Action:       "",
	}

	result, err := callisto.RPCClient.Client.LoginRiskControl(context.Background(), req)
	if err != nil {
		t.Fatalf("some error: %s", err)
	}

	t.Logf("result: %s", logger.ToJson(result))
}

func TestInitUserServerRPCClient(t *testing.T) {
	conf.Mock()

	xErr := InitGanymedeRPCClient()
	if xErr != nil {
		t.Fatalf("some error: %s", xErr)
	}

	req := &ganymedeservice.BlankParams{}

	result, err := ganymede.RPCClient.Client.Pong(context.Background(), req)
	if err != nil {
		t.Fatalf("some error: %s", err)
	}

	t.Logf("result: %s", logger.ToJson(result.ServerName))

	regex, err := ganymede.RPCClient.Client.Regexps(context.Background(), req)
	if err != nil {
		t.Fatalf("some error: %s", err)
	}

	t.Logf("regex: %s", regex)

	var loginReq = &ganymedeservice.LoginRequest{
		Type:             0,
		AreaCode:         "",
		Password:         "",
		VerificationCode: "",
		Phone:            "",
	}
	response, err := ganymede.RPCClient.Client.Login(context.Background(), loginReq)
	if err != nil {
		t.Logf("err: %s", err)
	}
	t.Logf("response: %s", logger.ToJson(response))

	joinReq := &ganymedeservice.JoinRequest{
		Nickname:         "wei,wei",
		Password:         "wewin123",
		AreaCode:         "87",
		Phone:            "18710565589",
		VerificationCode: "123",
	}
	joinResp, err := ganymede.RPCClient.Client.Join(context.Background(), joinReq)
	if err != nil {
		t.Logf("err: %s", err)
	}
	t.Logf("joinResp: %s", joinResp)

	code, err := ganymede.RPCClient.Client.CountryCodes(context.Background(), req)
	if err != nil {
		t.Logf("CountryCodes err: %s", err)
	}

	t.Logf("code: %s", code)

	vcReq := &ganymedeservice.VerificationCodeRequest{
		Phone:    "",
		Action:   "",
		AreaCode: "",
	}

	vsResult, err := ganymede.RPCClient.Client.GetVerificationCode(context.Background(), vcReq)
	if err != nil {
		t.Logf("GetVerificationCode err: %s", err)
	}
	t.Logf("vsResult: %s", vsResult)

	vccReq := &ganymedeservice.VerificationCodeRequest{
		Phone:    "",
		Action:   "",
		AreaCode: "",
	}
	vccResult, err := ganymede.RPCClient.Client.GetVerificationCode(context.Background(), vccReq)
	if err != nil {
		t.Logf("GetVerificationCode err: %s", err)
	}
	t.Logf("vccResult: %s", vccResult)

	rspReq := &ganymedeservice.ResetPasswordRequest{
		Password:         "",
		AreaCode:         "",
		Phone:            "",
		VerificationCode: "",
	}
	rspResult, err := ganymede.RPCClient.Client.ResetPassword(context.Background(), rspReq)
	if err != nil {
		t.Logf("ResetPassword err: %s", err)
	}
	t.Logf("rspResult: %s", rspResult)
}

func TestSomeError(t *testing.T) {
	err := _error()
	err = _aerror()
	t.Log(err.(aerror.Error).Code())
	t.Log(err.(aerror.Error).Message())
}

func _error() error {
	return errors.New("text")
}

func _aerror() aerror.Error {
	return aerror.NewErrorf(nil, "code", "message")
}
