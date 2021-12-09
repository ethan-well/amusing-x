package ganymedeservice

import (
	"amusingx.fit/amusingx/regexp"
	"github.com/ItsWewin/superfactory/verificationcode/randomcode"
	"github.com/ItsWewin/superfactory/xerror"
)

func (x *ResetPasswordRequest) Valid() *xerror.Error {
	if err := regexp.PhoneNumberValid(x.Phone); err != nil {
		return err
	}

	if err := regexp.AreaCodeValid(x.AreaCode); err != nil {
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
