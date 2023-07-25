package getway

import (
	panguservice "amusingx.fit/amusingx/protos/pangu/service/pangu/proto"
	"amusingx.fit/amusingx/services/admin/rpcserver/session"
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/proto"
	"net/http"
)

const (
	DeleteSessionKey            = "delete-session"
	SessionIdKey                = "session_id"
	GrpcGatewayCookieKey        = "grpcgateway-cookie"
	GrpcGatewayCookieSessionKey = "user_session"
)

func CustomMatcher(key string) (string, bool) {
	switch key {
	case "":
		return key, true
	default:
		return runtime.DefaultHeaderMatcher(key)
	}
}

func HttpResponseModifier(ctx context.Context, w http.ResponseWriter, resp proto.Message) error {
	md, ok := runtime.ServerMetadataFromContext(ctx)
	if !ok {
		return nil
	}

	setSession(w, md)

	delSession(w, md)

	return nil
}

func delSession(w http.ResponseWriter, md runtime.ServerMetadata) {
	sessionIds := md.HeaderMD.Get(DeleteSessionKey)
	if len(sessionIds) == 0 {
		return
	}
	session.DeleteSession(w)
}

func setSession(w http.ResponseWriter, md runtime.ServerMetadata) {
	sessionIds := md.HeaderMD.Get(SessionIdKey)
	if len(sessionIds) == 0 {
		return
	}
	sessionInfo := &panguservice.SessionInfo{
		SessionID: sessionIds[0],
		MaxAge:    1000,
	}
	session.SetSession(w, sessionInfo)
}

// look up session and pass userId in to context if it exists
func GatewayMetadataAnnotator(_ context.Context, r *http.Request) metadata.MD {
	return nil
}
