package amusinguser

import (
	"amusingx.fit/amusingx/services/ganymede/conf"
	"github.com/ItsWewin/superfactory/db/mysql"
)

func Mock() {
	mysqlDB := mysql.NewMysqlDB(conf.Conf.MysqlAmusinguserDatabase, conf.Conf.MysqlAmusinguserUsername, conf.Conf.MysqlAmusinguserPassword,
		conf.Conf.MysqlAmusinguserHost, conf.Conf.MysqlAmusinguserMaxOpenConns, conf.Conf.MysqlAmusinguserMaxIdleConns, conf.Conf.MysqlAmusinguserConnMaxLifetime)

	AmusingUserDB = mysqlDB.Connect()
}
