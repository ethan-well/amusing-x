package conf

import (
	amusing_x "amusingx.fit/amusingx"
	"github.com/ItsWewin/superfactory/basicmatter/config"
)

func Mock() {
	conf := config.NewBasicYamlConf()
	conf.Unmarshal(Conf, amusing_x.Root()+"/cmd/ganymede/config_local.yaml")
}
