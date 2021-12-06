package github

import (
	"amusingx.fit/amusingx/services/europa/oauth"
	"github.com/ItsWewin/superfactory/httputil"
	"github.com/ItsWewin/superfactory/httputil/rest"
	"github.com/ItsWewin/superfactory/xerror"
)

func (c *OAuth) GetUserProfile(url string, accessToken string) (*oauth.UserProfile, *xerror.Error) {
	opts := func(opts *rest.Options) {
		opts.Header = map[string]string{
			httputil.HeaderContent:       httputil.JsonHeaderContent,
			httputil.HeaderAccept:        httputil.JsonHeaderAccept,
			httputil.HeaderAuthorization: "token " + accessToken,
		}
	}
	var user oauth.UserProfile
	err := rest.Get(url, &user, opts)
	if err != nil {
		return nil, xerror.NewErrorf(err, xerror.Code.OtherNetworkError, "get user profile failed")
	}

	return &user, nil
}
