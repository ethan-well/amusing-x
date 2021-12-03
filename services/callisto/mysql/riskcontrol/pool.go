package riskcontrol

import (
	"amusingx.fit/amusingx/services/callisto/conf"
	"github.com/ItsWewin/superfactory/db/mysql"
	"github.com/ItsWewin/superfactory/logger"
	"github.com/jmoiron/sqlx"
)

var AmusingRiskDB *sqlx.DB

// 初始化数据库
func InitMySQL() {
	riskDB := conf.Conf.Mysql.RiskDB
	mysqlDB := mysql.NewMysqlDB(riskDB.DB, riskDB.User, riskDB.Password,
		riskDB.Host, riskDB.Port, riskDB.Protocol, riskDB.MaxOpenConns, riskDB.MaxIdleConns, riskDB.MaxLifeTime)

	logger.Infof("riskDB: %s", logger.ToJson(riskDB))

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
