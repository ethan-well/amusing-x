package amusingxwebapi

import "amusingx.fit/amusingx/services/webApi/conf"

func Mock() {
	conf.Mock()
	InitMySQL()
}
