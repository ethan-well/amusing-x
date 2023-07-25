package charonservice

import "github.com/ItsWewin/superfactory/aerror"

func (x *CategoryCreateRequest) Valid() aerror.Error {
	if len(x.Name) == 0 {
		return aerror.NewErrorf(nil, aerror.Code.CParamsErr, "params 'name' is invalid")
	}

	if len(x.Desc) == 0 {
		return aerror.NewErrorf(nil, aerror.Code.CParamsErr, "params 'desc' is invalid")
	}

	return nil
}
