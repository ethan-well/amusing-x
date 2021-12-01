package conf

func Mock() {
	Conf = &Config{
		Addr:                            "10001",
		Port:                            "10001",
		MysqlAmusingriskDatabase:        "amusingrisk",
		MysqlAmusingriskUsername:        "amusing",
		MysqlAmusingriskPassword:        "amusing111",
		MysqlAmusingriskHost:            "localhost",
		MysqlAmusingriskPort:            "3306",
		MysqlAmusingriskProtocol:        "tcp",
		MysqlAmusingriskConnMaxLifetime: 2 * 60,
		MysqlAmusingriskMaxIdleConns:    10,
		MysqlAmusingriskMaxOpenConns:    10,
		RedisAddr:                       "localhost:6379",
		RedisPassword:                   "",
		RedisDB:                         0,
		SessionStore:                    "redis",
		SessionStoreRedisAddr:           "localhost:6379",
		SessionStoreRedisPassword:       "",
		SessionStoreRedisDB:             2,
	}
}
