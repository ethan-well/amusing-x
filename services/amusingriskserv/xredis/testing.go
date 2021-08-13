package xredis

import (
	"amusingx.fit/amusingx/services/amusingriskserv/conf"
)

func Mock() {
	conf.Mock()

	InitRedis(conf.Conf.RedisAddr, conf.Conf.RedisPassword, conf.Conf.RedisDB)
}
