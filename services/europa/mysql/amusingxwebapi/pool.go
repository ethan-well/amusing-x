package amusingxwebapi

import (
	"amusingx.fit/amusingx/services/europa/conf"
	"github.com/ItsWewin/superfactory/db/mysql"
	"github.com/jmoiron/sqlx"
)

var AmusingUserDB *sqlx.DB

// 初始化数据库
func InitMySQL() {
	mysqlDB := mysql.NewMysqlDB(conf.ConfIns.MysqlAmusinguserDatabase, conf.ConfIns.MysqlAmusinguserUsername, conf.ConfIns.MysqlAmusinguserPassword,
		conf.ConfIns.MysqlAmusinguserHost, conf.ConfIns.MysqlAmusinguserMaxOpenConns, conf.ConfIns.MysqlAmusinguserMaxIdleConns, conf.ConfIns.MysqlAmusinguserConnMaxLifetime)

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
