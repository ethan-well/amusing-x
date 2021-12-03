package amusinguser

import (
	"amusingx.fit/amusingx/services/ganymede/conf"
	"github.com/ItsWewin/superfactory/db/mysql"
	"github.com/jmoiron/sqlx"
)

var AmusingUserDB *sqlx.DB

// 初始化数据库
func InitMySQL() {
	mysqlDB := mysql.NewMysqlDB(conf.ConfSect.MysqlAmusinguserDatabase, conf.ConfSect.MysqlAmusinguserUsername, conf.ConfSect.MysqlAmusinguserPassword,
		conf.ConfSect.MysqlAmusinguserHost, conf.ConfSect.MysqlAmusinguserMaxOpenConns, conf.ConfSect.MysqlAmusinguserMaxIdleConns, conf.ConfSect.MysqlAmusinguserConnMaxLifetime)

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
