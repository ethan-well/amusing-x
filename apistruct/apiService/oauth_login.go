package apiService

import "github.com/ItsWewin/superfactory/aerror"

type OAuthLogin struct {
	Code     string `json:"code"`
	Provider string `json:"provider"`
	Service  string `json:"service"`
}

func (login *OAuthLogin) Valid() aerror.Error {
	if len(login.Code) == 0 {
		return aerror.NewErrorf(nil, aerror.Code.CParamsErr, "code is blank")
	}

	if len(login.Provider) == 0 {
		return aerror.NewErrorf(nil, aerror.Code.CParamsErr, "provider is blank")
	}

	return nil
}
