package model

import (
	"amusingx.fit/amusingx/mysqlstruct/ganymede"
	"context"
	"github.com/ItsWewin/superfactory/aerror"
)

func QueryCountryCodes(ctx context.Context) ([]*ganymede.CountryCode, aerror.Error) {
	var countryCodeList []*ganymede.CountryCode

	query := "SELECT id, country_code, cname FROM country_code"
	err := GanymedeDB.SelectContext(ctx, &countryCodeList, query)
	if err != nil {
		return countryCodeList, aerror.NewError(err, aerror.Code.SSqlExecuteErr, "查询 country code 失败")
	}

	return countryCodeList, nil
}
