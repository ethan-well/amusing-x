package amusingriskcontrol

import (
	"amusingx.fit/amusingx/services/amusingriskcontrolserv/conf"
	"github.com/ItsWewin/superfactory/db/mysql"
)

func Mock() {
	mysqlDB := mysql.NewMysqlDB(conf.Conf.MysqlAmusingriskDatabase, conf.Conf.MysqlAmusingriskUsername, conf.Conf.MysqlAmusingriskPassword,
		conf.Conf.MysqlAmusingriskHost, conf.Conf.MysqlAmusingriskMaxOpenConns, conf.Conf.MysqlAmusingriskMaxIdleConns, conf.Conf.MysqlAmusingriskConnMaxLifetime)

	AmusingRiskDB = mysqlDB.Connect()
}
