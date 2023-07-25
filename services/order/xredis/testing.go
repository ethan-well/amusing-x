package xredis

import "amusingx.fit/amusingx/services/order/conf"

func Mock() {
	conf.Mock()

	InitRedis(conf.Conf.Redis.RedisO.Addr, conf.Conf.Redis.RedisO.Password, conf.Conf.Redis.RedisO.DBNo)
}
