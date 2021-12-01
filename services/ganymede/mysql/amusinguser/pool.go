package amusinguser

import (
	"amusingx.fit/amusingx/services/ganymede/conf"
	"github.com/ItsWewin/superfactory/db/mysql"
	"github.com/jmoiron/sqlx"
)

var AmusingUserDB *sqlx.DB

// 初始化数据库
func InitMySQL() {
	mysqlDB := mysql.NewMysqlDB(conf.Conf.MysqlAmusinguserDatabase, conf.Conf.MysqlAmusinguserUsername, conf.Conf.MysqlAmusinguserPassword,
		conf.Conf.MysqlAmusinguserHost, conf.Conf.MysqlAmusinguserMaxOpenConns, conf.Conf.MysqlAmusinguserMaxIdleConns, conf.Conf.MysqlAmusinguserConnMaxLifetime)

	AmusingUserDB = mysqlDB.Connect()
}

// 关闭数据库
func MysqlDisConnect() {
	if AmusingUserDB == nil {
		return
	}

	err := AmusingUserDB.Close()
	if err != nil {
		panic(err)
	}
}
