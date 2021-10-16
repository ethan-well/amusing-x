package amusinguserserv

import (
	"amusingx.fit/amusingx/regexp"
	"github.com/ItsWewin/superfactory/verificationcode/randomcode"
	"github.com/ItsWewin/superfactory/xerror"
	"strings"
)

type ResetPasswordRequest struct {
	Password         string `json:"password,omitempty"`
	AreaCode         string `json:"area_code"`
	Phone            string `json:"phone"`
	VerificationCode string `json:"verification_code"`
}

func (r *ResetPasswordRequest) Valid() *xerror.Error {
	r.Phone = strings.TrimSpace(r.Phone)
	r.AreaCode = strings.TrimSpace(r.AreaCode)
	r.Password = strings.TrimSpace(r.Password)

	if err := regexp.PhoneNumberValid(r.Phone); err != nil {
		return err
	}

	if err := regexp.AreaCodeValid(r.AreaCode); err != nil {
		return err
	}

	if err := regexp.PasswordValid(r.Password); err != nil {
		return err
	}

	if err := regexp.VerificationCodeValid(r.VerificationCode); err != nil {
		return err
	}

	codeStore := randomcode.RandomCodeStoreInit()
	if !codeStore.Check(r.VerificationCode) {
		return xerror.NewError(nil, xerror.Code.CParamsError, "验证码错误")
	}

	return nil
}
