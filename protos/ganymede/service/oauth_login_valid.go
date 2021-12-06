package userservice

import "github.com/ItsWewin/superfactory/xerror"

func (req *OAuthLoginRequest) Valid() *xerror.Error {
	if len(req.Code) == 0 {
		return xerror.NewErrorf(nil, xerror.Code.CParamsError, "'code' is invalid")
	}

	if len(req.Provider) == 0 {
		return xerror.NewErrorf(nil, xerror.Code.CParamsError, "'provider' is invalid")
	}

	return nil
}
