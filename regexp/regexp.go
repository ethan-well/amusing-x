package regexp

import (
	"fmt"
	"github.com/ItsWewin/superfactory/aerror"
	"regexp"
	"unicode"
)

const (
	PhoneRegex            = `^\d{5,11}$`
	AreaRegex             = `^[0-9a-zA-Z]{2,10}$`
	NicknameRegex         = `\w{5,20}`
	PasswordRegexpForWeb  = "(?=.*([a-zA-Z].*))(?=.*[0-9].*)[a-zA-Z0-9-*/+.~!@#$%^&*]{6,18}$"
	VerificationCodeRegex = `^[0-9a-zA-Z]{4,6}$`
)

type Regexp struct {
	ID    int64  `json:"id"`
	Name  string `json:"name"`
	Rules string `json:"rules"`
	Desc  string `json:"desc"`
}

var AllRegexps = []*Regexp{
	{
		ID:    1,
		Name:  "phone",
		Rules: PhoneRegex,
		Desc:  "手机号码是 5 - 11 位数字",
	},
	{
		ID:    2,
		Name:  "area_code",
		Rules: AreaRegex,
		Desc:  "区号为 3 - 10 位数组或者字母",
	},
	{
		ID:    3,
		Name:  "nickname",
		Rules: NicknameRegex,
		Desc:  "昵称为 5 - 20 位字母或者数字",
	},
	{
		ID:    4,
		Name:  "verification_code",
		Rules: VerificationCodeRegex,
		Desc:  "验证码为 4 - 6 位字母或者数字",
	},
	{
		ID:    5,
		Name:  "password",
		Rules: PasswordRegexpForWeb,
		Desc:  "6-16 位字母、数字、特殊字符(~!@#$%^&*)组成，至少包含一个数字和字母",
	},
}

func PhoneNumberValid(phone string) aerror.Error {
	matched, err := regexp.MatchString(PhoneRegex, phone)
	if err != nil {
		return aerror.NewError(err, aerror.Code.SUnexpectedErr, "Unexpected error. ")
	}

	if !matched {
		msg := fmt.Sprintf("Phone number is invalid, expected regex: %s but get: %s", PhoneRegex, phone)
		return aerror.NewError(nil, aerror.Code.CParamsErr, msg)
	}

	return nil
}

func AreaCodeValid(areaCode string) aerror.Error {
	matched, err := regexp.MatchString(AreaRegex, areaCode)
	if err != nil {
		return aerror.NewError(err, aerror.Code.SUnexpectedErr, "Unexpected error. ")
	}

	if !matched {
		msg := fmt.Sprintf("area code is invalid, expected regex: %s but code: %s", AreaRegex, areaCode)
		return aerror.NewError(err, aerror.Code.CParamsErr, msg)
	}

	return nil
}

func NicknameValid(nickname string) aerror.Error {
	matched, err := regexp.MatchString(NicknameRegex, nickname)
	if err != nil {
		return aerror.NewError(err, aerror.Code.SUnexpectedErr, "Unexpected error. ")
	}

	if !matched {
		return aerror.NewError(err, aerror.Code.CParamsErr, "昵称由 5 到 20 位字母、数字、下划线或者汉子组成")
	}

	return nil
}

func VerificationCodeValid(code string) aerror.Error {
	matched, err := regexp.MatchString(VerificationCodeRegex, code)
	if err != nil {
		return aerror.NewError(err, aerror.Code.SUnexpectedErr, "Unexpected error. ")
	}

	if !matched {
		return aerror.NewError(err, aerror.Code.CParamsErr, "验证码不正确")
	}

	return nil
}

func PasswordValid(str string) aerror.Error {
	var (
		isUpper   = false
		isLower   = false
		isNumber  = false
		isSpecial = false
	)

	if len(str) < 6 || len(str) > 16 {
		return aerror.NewError(nil, aerror.Code.CParamsErr,
			"密码由 6-16 位字母、数字、特殊字符(~!@#$%^&*)组成，至少包含一个数字和字母")
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

	if (isUpper || isLower) && (isNumber || isSpecial) {
		return nil
	}

	return aerror.NewError(nil, aerror.Code.CParamsErr,
		"密码由 6-16 位字母、数字、特殊字符(~!@#$%^&*)组成，至少包含一个数字和字母")
}
