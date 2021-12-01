package mysql

import "C"
import (
	"amusingx.fit/amusingx/services/pluto/conf"
)

func Mock() {
	conf.Mock()
	InitMySQL()
}
