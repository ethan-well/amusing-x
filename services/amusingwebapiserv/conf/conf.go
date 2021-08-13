package conf

import (
	"log"
	"reflect"
)

var Conf = new(Config)

type Config struct {
	Addr string `config:"base:addr"`
	Port string `config:"base:http.port"`

	MysqlAmusinguserDatabase        string `config:"base:mysql.amusinguser.database"`
	MysqlAmusinguserUsername        string `config:"base:mysql.amusinguser.username"`
	MysqlAmusinguserPassword        string `config:"base:mysql.amusinguser.password"`
	MysqlAmusinguserHost            string `config:"base:mysql.amusinguser.host"`
	MysqlAmusinguserPort            string `config:"base:mysql.amusinguser.port"`
	MysqlAmusinguserProtocol        string `config:"base:mysql.amusinguser.protocol"`
	MysqlAmusinguserConnMaxLifetime int64  `config:"base:mysql.amusinguser.connMaxLifetime"`
	MysqlAmusinguserMaxIdleConns    int    `config:"base:mysql.amusinguser.maxIdleConns"`
	MysqlAmusinguserMaxOpenConns    int    `config:"base:mysql.amusinguser.maxOpenConns"`

	RedisAddr     string `config:"base:redis.addr"`
	RedisPassword string `config:"base:redis.password"`
	RedisDB       int    `config:"base:redis.db"`

	SessionStore               string `config:"session:store"`
	SessionStoreRedisAddr      string `config:"session:store.redis.addr"`
	SessionStoreRedisPassword  string `config:"session:store.redis.password"`
	SessionStoreRedisDB        int    `config:"session:store.redis.db"`
	SessionStoreRedisKeyPrefix string `config:"session:store.redis.key.prefix"`

	RiskServiceRPCAddress string `config:"rpc:amusingriskservice.address"`

	RPCNetwork            string `rpc:network`
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
