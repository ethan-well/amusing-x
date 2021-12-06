package conf

import (
	amusing_x "amusingx.fit/amusingx"
	"github.com/ItsWewin/superfactory/basicmatter/config"
	"github.com/ItsWewin/superfactory/logger"
	"runtime"
)

func Mock() {
	TestingGetFileName()

	conf := config.NewBasicYamlConf()
	conf.Unmarshal(Conf, amusing_x.Root()+"/cmd/ganymede/config_local.yaml")
}

func TestingGetFileName() {
	_, filename, _, _ := runtime.Caller(0)
	logger.Infof("file name: %s", filename)
}
