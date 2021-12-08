package countrycode

import (
	"amusingx.fit/amusingx/protos/ganymede/service"
	"amusingx.fit/amusingx/services/ganymede/mysql/ganymededb/model"
	"context"
	"github.com/ItsWewin/superfactory/xerror"
)

func HandleCountryCodeList(ctx context.Context) (*ganymedeservice.CountryCodeList, *xerror.Error) {
	codeList, err := model.QueryCountryCodes(ctx)
	if err != nil {
		return nil, err
	}

	var countryCodeList = make([]*ganymedeservice.CountryCode, 0)
	for _, code := range codeList {
		countryCodeList = append(countryCodeList, &ganymedeservice.CountryCode{
			Id:         code.ID,
			Cname:      code.Cname,
			Code:       code.Code,
			CreateTime: code.CreateTime,
			UpdateTime: code.UpdateTime,
		})
	}

	return &ganymedeservice.CountryCodeList{CountryCodeList: countryCodeList}, nil
}
