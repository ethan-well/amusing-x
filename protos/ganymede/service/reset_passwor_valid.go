package ganymedeservice

import (
	"amusingx.fit/amusingx/regexp"
	"github.com/ItsWewin/superfactory/aerror"
	"github.com/ItsWewin/superfactory/verificationcode/randomcode"
)

func (x *ResetPasswordRequest) Valid() aerror.Error {
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
		return aerror.NewError(nil, aerror.Code.CParamsError, "验证码错误")
	}

	return nil
}
