package service

import (
	"github.com/ItsWewin/superfactory/aerror"
)

func (req *OAuthLoginRequest) Valid() aerror.Error {
	if len(req.Code) == 0 {
		return aerror.NewErrorf(nil, aerror.Code.CParamsErr, "'code' is invalid")
	}

	if len(req.Provider) == 0 {
		return aerror.NewErrorf(nil, aerror.Code.CParamsErr, "'provider' is invalid")
	}

	return nil
}
