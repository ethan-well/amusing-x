package lock

import (
	"amusingx.fit/amusingx/services/webApi/xredis"
	"context"
	"github.com/ItsWewin/superfactory/db/redis/redislocker"
	"github.com/ItsWewin/superfactory/logger"
)

type Locker struct {
	Key    string
	Value  interface{}
	Expire int64
}

func NewLocker(key, value string, expire int64) *Locker {
	return &Locker{
		Key:    key,
		Value:  value,
		Expire: expire,
	}
}

func (l *Locker) Lock(ctx context.Context) error {
	return redislocker.Lock(ctx, xredis.Client, l.Key, l.Value, l.Expire)
}

func (l *Locker) Unlock(ctx context.Context) error {
	err := redislocker.Unlock(ctx, xredis.Client, l.Key)
	if err != nil {
		logger.Errorf("redis unlock failed, key: %s, err: %s", l.Key, err)

		return err
	}

	return nil
}

// key 和 value 同时相等的时候，更新过期时间
func (l *Locker) Refresh(ctx context.Context) error {
	return redislocker.Refresh(ctx, xredis.Client, l.Key, l.Value, l.Expire)
}
