package github

import (
	"amusingx.fit/amusingx/apistruct/github"
	"amusingx.fit/amusingx/services/user/oauth/oauthstruct"
	"github.com/ItsWewin/superfactory/aerror"
	"github.com/ItsWewin/superfactory/httputil"
	"github.com/ItsWewin/superfactory/httputil/rest"
)

func (c *OAuth) GetUserProfile(url string, accessToken string) (*oauthstruct.UserProfile, aerror.Error) {
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
		return nil, aerror.NewErrorf(err, aerror.Code.OtherNetworkErr, "get user profile failed")
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
