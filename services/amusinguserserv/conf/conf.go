package conf

import (
	"log"
	"reflect"
)

var Conf = new(Config)

type Config struct {
	Addr string `config:"base:addr"`
	Port string `config:"base:http.port"`

	MysqlAmusinguserDatabase string `config:"base:mysql.amusinguser.database"`
	MysqlAmusinguserUsername string `config:"base:mysql.amusinguser.username"`
	MysqlAmusinguserPassword string `config:"base:mysql.amusinguser.password"`
	MysqlAmusinguserHost string `config:"base:mysql.amusinguser.host"`
	MysqlAmusinguserPort string  `config:"base:mysql.amusinguser.port"`
	MysqlAmusinguserProtocol string `config:"base:mysql.amusinguser.protocol"`
	MysqlAmusinguserConnMaxLifetime int64 `config:"base:mysql.amusinguser.connMaxLifetime"`
	MysqlAmusinguserMaxIdleConns int `config:"base:mysql.amusinguser.maxIdleConns"`
	MysqlAmusinguserMaxOpenConns int `config:"base:mysql.amusinguser.maxOpenConns"`
}

func (c *Config) GetAddr() string {
	return c.Addr
}

func (c *Config) Print() {
	rv := reflect.ValueOf(*c)
	rt := reflect.TypeOf(*c)

	log.Println("==================== service config ====================")
	for i := 0; i < rv.NumField(); i ++ {
		log.Printf("%v:%v\n", rt.Field(i).Name, rv.Field(i))
	}
	log.Println("==================== service config ====================")
}