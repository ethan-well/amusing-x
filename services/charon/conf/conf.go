package conf

import (
	"github.com/ItsWewin/superfactory/basicmatter"
	"log"
	"reflect"
)

var Conf = new(Config)

type Config struct {
	Addr       string `config:"base:addr"`
	Port       string `config:"base:http.port"`
	ServerName string `config:"base:server.name"`

	RedisAddr     string `config:"redis:redis.addr"`
	RedisPassword string `config:"redis:redis.password"`
	RedisDB       int    `config:"redis:redis.db"`

	SessionStore               string `config:"session:store"`
	SessionStoreRedisAddr      string `config:"session:store.redis.addr"`
	SessionStoreRedisPassword  string `config:"session:store.redis.password"`
	SessionStoreRedisDB        int    `config:"session:store.redis.db"`
	SessionStoreRedisKeyPrefix string `config:"session:store.redis.key.prefix"`

	RPCNetwork string `config:"rpc:network"`
	RPCAddress string `config:"rpc:address"`
}

func (c *Config) MaterType() string {
	return basicmatter.MasterConfigBasicSection
}

func (c *Config) HttpAddr() string {
	return c.Addr
}

func (c *Config) Print() {
	rv := reflect.ValueOf(*c)
	rt := reflect.TypeOf(*c)

	log.Println("==================== service config ====================")
	for i := 0; i < rv.NumField(); i++ {
		if rt.Field(i).Name == "Port" {
			log.Printf("%v:%v\n", rt.Field(i).Name, rv.Field(i))
		}
	}
	log.Println("==================== service config ====================")
}
