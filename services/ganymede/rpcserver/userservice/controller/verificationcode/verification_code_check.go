package verificationcode

import (
	"amusingx.fit/amusingx/protos/ganymede/userservice"
	"context"
	"github.com/ItsWewin/superfactory/logger"
	"github.com/ItsWewin/superfactory/verificationcode/randomcode"
	"github.com/ItsWewin/superfactory/xerror"
)

func HandlerVerificationCheck(ctx context.Context, req *userservice.VerificationCodeCheckRequest) (*userservice.VerificationCodeCheckResponse, *xerror.Error) {
	codeStore := randomcode.RandomCodeStoreInit()

	logger.Infof("code: %s", req.Code)

	if codeStore.Check(req.Code) {
		return &userservice.VerificationCodeCheckResponse{Result: "验证成功"}, nil
	} else {
		return &userservice.VerificationCodeCheckResponse{Result: "验证失败"}, nil
	}
}
