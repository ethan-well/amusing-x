package amusinguser

import (
	"amusingx.fit/amusingx/services/ganymede/conf"
	"github.com/ItsWewin/superfactory/db/mysql"
)

func Mock() {
	mysqlDB := mysql.NewMysqlDB(conf.ConfSect.MysqlAmusinguserDatabase, conf.ConfSect.MysqlAmusinguserUsername, conf.ConfSect.MysqlAmusinguserPassword,
		conf.ConfSect.MysqlAmusinguserHost, conf.ConfSect.MysqlAmusinguserMaxOpenConns, conf.ConfSect.MysqlAmusinguserMaxIdleConns, conf.ConfSect.MysqlAmusinguserConnMaxLifetime)

	AmusingUserDB = mysqlDB.Connect()
}
