package amusingxwebapi

import "amusingx.fit/amusingx/services/europa/conf"

func Mock() {
	conf.Mock()
	InitMySQL()
}
