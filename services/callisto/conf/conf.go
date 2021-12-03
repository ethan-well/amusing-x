package conf

import (
	"github.com/ItsWewin/superfactory/basicmatter"
	"log"
	"reflect"
)

var Conf = new(Config)

type Config struct {
	Addr string `config:"base:addr"`
	Port string `config:"base:http.port"`

	MysqlCallistoDatabase        string `config:"base:mysql.callisto.database"`
	MysqlCallistoUsername        string `config:"base:mysql.callisto.username"`
	MysqlCallistoPassword        string `config:"base:mysql.callisto.password"`
	MysqlCallistoHost            string `config:"base:mysql.callisto.host"`
	MysqlCallistoPort            string `config:"base:mysql.callisto.port"`
	MysqlCallistoProtocol        string `config:"base:mysql.callisto.protocol"`
	MysqlCallistoConnMaxLifetime int64  `config:"base:mysql.callisto.connMaxLifetime"`
	MysqlCallistoMaxIdleConns    int    `config:"base:mysql.callisto.maxIdleConns"`
	MysqlCallistoMaxOpenConns    int    `config:"base:mysql.callisto.maxOpenConns"`

	RedisAddr     string `config:"base:redis.addr"`
	RedisPassword string `config:"base:redis.password"`
	RedisDB       int    `config:"base:redis.db"`

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
