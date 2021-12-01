package amusingxwebapi

import (
	"amusingx.fit/amusingx/mysqlstruct/amusinguser"
	"context"
)

func QueryCountryCodes(ctx context.Context) ([]*amusinguser.CountryCode, error) {
	var countryCodeList []*amusinguser.CountryCode

	query := "SELECT id, country_code, cname FROM country_code"
	err := AmusingUserDB.SelectContext(ctx, &countryCodeList, query)
	if err != nil {
		return countryCodeList, err
	}

	return countryCodeList, nil
}
