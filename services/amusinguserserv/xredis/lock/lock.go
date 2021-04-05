package lock

import (
	"amusingx.fit/amusingx/services/amusinguserserv/xredis"
	"context"
	"errors"
	"fmt"
	"github.com/go-redis/redis/v8"
)

type Locker struct {
	Key    string
	Value  string
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
	script := redis.NewScript(`
		if redis.call("SETNX", KEYS[1], ARGV[1]) == 1 then
			return redis.call("expire", KEYS[1], ARGV[2])
		else
			return 0
		end
	`)

	result, err := script.Run(ctx, xredis.Client, []string{l.Key}, l.Value, l.Expire).Result()
	if err != nil {
		return err
	}

	if result.(int64) != 1 {
		return errors.New(fmt.Sprintf("lock failed, result: %d", result))
	}

	return nil
}

func (l *Locker) Unlock(ctx context.Context) error {
	script := redis.NewScript(`return redis.call('del', KEYS[1]) ~= false`)

	_, err := script.Run(ctx, xredis.Client, []string{l.Key}).Result()
	if err != nil {
		return err
	}

	return nil
}

// key 和 value 同时相等的时候，更新过期时间
func (l *Locker) Refresh(ctx context.Context) error {
	script := redis.NewScript(`if redis.call("get", KEYS[1]) == ARGV[1] then return redis.call("expire", KEYS[1], ARGV[2]) else return 0 end`)

	r, err := script.Run(ctx, xredis.Client, []string{l.Key}, l.Value, l.Expire).Result()
	if err != nil {
		return err
	}
	if r.(int64) != 1 {
		return errors.New(fmt.Sprintf("refresh failed, key: %s, value: %s, expire: %d", l.Key, l.Value, l.Expire))
	}

	return nil
}
