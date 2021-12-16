package session

import (
	"context"
	"github.com/ItsWewin/superfactory/xerror"
)

type RedisSession struct {
	Sid   string
	Value map[string]interface{}
	Store Store
}

func (s *RedisSession) Set(ctx context.Context, key string, value interface{}) error {
	if len(s.Sid) == 0 {
		return xerror.NewError(nil, xerror.Code.BUnexpectedBlankVariable, "session is not init")
	}

	if s.Value == nil {
		s.Value = make(map[string]interface{})
	}

	s.Value[key] = value

	// 写入 Redis
	err := s.Store.SessionWrite(ctx, s.Sid, s.Value)
	if err != nil {
		return err
	}

	return nil
}

func (s *RedisSession) Get(ctx context.Context, key string) (interface{}, error) {
	if len(s.Sid) == 0 {
		return nil, xerror.NewError(nil, xerror.Code.BUnexpectedBlankVariable, "session is not init")
	}

	if s.Value == nil {
		session, err := s.Store.SessionRead(ctx, s.Sid)
		if err != nil {
			return nil, xerror.NewError(err, xerror.Code.SRedisExecuteErr, "sessionRead failed")
		}

		return session.(*RedisSession).Value[key], nil
	}

	return s.Value[key], nil
}

func (s *RedisSession) Delete(ctx context.Context, key string) error {
	if len(s.Sid) == 0 {
		return xerror.NewError(nil, xerror.Code.BUnexpectedBlankVariable, "session is not init")
	}

	if s.Value != nil {
		// 删除键
		delete(s.Value, key)
	}

	// 写入 Redis
	err := s.Store.SessionDestroy(ctx, s.Sid)
	if err != nil {
		return err
	}

	return nil
}

func (s *RedisSession) SessionID() string {
	return s.Sid
}
