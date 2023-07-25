package panguserver

import (
	panguservice "amusingx.fit/amusingx/protos/pangu/service/pangu/proto"
	"amusingx.fit/amusingx/services/admin/conf"
	"amusingx.fit/amusingx/services/admin/rpcserver/getway"
	"amusingx.fit/amusingx/services/admin/rpcserver/handler/login"
	"bufio"
	"context"
	"fmt"
	"github.com/ItsWewin/superfactory/aerror"
	"github.com/ItsWewin/superfactory/logger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"net/http"
	"strings"
)

type PanguServer struct {
	panguservice.UnimplementedPanGuServiceServer
}

// Pong(context.Context, *BlankParams) (*PongResponse, error)
func (s *PanguServer) Pong(ctx context.Context, in *panguservice.BlankParams) (*panguservice.PongResponse, error) {
	md, _ := metadata.FromIncomingContext(ctx)
	logger.Infof("cookie: %#v", md.Get("grpcgateway-cookie"))

	return &panguservice.PongResponse{ServerName: conf.Conf.Server.Name}, nil
}

func (s *PanguServer) OauthLogin(ctx context.Context, in *panguservice.OAuthLoginRequest) (*panguservice.OAuthLoginResponse, error) {
	login, err := login.HandlerOauthLogin(ctx, in)
	if err != nil || login == nil || login.Result.UserInfo == nil || login.Result.SessionInfo == nil {
		logger.Errorf("err: %s", err)
		return nil, aerror.NewError(err, aerror.Code.BUnexpectedData, "login failed")
	}

	header := metadata.Pairs(
		getway.SessionIdKey, login.Result.SessionInfo.SessionID,
	)
	grpc.SetHeader(ctx, header)

	return login, nil
}

func (s *PanguServer) OauthProviderInfo(ctx context.Context, in *panguservice.OauthProviderInfoRequest) (*panguservice.OAuthProviderInfoResponse, error) {
	return login.HandlerOauthProviderInfo(ctx, in)
}

func (s *PanguServer) Logout(ctx context.Context, in *panguservice.BlankParams) (*panguservice.LogoutResponse, error) {
	sessionID, err := GetSessionID(ctx)
	if err != nil {
		return nil, err
	}

	logger.Infof("sessionID from cookie: %s", sessionID)

	if len(sessionID) == 0 {
		return &panguservice.LogoutResponse{Succeed: true}, nil
	}

	result, err := login.HandlerLogout(ctx, sessionID)
	if err != nil {
		return nil, err
	}

	if result.Succeed {
		header := metadata.Pairs(
			getway.DeleteSessionKey, "true",
		)
		grpc.SetHeader(ctx, header)
	}

	return result, nil
}

func GetSessionID(ctx context.Context) (string, aerror.Error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", aerror.NewErrorf(nil, aerror.Code.CParamsErr, "get session id failed")
	}

	rawCookies := strings.Join(md.Get(getway.GrpcGatewayCookieKey), "")
	if len(rawCookies) == 0 {
		return "", nil
	}

	rawRequest := fmt.Sprintf("GET / HTTP/1.0\r\nCookie: %s\r\n\r\n", rawCookies)
	req, err := http.ReadRequest(bufio.NewReader(strings.NewReader(rawRequest)))
	if err != nil {
		return "", aerror.NewErrorf(nil, aerror.Code.CParamsErr, "get session id failed, cookie info is invalid")
	}

	cookies := req.Cookies()
	for _, cookie := range cookies {
		if cookie.Name == getway.GrpcGatewayCookieSessionKey {
			return cookie.Value, nil
		}
	}

	return "", nil
}
