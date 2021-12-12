package ganymedeservice

import "github.com/ItsWewin/superfactory/xerror"

func (req *OAuthInfoRequest) Valid() *xerror.Error {
	if req == nil || len(req.Provider) == 0 {
		return xerror.NewErrorf(nil, xerror.Code.CParamsError, "'provider' is invalid")
	}

	return nil
}
