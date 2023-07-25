package xredis

import (
	"amusingx.fit/amusingx/services/webApi/conf"
)

func Mock() {
	conf.Mock()

	redis := conf.Conf.Redis.RedisO
	InitRedis(redis.Addr, redis.Password, redis.DBNo)
}
