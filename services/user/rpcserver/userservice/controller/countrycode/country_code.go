package countrycode

import (
	service "amusingx.fit/amusingx/protos/user/service/user/proto"
	"amusingx.fit/amusingx/services/user/mysql/ganymededb/model"
	"context"
	"github.com/ItsWewin/superfactory/aerror"
)

func HandleCountryCodeList(ctx context.Context) (*service.CountryCodeList, aerror.Error) {
	codeList, err := model.QueryCountryCodes(ctx)
	if err != nil {
		return nil, err
	}

	var countryCodeList = make([]*service.CountryCode, 0)
	for _, code := range codeList {
		countryCodeList = append(countryCodeList, &service.CountryCode{
			Id:         code.ID,
			Cname:      code.Cname,
			Code:       code.Code,
			CreateTime: code.CreateTime,
			UpdateTime: code.UpdateTime,
		})
	}

	return &service.CountryCodeList{CountryCodeList: countryCodeList}, nil
}
