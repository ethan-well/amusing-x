package conf

import (
	"github.com/ItsWewin/superfactory/logger"
	"runtime"
)

func Mock() {
	TestingGetFileName()
	
 	//conf := config.NewBasicYamlConf()
	//conf.Unmarshal(Conf, "")
}


func TestingGetFileName() {
	_, filename, _, _ := runtime.Caller(0)
	logger.Infof("file name: %s", filename)
}
