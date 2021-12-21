package ganymedeservice

import (
	"amusingx.fit/amusingx/regexp"
	"github.com/ItsWewin/superfactory/aerror"
)

func (req *LoginRequest) LoginByPassword() bool {
	return req.Type == 0
}

func (req *LoginRequest) LoginByVerificationCode() bool {
	return req.Type == 1
}

func (req *LoginRequest) Valid() aerror.Error {
	if req.Type != 0 && req.Type != 1 {
		return aerror.NewError(nil, aerror.Code.CParamsError, "Type is invalid. ")
	}

	if err := regexp.PhoneNumberValid(req.Phone); err != nil {
		return err
	}

	if err := regexp.AreaCodeValid(req.AreaCode); err != nil {
		return err
	}

	if req.Type == 0 {
		if err := regexp.PasswordValid(req.Password); err != nil {
			return err
		}
	} else {
		if err := regexp.VerificationCodeValid(req.VerificationCode); err != nil {
			return err
		}
	}

	return nil
}
