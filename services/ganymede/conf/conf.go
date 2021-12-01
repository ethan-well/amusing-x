package conf

import (
	"log"
	"reflect"
)

var Conf = new(Config)

type Config struct {
	Addr string `config:"base:addr"`
	Port string `config:"base:http.port"`

	MysqlGanymedeDatabase        string `config:"base:mysql.ganymede.database"`
	MysqlGanymedeUsername        string `config:"base:mysql.ganymede.username"`
	MysqlGanymedePassword        string `config:"base:mysql.ganymede.password"`
	MysqlGanymedeHost            string `config:"base:mysql.ganymede.host"`
	MysqlGanymedePort            string `config:"base:mysql.ganymede.port"`
	MysqlGanymedeProtocol        string `config:"base:mysql.ganymede.protocol"`
	MysqlGanymedeConnMaxLifetime int64  `config:"base:mysql.ganymede.connMaxLifetime"`
	MysqlGanymedeMaxIdleConns    int    `config:"base:mysql.ganymede.maxIdleConns"`
	MysqlGanymedeMaxOpenConns    int    `config:"base:mysql.ganymede.maxOpenConns"`

	RedisAddr     string `config:"base:redis.addr"`
	RedisPassword string `config:"base:redis.password"`
	RedisDB       int    `config:"base:redis.db"`

	SessionStore               string `config:"session:store"`
	SessionStoreRedisAddr      string `config:"session:store.redis.addr"`
	SessionStoreRedisPassword  string `config:"session:store.redis.password"`
	SessionStoreRedisDB        int    `config:"session:store.redis.db"`
	SessionStoreRedisKeyPrefix string `config:"session:store.redis.key.prefix"`

	RiskServiceRPCAddress string `config:"rpc:amusingriskservice.address"`

	RPCNetwork            string `config:"rpc:network"`
	UserServiceRPCAddress string `config:"rpc:amusingxuserservice.address"`
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
