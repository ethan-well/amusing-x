package conf

import (
	"github.com/ItsWewin/superfactory/basicmatter"
	"log"
	"reflect"
)

var ConfIns = new(ConfigSection)

type ConfigSection struct {
	Addr string `config:"base:addr"`
	Port string `config:"base:http.port"`

	MysqlEuropaDatabase        string `config:"base:mysql.europa.database"`
	MysqlEuropaUsername        string `config:"base:mysql.europa.username"`
	MysqlEuropaPassword        string `config:"base:mysql.europa.password"`
	MysqlEuropaHost            string `config:"base:mysql.europa.host"`
	MysqlEuropaPort            string `config:"base:mysql.europa.port"`
	MysqlEuropaProtocol        string `config:"base:mysql.europa.protocol"`
	MysqlEuropaConnMaxLifetime int64  `config:"base:mysql.europa.connMaxLifetime"`
	MysqlEuropaMaxIdleConns    int    `config:"base:mysql.europa.maxIdleConns"`
	MysqlEuropaMaxOpenConns    int    `config:"base:mysql.europa.maxOpenConns"`

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

func (c *ConfigSection) MaterType() string {
	return basicmatter.MasterConfigBasicSection
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
