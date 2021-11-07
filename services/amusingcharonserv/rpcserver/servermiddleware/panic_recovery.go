package servermiddleware

import (
	"context"
	"github.com/ItsWewin/superfactory/logger"
	"google.golang.org/grpc"
)

func UnaryServerInterceptorPanicRecover() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		defer func() {
			if err := recover(); err != nil {
				logger.Errorf("server panic: %s", err)
			}
		}()

		return handler(ctx, req)
	}
}

func StreamServerInterceptorPanicRecover() grpc.StreamServerInterceptor {
	return func(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		defer func() {
			if err := recover(); err != nil {
				logger.Errorf("server panic: %s", err)
			}
		}()

		return handler(srv, ss)
	}
}
