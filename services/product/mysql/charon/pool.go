package charon

import (
	"amusingx.fit/amusingx/services/product/conf"
	"github.com/ItsWewin/superfactory/db/mysql"
	"github.com/jmoiron/sqlx"
)

var CharonDB *sqlx.DB

// InitMySQL 初始化数据库
func InitMySQL() {
	pangu := conf.Conf.Mysql.Charon
	mysqlDB := mysql.NewMysqlDB(pangu.DB, pangu.User, pangu.Password,
		pangu.Host, pangu.Port, pangu.Protocol, pangu.MaxOpenConns, pangu.MaxIdleConns, pangu.MaxLifeTime)

	CharonDB = mysqlDB.Connect()
}

// MysqlDisConnect 关闭数据库
func MysqlDisConnect() {
	if CharonDB == nil {
		return
	}

	err := CharonDB.Close()
	if err != nil {
		panic(err)
	}
}
