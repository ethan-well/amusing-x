package amusingxwebapi

import (
	"amusingx.fit/amusingx/mysqlstruct/ganymede"
	"context"
)

func QueryCountryCodes(ctx context.Context) ([]*ganymede.CountryCode, error) {
	var countryCodeList []*ganymede.CountryCode

	query := "SELECT id, country_code, cname FROM country_code"
	err := AmusingUserDB.SelectContext(ctx, &countryCodeList, query)
	if err != nil {
		return countryCodeList, err
	}

	return countryCodeList, nil
}
