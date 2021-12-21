package europa

import (
	"amusingx.fit/amusingx/regexp"
	"github.com/ItsWewin/superfactory/aerror"
	"github.com/ItsWewin/superfactory/verificationcode/randomcode"
	"strings"
)

type JoinRequest struct {
	Nickname         string `json:"nickname"`
	Password         string `json:"password,omitempty"`
	AreaCode         string `json:"area_code"`
	Phone            string `json:"phone"`
	VerificationCode string `json:"verification_code"`
}

type JoinResponse struct {
	Nickname string `json:"nickname"`
	AreaCode string `json:"area_code"`
	Phone    string `json:"phone"`
}

func (r *JoinRequest) Valid() aerror.Error {
	r.Nickname = strings.TrimSpace(r.Nickname)
	r.Password = strings.TrimSpace(r.Password)
	r.AreaCode = strings.TrimSpace(r.AreaCode)
	r.Phone = strings.TrimSpace(r.Phone)
	r.VerificationCode = strings.TrimSpace(r.VerificationCode)

	if err := regexp.NicknameValid(r.Nickname); err != nil {
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
		return aerror.NewError(nil, aerror.Code.CParamsError, "验证码错误")
	}

	return nil
}
