package amusinguserserv

import (
	"errors"
	"strings"
)

type JoinRequest struct {
	Nickname         string `json:"nickname"`
	Password         string `json:"password"`
	AreaCode         string `json:"area_code"`
	Phone            string `json:"phone"`
	VerificationCode string `json:"verification_code"`
}

func (r *JoinRequest) Valid() error {
	r.Nickname = strings.TrimSpace(r.Nickname)
	r.Password = strings.TrimSpace(r.Password)
	r.AreaCode = strings.TrimSpace(r.AreaCode)
	r.Phone = strings.TrimSpace(r.Phone)
	r.VerificationCode = strings.TrimSpace(r.VerificationCode)

	if len(r.Nickname) <= 3 {
		return errors.New("nickname is invalid")
	}

	if len(r.Password) <= 6 {
		return errors.New("password length should should >= 6")
	}

	if len(r.VerificationCode) != 6 {
		return errors.New("verification code is invalid")
	}

	return nil
}
