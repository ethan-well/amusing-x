package verificationcode

import (
	"amusingx.fit/amusingx/protos/ganymede/service"
	"context"
	"github.com/ItsWewin/superfactory/aerror"
	"github.com/ItsWewin/superfactory/logger"
	"github.com/ItsWewin/superfactory/verificationcode/randomcode"
)

func HandlerVerificationCheck(ctx context.Context, req *ganymedeservice.VerificationCodeCheckRequest) (*ganymedeservice.VerificationCodeCheckResponse, aerror.Error) {
	codeStore := randomcode.RandomCodeStoreInit()

	logger.Infof("code: %s", req.Code)

	if codeStore.Check(req.Code) {
		return &ganymedeservice.VerificationCodeCheckResponse{Result: "验证成功"}, nil
	} else {
		return &ganymedeservice.VerificationCodeCheckResponse{Result: "验证失败"}, nil
	}
}
