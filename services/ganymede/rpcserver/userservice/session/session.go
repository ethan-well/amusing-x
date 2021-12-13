package session

import (
	"amusingx.fit/amusingx/services/ganymede/session"
	"context"
	"encoding/json"
	"github.com/ItsWewin/superfactory/xerror"
	uuid "github.com/satori/go.uuid"
	"strconv"
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

	switch v.(type) {
	case int64:
		return v.(int64), nil
	case string:
		id, err := strconv.ParseInt(v.(string), 10, 64)
		return id, xerror.NewErrorf(err, xerror.Code.BUnexpectedData, "get user id failed")
	case float64:
		return int64(v.(float64)), nil
	default:
		return 0, xerror.NewErrorf(nil, xerror.Code.BUnexpectedData, "no user id")
	}
}

func (m *Model) Delete(ctx context.Context, sid string) *xerror.Error {
	sess, err := session.GlobalSessionManager.Store.SessionInit(ctx, sid)
	if err != nil {
		return xerror.NewError(nil, xerror.Code.SUnexpectedErr, "destroy session failed")
	}

	err = sess.Delete(ctx, sid)
	if err != nil {
		return xerror.NewError(nil, xerror.Code.SUnexpectedErr, "delete session failed")
	}

	return nil
}
