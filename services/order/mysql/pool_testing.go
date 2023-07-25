package mysql

import "C"
import (
	"amusingx.fit/amusingx/services/order/conf"
)

func Mock() {
	conf.Mock()
	InitMySQL()
}
