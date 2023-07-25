package verificationcode

import (
	service "amusingx.fit/amusingx/protos/user/service/user/proto"
	"context"
	"github.com/ItsWewin/superfactory/aerror"
	"github.com/ItsWewin/superfactory/logger"
	"github.com/ItsWewin/superfactory/verificationcode/randomcode"
)

func HandlerVerificationCheck(ctx context.Context, req *service.VerificationCodeCheckRequest) (*service.VerificationCodeCheckResponse, aerror.Error) {
	codeStore := randomcode.RandomCodeStoreInit()

	logger.Infof("code: %s", req.Code)

	if codeStore.Check(req.Code) {
		return &service.VerificationCodeCheckResponse{Result: "验证成功"}, nil
	} else {
		return &service.VerificationCodeCheckResponse{Result: "验证失败"}, nil
	}
}
