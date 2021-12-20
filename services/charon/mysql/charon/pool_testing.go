package charon

import "amusingx.fit/amusingx/services/charon/conf"

func Mock() {
	conf.Mock()
	InitMySQL()
}
