package amusinguserserv

import (
	"amusingx.fit/amusingx/services/amusinguserserv/regexp"
	"github.com/ItsWewin/superfactory/xerror"
	"strings"
)

type LoginRequest struct {
	Type     int    `json:"type"`
	AreaCode string `json:"area_code"`
	Phone    string `json:"phone"`
	Password string `json:"password"`
}

func (r *LoginRequest) Valid() *xerror.Error {
	r.Phone = strings.TrimSpace(r.Phone)
	r.AreaCode = strings.TrimSpace(r.AreaCode)
	r.Password = strings.TrimSpace(r.Password)

	if r.Type != 0 && r.Type != 1 {
		return xerror.NewError(nil, xerror.Code.CParamsError, "Type is invalid. ")
	}

	if err := regexp.PhoneNumberValid(r.Phone); err != nil {
		return err
	}

	if err := regexp.AreaCodeValid(r.AreaCode); err != nil {
		return err
	}

	if err := regexp.PasswordValid(r.Password); err != nil {
		return err
	}

	return nil
}
