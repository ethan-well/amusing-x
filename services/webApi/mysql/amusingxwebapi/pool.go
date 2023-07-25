package amusingxwebapi

import (
	"amusingx.fit/amusingx/services/webApi/conf"
	"github.com/ItsWewin/superfactory/db/mysql"
	"github.com/jmoiron/sqlx"
)

var AmusingUserDB *sqlx.DB

// 初始化数据库
func InitMySQL() {
	europa := conf.Conf.Mysql.Europadb
	mysqlDB := mysql.NewMysqlDB(europa.DB, europa.User, europa.Password,
		europa.Host, europa.Port, europa.Protocol, europa.MaxOpenConns, europa.MaxIdleConns, europa.MaxLifeTime)

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
