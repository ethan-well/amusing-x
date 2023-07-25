package conf

import (
	amusing_x "amusingx.fit/amusingx"
	"github.com/ItsWewin/superfactory/basicmatter/config"
	"path"
)

func Mock() {
	conf := config.NewBasicYamlConf()
	err := conf.Unmarshal(Conf, path.Join(amusing_x.Root(), "cmd/europa/config_local.yaml"))
	if err != nil {
		panic(err)
	}
}
