package europa

import (
	"amusingx.fit/amusingx/regexp"
	"github.com/ItsWewin/superfactory/aerror"
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

func (r *VerificationCodeRequest) Valid() aerror.Error {
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
		return aerror.NewError(nil, aerror.Code.CParamsError, "'action' s is invalid")
	}

	return nil
}
