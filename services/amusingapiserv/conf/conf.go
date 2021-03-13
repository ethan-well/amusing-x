package conf

type Conf struct {
	Addr string `conf:"base:addr"`
	Port int `conf:"base:port"`
}

func Init(cf string) {

}
