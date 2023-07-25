package conf

import (
	"github.com/ItsWewin/superfactory/basicmatter"
	"log"
	"reflect"
)

var ConfSection = new(ConfigSection)

type ConfigSection struct {
	Addr       string `config:"base:addr"`
	Port       string `config:"base:http.port"`
	ServerName string `config:"base:server.name"`

	MysqlPlutoDB              string `config:"mysql:mysql.pluto.db"`
	MysqlPlutoUsername        string `config:"mysql:mysql.pluto.username"`
	MysqlPlutoPassword        string `config:"mysql:mysql.pluto.password"`
	MysqlPlutoHost            string `config:"mysql:mysql.pluto.host"`
	MysqlPlutoPort            string `config:"mysql:mysql.pluto.port"`
	MysqlPlutoProtocol        string `config:"mysql:mysql.pluto.protocol"`
	MysqlPlutoConnMaxLifetime int64  `config:"mysql:mysql.pluto.connMaxLifetime"`
	MysqlPlutoMaxIdleConns    int    `config:"mysql:mysql.pluto.maxIdleConns"`
	MysqlPlutoMaxOpenConns    int    `config:"mysql:mysql.pluto.maxOpenConns"`

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

func (c *ConfigSection) MaterType() string {
	return basicmatter.MasterConfigBasicYaml
}

func (c *ConfigSection) RunHttpServer() bool {
	return true
}

func (c *ConfigSection) HttpAddr() string {
	return c.Addr
}

func (c *ConfigSection) Print() {
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
