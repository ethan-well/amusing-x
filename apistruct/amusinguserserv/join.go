package amusinguserserv

import (
	"amusingx.fit/amusingx/services/amusinguserserv/regex"
	"github.com/ItsWewin/superfactory/xerror"
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

func (r *JoinRequest) Valid() *xerror.Error {
	r.Nickname = strings.TrimSpace(r.Nickname)
	r.Password = strings.TrimSpace(r.Password)
	r.AreaCode = strings.TrimSpace(r.AreaCode)
	r.Phone = strings.TrimSpace(r.Phone)
	r.VerificationCode = strings.TrimSpace(r.VerificationCode)

	if err := regex.NicknameValid(r.Nickname); err != nil {
		return err
	}

	if err := regex.PasswordValid(r.Password); err != nil {
		return err
	}

	if err := regex.VerificationCodeValid(r.VerificationCode); err != nil {
		return err
	}

	return nil
}
