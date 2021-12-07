package xredis

import (
	"amusingx.fit/amusingx/services/europa/conf"
)

func Mock() {
	conf.Mock()

	redis := conf.Conf.Redis.RedisO
	InitRedis(redis.Addr, redis.Password, redis.DBNo)
}
