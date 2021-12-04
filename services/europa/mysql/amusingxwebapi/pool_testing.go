package amusingxwebapi

import (
	"amusingx.fit/amusingx/services/europa/conf"
	"github.com/ItsWewin/superfactory/db/mysql"
)

func Mock() {
	mysqlDB := mysql.NewMysqlDB(conf.ConfIns.MysqlAmusinguserDatabase, conf.ConfIns.MysqlAmusinguserUsername, conf.ConfIns.MysqlAmusinguserPassword,
		conf.ConfIns.MysqlAmusinguserHost, conf.ConfIns.MysqlAmusinguserMaxOpenConns, conf.ConfIns.MysqlAmusinguserMaxIdleConns, conf.ConfIns.MysqlAmusinguserConnMaxLifetime)

	AmusingUserDB = mysqlDB.Connect()
}
