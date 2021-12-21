package oauth

import (
	"context"
	"fmt"
	"github.com/ItsWewin/superfactory/aerror"
	"github.com/go-redis/redis/v8"
	"math/rand"
	"time"
)

type Store interface {
	RandomStateCode(context.Context) (string, aerror.Error)
	ValidSate(context.Context, string) aerror.Error
}

type RedisStore struct {
	Client        *redis.Client
	KeyPrefix     string
	MaxLifeSecond int64
}

var StoreIns Store

func InitOAuthRedisStore(addr, password string, db int, prefix string, maxLifeSecond int64) aerror.Error {
	if len(addr) == 0 {
		return aerror.NewError(nil, aerror.Code.BUnexpectedBlankVariable, "addr is blank")
	}

	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})

	StoreIns = &RedisStore{Client: client, KeyPrefix: prefix, MaxLifeSecond: maxLifeSecond}

	return nil
}

func (s *RedisStore) key(code string) string {
	return fmt.Sprintf("oauth:sate:store:%s:%s" + s.KeyPrefix + code)
}

func (s *RedisStore) save(ctx context.Context, code string) aerror.Error {

	cmd := s.Client.SetEX(ctx, s.key(code), code, time.Duration(s.MaxLifeSecond)*time.Second)
	if cmd.Err() != nil {
		return aerror.NewErrorf(nil, aerror.Code.SRedisExecuteErr, cmd.Err().Error())
	}

	return nil
}

func (s *RedisStore) get(ctx context.Context, code string) (string, aerror.Error) {
	cmd := s.Client.Get(ctx, s.key(code))
	if cmd.Err() != nil {
		return "", aerror.NewErrorf(nil, aerror.Code.SRedisExecuteErr, cmd.Err().Error())
	}

	code, err := cmd.Result()
	if err != nil {
		return "", aerror.NewErrorf(nil, aerror.Code.SRedisExecuteErr, err.Error())
	}

	return code, nil
}

func (s RedisStore) RandomStateCode(ctx context.Context) (string, aerror.Error) {
	rand.Seed(time.Now().UnixNano())
	min := 200000
	max := 999999
	code := fmt.Sprintf("%d", rand.Intn(max-min+1)+min)

	return code, s.save(ctx, code)
}

func (s RedisStore) ValidSate(ctx context.Context, code string) aerror.Error {
	_, err := s.get(ctx, code)
	if err != nil {
		return err
	}

	return nil
}
