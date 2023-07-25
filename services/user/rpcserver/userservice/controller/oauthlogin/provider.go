package oauthlogin

import (
	"amusingx.fit/amusingx/services/user/conf"
	"amusingx.fit/amusingx/services/user/oauth/oauthstruct"
	"context"
	"github.com/ItsWewin/superfactory/aerror"
)

func GetProvider(ctx context.Context, provider, clientServer string) (*conf.OAuthProviderInfo, aerror.Error) {
	switch provider {
	case oauthstruct.ProviderGitHub:
		return getGithubProvider(clientServer)
	case oauthstruct.ProviderWeChat:
		return getWechatProvider(clientServer)
	default:
		return nil, aerror.NewErrorf(nil, aerror.Code.CParamsErr, "provider '%s' is invalid", provider)
	}
}

func getGithubProvider(server string) (*conf.OAuthProviderInfo, aerror.Error) {
	switch server {
	case conf.Conf.OAuth.Github.ForAmusingX.ClientServer:
		return conf.Conf.OAuth.Github.ForAmusingX, nil
	case conf.Conf.OAuth.Github.ForPanGu.ClientServer:
		return conf.Conf.OAuth.Github.ForPanGu, nil
	default:
		return nil, aerror.NewErrorf(nil, aerror.Code.CParamsErr, "client server: %s is invalid", server)
	}
}

func getWechatProvider(server string) (*conf.OAuthProviderInfo, aerror.Error) {
	switch server {
	case conf.Conf.OAuth.WeChat.ForAmusingX.ClientServer:
		return conf.Conf.OAuth.WeChat.ForAmusingX, nil
	case conf.Conf.OAuth.WeChat.ForPanGu.ClientSecret:
		return conf.Conf.OAuth.WeChat.ForPanGu, nil
	default:
		return nil, aerror.NewErrorf(nil, aerror.Code.CParamsErr, "client server is invalid")
	}
}
