package amusinguser

import (
	"amusingx.fit/amusingx/mysqlstruct/amusinguser"
	"context"
	"github.com/ItsWewin/superfactory/xerror"
)

func QueryCountryCodes(ctx context.Context) ([]*amusinguser.CountryCode, *xerror.Error) {
	var countryCodeList []*amusinguser.CountryCode

	query := "SELECT id, country_code, cname FROM country_code"
	err := AmusingUserDB.SelectContext(ctx, &countryCodeList, query)
	if err != nil {
		return countryCodeList, xerror.NewError(err, xerror.Code.SSqlExecuteErr, "查询 country code 失败")
	}

	return countryCodeList, nil
}
