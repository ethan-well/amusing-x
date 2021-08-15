package userservice

import (
	"amusingx.fit/amusingx/services/amusinguserserv/regexp"
	"github.com/ItsWewin/superfactory/xerror"
)

func (x *LoginRequest) LoginByPassword() bool {
	return x.Type == 0
}

func (x *LoginRequest) LoginByVerificationCode() bool {
	return x.Type == 1
}

func (x *LoginRequest) Valid() *xerror.Error {
	if x.Type != 0 && x.Type != 1 {
		return xerror.NewError(nil, xerror.Code.CParamsError, "Type is invalid. ")
	}

	if err := regexp.PhoneNumberValid(x.Phone); err != nil {
		return err
	}

	if err := regexp.AreaCodeValid(x.AreaCode); err != nil {
		return err
	}

	if x.Type == 0 {
		if err := regexp.PasswordValid(x.Password); err != nil {
			return err
		}
	} else {
		if err := regexp.VerificationCodeValid(x.VerificationCode); err != nil {
			return err
		}
	}

	return nil
}
