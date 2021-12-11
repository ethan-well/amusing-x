package github

import (
	"amusingx.fit/amusingx/apistruct/github"
	"amusingx.fit/amusingx/services/ganymede/oauth/oauthstruct"
	"github.com/ItsWewin/superfactory/httputil"
	"github.com/ItsWewin/superfactory/httputil/rest"
	"github.com/ItsWewin/superfactory/xerror"
)

func (c *OAuth) GetUserProfile(url string, accessToken string) (*oauthstruct.UserProfile, *xerror.Error) {
	opts := func(opts *rest.Options) {
		opts.Header = map[string]string{
			httputil.HeaderContent:       httputil.JsonHeaderContent,
			httputil.HeaderAccept:        httputil.JsonHeaderAccept,
			httputil.HeaderAuthorization: "token " + accessToken,
		}
	}
	var user github.UserProfile
	err := rest.Get(url, &user, opts)
	if err != nil {
		return nil, xerror.NewErrorf(err, xerror.Code.OtherNetworkError, "get user profile failed")
	}

	return &oauthstruct.UserProfile{
		Login:       user.Login,
		OuterUserID: user.ID,
		AvatarUrl:   user.AvatarUrl,
		Name:        user.Name,
		Company:     user.Company,
		Blog:        user.Blog,
		Email:       user.Email,
		Location:    user.Location,
	}, nil
}
