package xredis

import (
	"amusingx.fit/amusingx/services/europa/conf"
)

func Mock() {
	conf.Mock()

	InitRedis(conf.ConfIns.RedisAddr, conf.ConfIns.RedisPassword, conf.ConfIns.RedisDB)
}
