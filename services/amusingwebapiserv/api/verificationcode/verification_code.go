package verificationcode

import (
	"amusingx.fit/amusingx/apistruct/amusinguserserv"
	"amusingx.fit/amusingx/protos/amusingriskcontrolserv/loginrisk/loginrisk"
	"amusingx.fit/amusingx/services/amusingwebapiserv/model"
	"amusingx.fit/amusingx/services/amusingwebapiserv/rpcclient/riskrpcserver"
	"context"
	"github.com/ItsWewin/superfactory/httputil/rest"
	"github.com/ItsWewin/superfactory/logger"
	"github.com/ItsWewin/superfactory/verificationcode/randomcode"
	"github.com/ItsWewin/superfactory/xerror"
	"net/http"
)

func HandlerVerificationCode(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	req, err := getAndValidParams(ctx, r)
	if err != nil {
		rest.FailJsonResponse(w, err.Code, err.Message)
		return
	}

	err = riskControl(ctx, req.Phone)
	if err != nil {
		logger.Errorf("[HandlerVerificationCode] riskControl failed: %s", err)
		rest.FailJsonResponse(w, err.Code, err.Message)
		return
	}

	codeStore := randomcode.RandomCodeStoreInit()
	randomCode, err := codeStore.Generate()
	if err != nil {
		logger.Errorf("get and valid params failed, err: %s", err.Error())

		rest.FailJsonResponse(w, xerror.Code.SUnexpectedErr, xerror.Message.ParamsError)
		return
	}

	rest.SucceedJsonResponse(w, map[string]string{"code": randomCode.GetCode()})

	go riskControlValueVerifyAdd(ctx, req.Phone)
}

func getAndValidParams(ctx context.Context, r *http.Request) (*amusinguserserv.VerificationCodeRequest, *xerror.Error) {
	request := &amusinguserserv.VerificationCodeRequest{
		Phone:    r.FormValue("phone"),
		Action:   r.FormValue("action"),
		AreaCode: r.FormValue("area_code"),
	}

	logger.Infof("request: %s", logger.ToJson(request))

	if err := request.Valid(); err != nil {
		return nil, err
	}

	user := model.User{
		Phone:    request.Phone,
		AreaCode: request.AreaCode,
	}

	switch {
	case request.IsJoin():
		existed, err := user.ExistedWithPhone(ctx)
		if err != nil {
			return nil, xerror.NewError(err, err.Code, "getAndValidParams failed")
		}
		if existed {
			return nil, xerror.NewError(nil, xerror.Code.CParamsError, "phone is token")
		}

		return request, nil
	case request.IsLogin():
		existed, err := user.ExistedWithPhone(ctx)
		if err != nil {
			return nil, xerror.NewError(err, err.Code, "getAndValidParams failed")
		}
		if !existed {
			return nil, xerror.NewError(nil, xerror.Code.CParamsError, "phone number not join")
		}
	default:
		return nil, xerror.NewError(nil, xerror.Code.CParamsError, "'action' is invalid")
	}

	return request, nil
}

func riskControl(ctx context.Context, phone string) *xerror.Error {
	req := &loginrisk.LoginRiskRequest{
		StrategyType: "verification_code",
		Phone:        phone,
		Action:       "value_verify",
	}

	reply, err := riskrpcserver.RPCClient.RiskServerRPCClient.LoginRiskControl(ctx, req)
	if err != nil {
		return xerror.NewError(err, xerror.Code.BUnexpectedData, "riskControl request risk control failed")
	}

	if !reply.Result {
		return xerror.NewError(err, xerror.Code.BUnexpectedData, "不能使用验证码服务")
	}

	return nil
}

func riskControlValueVerifyAdd(ctx context.Context, phone string) {
	req := &loginrisk.LoginRiskRequest{
		StrategyType: "verification_code",
		Phone:        phone,
		Action:       "value_add",
	}

	reply, err := riskrpcserver.RPCClient.RiskServerRPCClient.LoginRiskControl(ctx, req)
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

func HandlerVerificationCheck(w http.ResponseWriter, r *http.Request) {
	codeStore := randomcode.RandomCodeStoreInit()
	code := r.FormValue("code")

	logger.Infof("code: %s", code)

	if codeStore.Check(code) {
		rest.SucceedJsonResponse(w, "check succeed")
		return
	} else {
		rest.FailJsonResponse(w, xerror.Code.CUnexpectRequestDate, "验证失败")
	}
}
