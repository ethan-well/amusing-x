package charon

import "amusingx.fit/amusingx/services/product/conf"

func Mock() {
	conf.Mock()
	InitMySQL()
}
