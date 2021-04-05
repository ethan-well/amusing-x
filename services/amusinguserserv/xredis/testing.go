package xredis

import (
	"amusingx.fit/amusingx/services/amusinguserserv/conf"
	"github.com/go-redis/redis/v8"
)

func Mock() {
	conf.Mock()

	Client = redis.NewClient(&redis.Options{
		Addr:     conf.Conf.RedisAddr,
		Password: conf.Conf.RedisPassword,
		DB:       conf.Conf.RedisDB,
	})
}
