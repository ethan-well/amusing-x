package pangu

import "amusingx.fit/amusingx/services/pangu/conf"

func Mock() {
	conf.Mock()
	InitMySQL()
}
