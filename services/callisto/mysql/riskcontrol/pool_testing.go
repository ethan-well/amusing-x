package riskcontrol

import (
	"amusingx.fit/amusingx/services/callisto/conf"
)

func Mock() {
	conf.Mock()

	InitMySQL()
}
