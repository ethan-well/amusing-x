package countrycode

import (
	"amusingx.fit/amusingx/protos/ganymede/service"
	"amusingx.fit/amusingx/services/ganymede/mysql/ganymede/model"
	"context"
	"github.com/ItsWewin/superfactory/xerror"
)

func HandleCountryCodeList(ctx context.Context) (*userservice.CountryCodeList, *xerror.Error) {
	codeList, err := model.QueryCountryCodes(ctx)
	if err != nil {
		return nil, err
	}

	var countryCodeList = make([]*userservice.CountryCode, 0)
	for _, code := range codeList {
		countryCodeList = append(countryCodeList, &userservice.CountryCode{
			Id:         code.ID,
			Cname:      code.Cname,
			Code:       code.Code,
			CreateTime: code.CreateTime,
			UpdateTime: code.UpdateTime,
		})
	}

	return &userservice.CountryCodeList{CountryCodeList: countryCodeList}, nil
}
