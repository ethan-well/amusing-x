package mysql

import "C"
import (
	"amusingx.fit/amusingx/services/amusingplutoserv/conf"
)

func Mock() {
	conf.Mock()
	InitMySQL()
}
