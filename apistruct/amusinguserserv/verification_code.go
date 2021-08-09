package amusinguserserv

import (
	"amusingx.fit/amusingx/services/amusinguserserv/regexp"
	"github.com/ItsWewin/superfactory/xerror"
)

type VerificationCodeRequest struct {
	Phone string `json:"phone"`
	// login or join
	Action   string `json:"action"`
	AreaCode string `json:"area_code"`
}

const login = "login"
const join = "join"

var verificationCodeActions = []string{login, join}

func (r *VerificationCodeRequest) IsLogin() bool {
	return r.Action == login
}

func (r *VerificationCodeRequest) IsJoin() bool {
	return r.Action == join
}

func (r *VerificationCodeRequest) Valid() *xerror.Error {
	if err := regexp.PhoneNumberValid(r.Phone); err != nil {
		return err
	}

	if err := regexp.AreaCodeValid(r.AreaCode); err != nil {
		return err
	}

	actionValid := false
	for _, action := range verificationCodeActions {
		if r.Action == action {
			actionValid = true
			break
		}
	}
	if !actionValid {
		return xerror.NewError(nil, xerror.Code.CParamsError, "'action' s is invalid")
	}

	return nil
}
