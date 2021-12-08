package verificationcode

import (
	"amusingx.fit/amusingx/protos/callisto/service"
	"amusingx.fit/amusingx/protos/ganymede/service"
	"amusingx.fit/amusingx/services/ganymede/rpcclient/amusingxriskrpcserver"
	"amusingx.fit/amusingx/services/ganymede/rpcserver/userservice/model"
	"context"
	"github.com/ItsWewin/superfactory/logger"
	"github.com/ItsWewin/superfactory/verificationcode/randomcode"
	"github.com/ItsWewin/superfactory/xerror"
)

func HandlerVerificationCode(ctx context.Context,
	req *ganymedeservice.VerificationCodeRequest) (*ganymedeservice.VerificationCodeResponse, *xerror.Error) {

	err := getAndValidParams(ctx, req)
	if err != nil {
		return nil, err
	}

	err = riskControl(ctx, req.Phone)
	if err != nil {
		return nil, err
	}

	logger.Infof("HandlerVerificationCode getAndValidParams: %s", logger.ToJson(req))

	codeStore := randomcode.RandomCodeStoreInit()
	randomCode, err := codeStore.Generate()
	if err != nil {
		return nil, err
	}

	go riskControlValueVerifyAdd(ctx, req.Phone)

	return &ganymedeservice.VerificationCodeResponse{Code: randomCode.GetCode()}, nil
}

func getAndValidParams(ctx context.Context, request *ganymedeservice.VerificationCodeRequest) *xerror.Error {
	logger.Infof("request: %s", logger.ToJson(request))

	if err := request.Valid(); err != nil {
		return err
	}

	user := model.UserComplex{
		Phone:    request.Phone,
		AreaCode: request.AreaCode,
	}

	switch {
	case request.IsJoin():
		existed, err := user.ExistedWithPhone(ctx)
		if err != nil {
			return xerror.NewError(err, err.Code, "getAndValidParams failed")
		}
		if existed {
			return xerror.NewError(nil, xerror.Code.CParamsError, "phone is token")
		}

		return nil
	case request.IsLogin():
		existed, err := user.ExistedWithPhone(ctx)
		if err != nil {
			return xerror.NewError(err, err.Code, "getAndValidParams failed")
		}
		if !existed {
			return xerror.NewError(nil, xerror.Code.CParamsError, "phone number not join")
		}
	default:
		return xerror.NewError(nil, xerror.Code.CParamsError, "'action' is invalid")
	}

	return nil
}

func riskControl(ctx context.Context, phone string) *xerror.Error {
	req := &riskservice.LoginRiskRequest{
		StrategyType: "verification_code",
		Phone:        phone,
		Action:       "value_verify",
	}

	reply, err := amusingxriskrpcserver.RiskServerRPCClient.RiskServerRPCClient.LoginRiskControl(ctx, req)
	if err != nil {
		return xerror.NewError(err, xerror.Code.BUnexpectedData, "riskControl request risk control failed")
	}

	if !reply.Result {
		return xerror.NewError(err, xerror.Code.BUnexpectedData, "不能使用验证码服务")
	}

	return nil
}

func riskControlValueVerifyAdd(ctx context.Context, phone string) {
	req := &riskservice.LoginRiskRequest{
		StrategyType: "verification_code",
		Phone:        phone,
		Action:       "value_add",
	}

	reply, err := amusingxriskrpcserver.RiskServerRPCClient.RiskServerRPCClient.LoginRiskControl(ctx, req)
	if err != nil {
		err := xerror.NewError(err, xerror.Code.BUnexpectedData, "request risk control failed")
		logger.Errorf("riskControlValueVerifyAdd failed: %s", err.Error())
		return
	}

	if !reply.Result {
		err := xerror.NewError(err, xerror.Code.BUnexpectedData, "不能使用验证码服务")
		logger.Errorf("riskControlValueVerifyAdd failed: %s", err.Error())
		return
	}

	logger.Infof("riskControlValueVerifyAdd succeed")
}
