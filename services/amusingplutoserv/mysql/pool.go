package mysql

import (
	"amusingx.fit/amusingx/services/amusingplutoserv/conf"
	"github.com/ItsWewin/superfactory/db/mysql"
	"github.com/jmoiron/sqlx"
)

var PlutoDB *sqlx.DB

func InitMySQL() {
	mysqlDB := mysql.NewMysqlDB(conf.Conf.MysqlPlutoDB, conf.Conf.MysqlPlutoUsername, conf.Conf.MysqlPlutoPassword,
		conf.Conf.MysqlPlutoHost, conf.Conf.MysqlPlutoMaxOpenConns, conf.Conf.MysqlPlutoMaxIdleConns, conf.Conf.MysqlPlutoConnMaxLifetime)

	PlutoDB = mysqlDB.Connect()
}

func DisConnectMySQL() {
	if PlutoDB == nil {
		return
	}

	err := PlutoDB.Close()
	if err != nil {
		panic(err)
	}
}
