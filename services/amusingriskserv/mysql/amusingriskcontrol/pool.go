package amusingriskcontrol

import (
	"amusingx.fit/amusingx/services/amusingriskserv/conf"
	"github.com/ItsWewin/superfactory/db/mysql"
	"github.com/jmoiron/sqlx"
)

var AmusingRiskDB *sqlx.DB

// 初始化数据库
func InitMySQL() {
	mysqlDB := mysql.NewMysqlDB(conf.Conf.MysqlAmusingriskDatabase, conf.Conf.MysqlAmusingriskUsername, conf.Conf.MysqlAmusingriskPassword,
		conf.Conf.MysqlAmusingriskHost, conf.Conf.MysqlAmusingriskMaxOpenConns, conf.Conf.MysqlAmusingriskMaxIdleConns, conf.Conf.MysqlAmusingriskConnMaxLifetime)

	AmusingRiskDB = mysqlDB.Connect()
}

// 关闭数据库
func MysqlDisConnect() {
	if AmusingRiskDB == nil {
		return
	}

	err := AmusingRiskDB.Close()
	if err != nil {
		panic(err)
	}
}
