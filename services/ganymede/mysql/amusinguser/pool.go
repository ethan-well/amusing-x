package amusinguser

import (
	"amusingx.fit/amusingx/services/ganymede/conf"
	"github.com/ItsWewin/superfactory/db/mysql"
	"github.com/ItsWewin/superfactory/logger"
	"github.com/jmoiron/sqlx"
)

var GanymedeDB *sqlx.DB

// 初始化数据库
func InitMySQL() {
	ganymedeDB := conf.Conf.Mysql.Ganymede
	mysqlDB := mysql.NewMysqlDB(ganymedeDB.DB, ganymedeDB.User, ganymedeDB.Password,
		ganymedeDB.Host, ganymedeDB.Port, ganymedeDB.Protocol, ganymedeDB.MaxOpenConns, ganymedeDB.MaxIdleConns, ganymedeDB.MaxLifeTime)
	
	logger.Infof("ganymedeDB: %s", logger.ToJson(ganymedeDB))
	
	GanymedeDB = mysqlDB.Connect()
}

// 关闭数据库
func MysqlDisConnect() {
	if GanymedeDB == nil {
		return
	}

	err := GanymedeDB.Close()
	if err != nil {
		panic(err)
	}
}
