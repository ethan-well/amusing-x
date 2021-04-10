package regex

import (
	"fmt"
	"github.com/ItsWewin/superfactory/xerror"
	"regexp"
	"unicode"
)

const (
	PhoneRegex    = `^\d{5,11}$`
	AreaRegex     = `^[0-9a-zA-Z]{3,10}$`
	NicknameRegex = `\w{5,20}`

	VerificationCodeRegex = `^[0-9a-zA-Z]{6}$`
)

func PhoneNumberValid(phone string) *xerror.Error {
	matched, err := regexp.MatchString(PhoneRegex, phone)
	if err != nil {
		return xerror.NewError(err, xerror.Code.SUnexpectedErr, "Unexpected error. ")
	}

	if !matched {
		msg := fmt.Sprintf("Phone number is invalid, expected regex: %s but get: %s", PhoneRegex, phone)
		xerror.NewError(nil, xerror.Code.CParamsError, msg)
	}

	return nil
}

func AreaCodeValid(areaCode string) *xerror.Error {
	matched, err := regexp.MatchString(AreaRegex, areaCode)
	if err != nil {
		return xerror.NewError(err, xerror.Code.SUnexpectedErr, "Unexpected error. ")
	}

	if !matched {
		msg := fmt.Sprintf("area code is invalid, expected regex: %s but code: %s", AreaRegex, areaCode)
		return xerror.NewError(err, xerror.Code.CParamsError, msg)
	}

	return nil
}

func NicknameValid(nickname string) *xerror.Error {
	matched, err := regexp.MatchString(NicknameRegex, nickname)
	if err != nil {
		return xerror.NewError(err, xerror.Code.SUnexpectedErr, "Unexpected error. ")
	}

	if !matched {
		return xerror.NewError(err, xerror.Code.CParamsError, "昵称由 5 到 20 位字母、数字、下划线或者汉子组成")
	}

	return nil
}

func VerificationCodeValid(code string) *xerror.Error {
	matched, err := regexp.MatchString(VerificationCodeRegex, code)
	if err != nil {
		return xerror.NewError(err, xerror.Code.SUnexpectedErr, "Unexpected error. ")
	}

	if !matched {
		return xerror.NewError(err, xerror.Code.CParamsError, "验证码不正确")
	}

	return nil
}

func PasswordValid(str string) *xerror.Error {
	var (
		isUpper   = false
		isLower   = false
		isNumber  = false
		isSpecial = false
	)

	if len(str) < 6 || len(str) > 16 {
		return xerror.NewError(nil, xerror.Code.CParamsError,
			"The password must contain uppercase and lowercase letters, numbers or punctuation, and must be 6-16 digits long. ")
	}

	for _, s := range str {
		switch {
		case unicode.IsUpper(s):
			isUpper = true
		case unicode.IsLower(s):
			isLower = true
		case unicode.IsNumber(s):
			isNumber = true
		case unicode.IsPunct(s) || unicode.IsSymbol(s):
			isSpecial = true
		default:
		}
	}

	if (isUpper && isLower) && (isNumber || isSpecial) {
		return nil
	}

	return xerror.NewError(nil, xerror.Code.CParamsError,
		"The password must contain uppercase and lowercase letters, numbers or punctuation, and must be 6-16 digits long. ")
}
