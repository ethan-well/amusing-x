package session

import (
	"amusingx.fit/amusingx/services/ganymede/session"
	"context"
	"encoding/json"
	"github.com/ItsWewin/superfactory/xerror"
	uuid "github.com/satori/go.uuid"
)

type Info struct {
	UserID int64 `json:"user_id"`
}

type Model struct {
}

func New() *Model {
	return &Model{}
}

func (m *Model) SetSession(ctx context.Context, info *Info) (string, *xerror.Error) {
	uid := uuid.NewV4().String()
	sess, err := session.GlobalSessionManager.Store.SessionInit(ctx, uid)
	if err != nil {
		return "", xerror.NewError(nil, xerror.Code.SUnexpectedErr, "get session failed")
	}

	err = sess.Set(ctx, "id", info.UserID)

	data, err := json.Marshal(info)
	if err != nil {
		return "", xerror.NewErrorf(nil, xerror.Code.BUnexpectedData, "json marshal failed")
	}
	err = sess.Set(ctx, "details", string(data))
	if err != nil {
		return "", xerror.NewError(nil, xerror.Code.SUnexpectedErr, "set session id failed")
	}

	return uid, nil
}

func (m *Model) GetUserID(ctx context.Context, sid string) (int64, *xerror.Error) {
	sess, err := session.GlobalSessionManager.Store.SessionInit(ctx, sid)
	if err != nil {
		return 0, xerror.NewError(nil, xerror.Code.SUnexpectedErr, "get session failed")
	}

	v, err := sess.Get(ctx, "id")

	id, ok := v.(int64)
	if !ok {
		return 0, xerror.NewErrorf(nil, xerror.Code.BUnexpectedData, "no user id")
	}

	return id, nil
}
