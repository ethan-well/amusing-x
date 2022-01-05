package ganymedeservice

import (
	"github.com/ItsWewin/superfactory/aerror"
)

func (req *OAuthLoginRequest) Valid() aerror.Error {
	if len(req.Code) == 0 {
		return aerror.NewErrorf(nil, aerror.Code.CParamsError, "'code' is invalid")
	}

	if len(req.Provider) == 0 {
		return aerror.NewErrorf(nil, aerror.Code.CParamsError, "'provider' is invalid")
	}

	return nil
}
