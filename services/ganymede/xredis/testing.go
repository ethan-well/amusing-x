package xredis

import (
	"amusingx.fit/amusingx/services/ganymede/conf"
)

func Mock() {
	conf.Mock()

	InitRedis(conf.ConfSect.RedisAddr, conf.ConfSect.RedisPassword, conf.ConfSect.RedisDB)
}
