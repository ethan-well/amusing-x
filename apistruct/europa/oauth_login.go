package europa

import "github.com/ItsWewin/superfactory/aerror"

type OAuthLogin struct {
	Code     string `json:"code"`
	Provider string `json:"provider"`
	Service  string `json:"service"`
}

func (login *OAuthLogin) Valid() aerror.Error {
	if len(login.Code) == 0 {
		return aerror.NewErrorf(nil, aerror.Code.CParamsError, "code is blank")
	}

	if len(login.Provider) == 0 {
		return aerror.NewErrorf(nil, aerror.Code.CParamsError, "provider is blank")
	}

	return nil
}
