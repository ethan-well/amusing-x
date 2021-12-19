package pangu

import (
	"amusingx.fit/amusingx/services/pangu/conf"
	"github.com/ItsWewin/superfactory/db/mysql"
	"github.com/jmoiron/sqlx"
)

var PanguDB *sqlx.DB

// 初始化数据库
func InitMySQL() {
	pangu := conf.Conf.Mysql.Pangudb
	mysqlDB := mysql.NewMysqlDB(pangu.DB, pangu.User, pangu.Password,
		pangu.Host, pangu.Port, pangu.Protocol, pangu.MaxOpenConns, pangu.MaxIdleConns, pangu.MaxLifeTime)

	PanguDB = mysqlDB.Connect()
}

// 关闭数据库
func MysqlDisConnect() {
	if PanguDB == nil {
		return
	}

	err := PanguDB.Close()
	if err != nil {
		panic(err)
	}
}
