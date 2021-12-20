package charonservice

import "github.com/ItsWewin/superfactory/xerror"

func (x *CategoryCreateRequest) Valid() *xerror.Error {
	if len(x.Name) != 0 {
		return xerror.NewErrorf(nil, xerror.Code.CParamsError, "params 'name' is invalid")
	}

	if len(x.Name) != 0 {
		return xerror.NewErrorf(nil, xerror.Code.CParamsError, "params 'name' is invalid")
	}

	return nil
}
