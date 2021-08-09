package conf

import (
	"log"
	"reflect"
)

var Conf = new(Config)

type Config struct {
	Addr string `config:"base:addr"`
	Port string `config:"base:http.port"`

	MysqlAmusingriskDatabase        string `config:"base:mysql.Amusingrisk.database"`
	MysqlAmusingriskUsername        string `config:"base:mysql.Amusingrisk.username"`
	MysqlAmusingriskPassword        string `config:"base:mysql.Amusingrisk.password"`
	MysqlAmusingriskHost            string `config:"base:mysql.Amusingrisk.host"`
	MysqlAmusingriskPort            string `config:"base:mysql.Amusingrisk.port"`
	MysqlAmusingriskProtocol        string `config:"base:mysql.Amusingrisk.protocol"`
	MysqlAmusingriskConnMaxLifetime int64  `config:"base:mysql.Amusingrisk.connMaxLifetime"`
	MysqlAmusingriskMaxIdleConns    int    `config:"base:mysql.Amusingrisk.maxIdleConns"`
	MysqlAmusingriskMaxOpenConns    int    `config:"base:mysql.Amusingrisk.maxOpenConns"`

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

func (c *Config) GetAddr() string {
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
