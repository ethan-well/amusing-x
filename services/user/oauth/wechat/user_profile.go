package wechat

import (
	"amusingx.fit/amusingx/apistruct/wechat"
	"amusingx.fit/amusingx/services/user/oauth/oauthstruct"
	"fmt"
	"github.com/ItsWewin/superfactory/aerror"
	"github.com/ItsWewin/superfactory/httputil"
	"github.com/ItsWewin/superfactory/httputil/rest"
	"github.com/ItsWewin/superfactory/lang"
	"net/url"
)

func (o *OAuth) GetUserProfile(profileUrl string, accessToken string) (*oauthstruct.UserProfile, aerror.Error) {
	var params = url.Values{}
	params.Set("openid", o.ClientID)
	params.Set("access_token", o.ClientSecret)
	params.Set("lang", lang.En)

	u, err := url.Parse(profileUrl)
	if err != nil {
		return nil, aerror.NewErrorf(nil, aerror.Code.BUnexpectedData, "profileUrl format err")
	}
	u.RawQuery = params.Encode()

	opts := func(opts *rest.Options) {
		opts.Header = map[string]string{
			httputil.HeaderContent: httputil.JsonHeaderContent,
			httputil.HeaderAccept:  httputil.JsonHeaderAccept,
		}
	}

	var dest wechat.UserProfile
	xErr := rest.Get(u.String(), &dest, opts)
	if xErr != nil {
		return nil, xErr
	}

	if dest.IsError() {
		return nil, aerror.NewErrorf(nil, aerror.Code.OtherNetworkErr, dest.Errmsg)
	}

	return &oauthstruct.UserProfile{
		Login:       dest.Unionid,
		OuterUserID: 0,
		AvatarUrl:   dest.AvatarUrl,
		Name:        dest.Nickname,
		Company:     "",
		Blog:        "",
		Email:       "",
		Location:    fmt.Sprintf("%s,%s,%s", dest.Country, dest.Province, dest.City),
	}, nil
}
