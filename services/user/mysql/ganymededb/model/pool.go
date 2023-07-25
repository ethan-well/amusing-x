package model

import (
	"amusingx.fit/amusingx/services/user/conf"
	"github.com/ItsWewin/superfactory/db/mysql"
	"github.com/jmoiron/sqlx"
)

var UserDB *sqlx.DB

// InitMySQL 初始化数据库
func InitMySQL() {
	userDB := conf.Conf.Mysql.User
	mysqlDB := mysql.NewMysqlDB(userDB.DB, userDB.User, userDB.Password,
		userDB.Host, userDB.Port, userDB.Protocol, userDB.MaxOpenConns, userDB.MaxIdleConns, userDB.MaxLifeTime)

	UserDB = mysqlDB.Connect()
}

// MysqlDisConnect 关闭数据库
func MysqlDisConnect() {
	if UserDB == nil {
		return
	}

	err := UserDB.Close()
	if err != nil {
		panic(err)
	}
}
