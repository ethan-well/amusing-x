package userservice

import (
	"amusingx.fit/amusingx/regexp"
	"github.com/ItsWewin/superfactory/verificationcode/randomcode"
	"github.com/ItsWewin/superfactory/xerror"
	"strings"
)

func (x *JoinRequest) Valid() *xerror.Error {
	x.Nickname = strings.TrimSpace(x.Nickname)
	x.Password = strings.TrimSpace(x.Password)
	x.AreaCode = strings.TrimSpace(x.AreaCode)
	x.Phone = strings.TrimSpace(x.Phone)
	x.VerificationCode = strings.TrimSpace(x.VerificationCode)

	if err := regexp.NicknameValid(x.Nickname); err != nil {
		return err
	}

	if err := regexp.PasswordValid(x.Password); err != nil {
		return err
	}

	if err := regexp.VerificationCodeValid(x.VerificationCode); err != nil {
		return err
	}

	codeStore := randomcode.RandomCodeStoreInit()
	if !codeStore.Check(x.VerificationCode) {
		return xerror.NewError(nil, xerror.Code.CParamsError, "验证码错误")
	}

	return nil
}
