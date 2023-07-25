package verificationcode

import (
	"amusingx.fit/amusingx/apistruct/apiService"
	"amusingx.fit/amusingx/protos/callisto/service"
	"amusingx.fit/amusingx/services/webApi/model"
	"amusingx.fit/amusingx/services/webApi/rpcclient/callisto"
	"context"
	"github.com/ItsWewin/superfactory/aerror"
	"github.com/ItsWewin/superfactory/httputil/rest"
	"github.com/ItsWewin/superfactory/logger"
	"github.com/ItsWewin/superfactory/verificationcode/randomcode"
	"net/http"
)

func HandlerVerificationCode(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	req, err := getAndValidParams(ctx, r)
	if err != nil {
		rest.FailJsonResponse(w, err.Code(), err.Message())
		return
	}

	err = riskControl(ctx, req.Phone)
	if err != nil {
		logger.Errorf("[HandlerVerificationCode] riskControl failed: %s", err)
		rest.FailJsonResponse(w, err.Code(), err.Message())
		return
	}

	codeStore := randomcode.RandomCodeStoreInit()
	randomCode, err := codeStore.Generate()
	if err != nil {
		logger.Errorf("get and valid params failed, err: %s", err.Error())

		rest.FailJsonResponse(w, aerror.Code.SUnexpectedErr, aerror.Message.ParamsError)
		return
	}

	rest.SucceedJsonResponse(w, map[string]string{"code": randomCode.GetCode()})

	go riskControlValueVerifyAdd(ctx, req.Phone)
}

func getAndValidParams(ctx context.Context, r *http.Request) (*apiService.VerificationCodeRequest, aerror.Error) {
	request := &apiService.VerificationCodeRequest{
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
			return nil, aerror.NewError(err, err.Code(), "getAndValidParams failed")
		}
		if existed {
			return nil, aerror.NewError(nil, aerror.Code.CParamsErr, "phone is token")
		}

		return request, nil
	case request.IsLogin():
		existed, err := user.ExistedWithPhone(ctx)
		if err != nil {
			return nil, aerror.NewError(err, err.Code(), "getAndValidParams failed")
		}
		if !existed {
			return nil, aerror.NewError(nil, aerror.Code.CParamsErr, "phone number not join")
		}
	default:
		return nil, aerror.NewError(nil, aerror.Code.CParamsErr, "'action' is invalid")
	}

	return request, nil
}

func riskControl(ctx context.Context, phone string) aerror.Error {
	req := &riskservice.LoginRiskRequest{
		StrategyType: "verification_code",
		Phone:        phone,
		Action:       "value_verify",
	}

	reply, err := callisto.RPCClient.Client.LoginRiskControl(ctx, req)
	if err != nil {
		return aerror.NewError(err, aerror.Code.BUnexpectedData, "riskControl request risk control failed")
	}

	if !reply.Result {
		return aerror.NewError(err, aerror.Code.BUnexpectedData, "不能使用验证码服务")
	}

	return nil
}

func riskControlValueVerifyAdd(ctx context.Context, phone string) {
	req := &riskservice.LoginRiskRequest{
		StrategyType: "verification_code",
		Phone:        phone,
		Action:       "value_add",
	}

	reply, err := callisto.RPCClient.Client.LoginRiskControl(ctx, req)
	if err != nil {
		err := aerror.NewError(err, aerror.Code.BUnexpectedData, "request risk control failed")
		logger.Errorf("riskControlValueVerifyAdd failed: %s", err.Error())
		return
	}

	if !reply.Result {
		err := aerror.NewError(err, aerror.Code.BUnexpectedData, "不能使用验证码服务")
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
		rest.FailJsonResponse(w, aerror.Code.CUnexpectRequestDate, "验证失败")
	}
}
