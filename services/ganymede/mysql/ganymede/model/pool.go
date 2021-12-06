package model

import (
	"amusingx.fit/amusingx/services/ganymede/conf"
	"github.com/ItsWewin/superfactory/db/mysql"
	"github.com/jmoiron/sqlx"
)

var GanymedeDB *sqlx.DB

// InitMySQL 初始化数据库
func InitMySQL() {
	ganymedeDB := conf.Conf.Mysql.Ganymede
	mysqlDB := mysql.NewMysqlDB(ganymedeDB.DB, ganymedeDB.User, ganymedeDB.Password,
		ganymedeDB.Host, ganymedeDB.Port, ganymedeDB.Protocol, ganymedeDB.MaxOpenConns, ganymedeDB.MaxIdleConns, ganymedeDB.MaxLifeTime)

	GanymedeDB = mysqlDB.Connect()
}

// MysqlDisConnect 关闭数据库
func MysqlDisConnect() {
	if GanymedeDB == nil {
		return
	}

	err := GanymedeDB.Close()
	if err != nil {
		panic(err)
	}
}
