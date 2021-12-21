package session

import (
	"amusingx.fit/amusingx/services/ganymede/session"
	"context"
	"encoding/json"
	"github.com/ItsWewin/superfactory/aerror"
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

func (m *Model) SetSession(ctx context.Context, info *Info) (string, aerror.Error) {
	uid := uuid.NewV4().String()
	sess, err := session.GlobalSessionManager.Store.SessionInit(ctx, uid)
	if err != nil {
		return "", aerror.NewError(nil, aerror.Code.SUnexpectedErr, "get session failed")
	}

	err = sess.Set(ctx, "id", info.UserID)

	data, err := json.Marshal(info)
	if err != nil {
		return "", aerror.NewErrorf(nil, aerror.Code.BUnexpectedData, "json marshal failed")
	}
	err = sess.Set(ctx, "details", string(data))
	if err != nil {
		return "", aerror.NewError(nil, aerror.Code.SUnexpectedErr, "set session id failed")
	}

	return uid, nil
}

func (m *Model) GetUserID(ctx context.Context, sid string) (int64, aerror.Error) {
	sess, err := session.GlobalSessionManager.Store.SessionInit(ctx, sid)
	if err != nil {
		return 0, aerror.NewError(nil, aerror.Code.SUnexpectedErr, "get session failed")
	}

	v, err := sess.Get(ctx, "id")

	switch v.(type) {
	case int64:
		return v.(int64), nil
	case string:
		id, err := strconv.ParseInt(v.(string), 10, 64)
		return id, aerror.NewErrorf(err, aerror.Code.BUnexpectedData, "get user id failed")
	case float64:
		return int64(v.(float64)), nil
	default:
		return 0, aerror.NewErrorf(nil, aerror.Code.BUnexpectedData, "no user id")
	}
}

func (m *Model) Delete(ctx context.Context, sid string) aerror.Error {
	sess, err := session.GlobalSessionManager.Store.SessionInit(ctx, sid)
	if err != nil {
		return aerror.NewError(err, aerror.Code.SUnexpectedErr, "session store init failed")
	}

	err = sess.Delete(ctx, sid)
	if err != nil {
		return aerror.NewError(err, aerror.Code.SUnexpectedErr, "delete session failed")
	}

	return nil
}
