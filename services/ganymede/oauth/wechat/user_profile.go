package wechat

import (
	"amusingx.fit/amusingx/apistruct/wechat"
	"amusingx.fit/amusingx/services/ganymede/oauth"
	"fmt"
	"github.com/ItsWewin/superfactory/httputil"
	"github.com/ItsWewin/superfactory/httputil/rest"
	"github.com/ItsWewin/superfactory/lang"
	"github.com/ItsWewin/superfactory/xerror"
	"net/url"
)

func (o *OAuth) GetUserProfile(profileUrl string, accessToken string) (*oauth.UserProfile, *xerror.Error) {
	var params = url.Values{}
	params.Set("openid", o.ClientID)
	params.Set("access_token", o.ClientSecret)
	params.Set("lang", lang.En)

	u, err := url.Parse(profileUrl)
	if err != nil {
		return nil, xerror.NewErrorf(nil, xerror.Code.BUnexpectedData, "profileUrl format err")
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
		return nil, xerror.NewErrorf(nil, xerror.Code.OtherNetworkError, dest.Errmsg)
	}

	return &oauth.UserProfile{
		Login:       dest.Unionid,
		OuterUserID: 0,
		AvatarUrl:   dest.AvatarUrl,
		Name:        dest.Nickname,
		Company:     ,
		Blog:        "",
		Email:       "",
		Location:    fmt.Sprintf("%s,%s,%s", dest.Country, dest.Province, dest.City),
	}, nil


}
