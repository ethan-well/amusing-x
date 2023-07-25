package xredis

import "amusingx.fit/amusingx/services/admin/conf"

func Mock() {
	conf.Mock()

	redis0 := conf.Conf.Redis.RedisO
	InitRedis(redis0.Addr, redis0.Password, redis0.DBNo)
}
