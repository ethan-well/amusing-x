package servermiddleware

import (
	"amusingx.fit/amusingx/services/admin/rpcserver/handler/login"
	"amusingx.fit/amusingx/services/admin/rpcserver/panguserver"
	"context"
	"github.com/ItsWewin/superfactory/aerror"
	"github.com/ItsWewin/superfactory/logger"
	"google.golang.org/grpc"
)

var skipAuthMethods = map[string]struct{}{
	"/panguservice.PanGuService/Pong":              {},
	"/panguservice.PanGuService/Logout":            {},
	"/panguservice.PanGuService/OauthLogin":        {},
	"/panguservice.PanGuService/OauthProviderInfo": {},
}

func UnaryServerInterceptorAuthentication() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		defer func() {
			if err := recover(); err != nil {
				logger.Errorf("server panic: %s", err)
			}
		}()

		// skip authentication
		if _, ok := skipAuthMethods[info.FullMethod]; ok {
			return handler(ctx, req)
		}

		err = authenticationFromCookie(ctx)
		if err != nil {
			return nil, forbiddenErr(nil)
		}

		return handler(ctx, req)
	}
}

func StreamServerInterceptorAuthentication() grpc.StreamServerInterceptor {
	return func(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		defer func() {
			if err := recover(); err != nil {
				logger.Errorf("server panic: %s", err)
			}
		}()

		return handler(srv, ss)
	}
}

func forbiddenErr(err error) aerror.Error {
	return aerror.NewErrorf(err, aerror.Code.CForbidden, "permission denied")
}

func authenticationFromCookie(ctx context.Context) aerror.Error {
	sessionId, err := panguserver.GetSessionID(ctx)
	if err != nil {
		return err
	}

	resp, err := login.HandlerIsLogin(ctx, sessionId)
	if err != nil {
		logger.Errorf("authenticationFromCookie err: %s", err)
		return err
	}

	if !resp.Login {
		return aerror.NewErrorf(nil, aerror.Code.CForbidden, "user no login")
	}

	return nil
}
