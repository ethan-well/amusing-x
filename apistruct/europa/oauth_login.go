package europa

import "github.com/ItsWewin/superfactory/xerror"

type OAuthLogin struct {
	Code     string `json:"code"`
	Provider string `json:"provider"`
}

func (login *OAuthLogin) Valid() *xerror.Error {
	if len(login.Code) == 0 {
		return xerror.NewErrorf(nil, xerror.Code.CParamsError, "code is blank")
	}

	if len(login.Provider) == 0 {
		return xerror.NewErrorf(nil, xerror.Code.CParamsError, "provider is blank")
	}

	return nil
}
