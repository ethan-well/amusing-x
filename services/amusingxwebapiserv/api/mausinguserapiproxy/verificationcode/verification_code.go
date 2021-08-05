package verificationcode

import (
	"github.com/ItsWewin/superfactory/httputil/rest"
	"github.com/ItsWewin/superfactory/logger"
	"github.com/ItsWewin/superfactory/verificationcode/randomcode"
	"github.com/ItsWewin/superfactory/xerror"
	"net/http"
)

func HandlerVerificationCode(w http.ResponseWriter, r *http.Request) {
	//ctx := context.Background()
	codeStore := randomcode.RandomCodeStoreInit()
	randomCode, err := codeStore.Generate()
	if err != nil {
		logger.Errorf("get and valid params failed, err: %s", err.Error())

		rest.FailJsonResponse(w, xerror.Code.SUnexpectedErr, xerror.Message.ParamsError)
		return
	}

	rest.SucceedJsonResponse(w, map[string]string{"code": randomCode.GetCode()})
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
