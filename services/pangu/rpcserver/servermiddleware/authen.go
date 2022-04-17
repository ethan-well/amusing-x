package servermiddleware

import (
	"amusingx.fit/amusingx/services/pangu/rpcserver/handler/login"
	"amusingx.fit/amusingx/services/pangu/rpcserver/panguserver"
	"context"
	"github.com/ItsWewin/superfactory/aerror"
	"github.com/ItsWewin/superfactory/logger"
	"google.golang.org/grpc"
)

var routers = map[string]struct{}{
	"/panguservice.PanGuService/CategoryList": {},
}

func UnaryServerInterceptorAuthentication() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		defer func() {
			if err := recover(); err != nil {
				logger.Errorf("server panic: %s", err)
			}
		}()

		logger.Errorf("FullMethod: %s", info.FullMethod)

		sessionId, err := panguserver.GetSessionID(ctx)
		if err != nil {
			return nil, err
		}

		logger.Errorf("sessionId: %s", sessionId)

		// skip authentication
		if _, ok := routers[info.FullMethod]; !ok {
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

	logger.Infof("sessionId xxx: %s", sessionId)

	resp, err := login.HandlerIsLogin(ctx, sessionId)
	if err != nil {
		logger.Errorf("authenticationFromCookie err: %s", err)
		return err
	}

	logger.Infof("resp: %s", logger.ToJson(resp))

	if !resp.Login {
		return aerror.NewErrorf(nil, aerror.Code.CForbidden, "user no login")
	}

	return nil
}
