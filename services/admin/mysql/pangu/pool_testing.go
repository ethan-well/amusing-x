package pangu

import "amusingx.fit/amusingx/services/admin/conf"

func Mock() {
	conf.Mock()
	InitMySQL()
}
