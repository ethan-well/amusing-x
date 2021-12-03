package mysql

import (
	"amusingx.fit/amusingx/services/pluto/conf"
	"github.com/ItsWewin/superfactory/db/mysql"
	"github.com/ItsWewin/superfactory/logger"
	"github.com/jmoiron/sqlx"
)

var PlutoDB *sqlx.DB

func InitMySQL() {
	plutodb := conf.Conf.Mysql.Plutodb
	mysqlDB := mysql.NewMysqlDB(plutodb.DB, plutodb.User, plutodb.Password,
		plutodb.Host, plutodb.Port, plutodb.Protocol, plutodb.MaxOpenConns, plutodb.MaxIdleConns, plutodb.MaxLifeTime)

	logger.Infof("plutodb: %s", logger.ToJson(plutodb))

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
