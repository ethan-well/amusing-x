package service

import (
	"github.com/ItsWewin/superfactory/aerror"
)

func (req *OAuthInfoRequest) Valid() aerror.Error {
	if req == nil || len(req.Provider) == 0 {
		return aerror.NewErrorf(nil, aerror.Code.CParamsErr, "'provider' is invalid")
	}

	if len(req.Service) == 0 {
		return aerror.NewErrorf(nil, aerror.Code.CParamsErr, "'service' is invalid")
	}

	return nil
}
