package riskcontrol

import (
	"amusingx.fit/amusingx/services/risk/conf"
)

func Mock() {
	conf.Mock()

	InitMySQL()
}
