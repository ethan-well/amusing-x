package conf

func Mock() {
	Conf = &Config{
		Addr:                      "10001",
		Port:                      "10001",
		MysqlPlutoDB:              "plutodb",
		MysqlPlutoUsername:        "amusing",
		MysqlPlutoPassword:        "amusing111",
		MysqlPlutoHost:            "localhost",
		MysqlPlutoPort:            "3306",
		MysqlPlutoProtocol:        "tcp",
		MysqlPlutoConnMaxLifetime: 2 * 60,
		MysqlPlutoMaxIdleConns:    10,
		MysqlPlutoMaxOpenConns:    10,
		RedisAddr:                 "localhost:6379",
		RedisPassword:             "",
		RedisDB:                   0,
		SessionStore:              "redis",
		SessionStoreRedisAddr:     "localhost:6379",
		SessionStoreRedisPassword: "",
		SessionStoreRedisDB:       2,
	}
}
