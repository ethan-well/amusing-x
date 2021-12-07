package riskcontrol

import (
	"amusingx.fit/amusingx/services/callisto/conf"
	"github.com/ItsWewin/superfactory/db/mysql"
	"github.com/jmoiron/sqlx"
)

var CallistoDB *sqlx.DB

// 初始化数据库
func InitMySQL() {
	callisto := conf.Conf.Mysql.Callistodb
	mysqlDB := mysql.NewMysqlDB(callisto.DB, callisto.User, callisto.Password,
		callisto.Host, callisto.Port, callisto.Protocol, callisto.MaxOpenConns, callisto.MaxIdleConns, callisto.MaxLifeTime)

	CallistoDB = mysqlDB.Connect()
}

// 关闭数据库
func MysqlDisConnect() {
	if CallistoDB == nil {
		return
	}

	err := CallistoDB.Close()
	if err != nil {
		panic(err)
	}
}
