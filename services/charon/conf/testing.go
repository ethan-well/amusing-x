package conf

func Mock() {
	Conf = &Config{
		Addr:                      "10001",
		Port:                      "10001",
		RedisAddr:                 "localhost:6379",
		RedisPassword:             "",
		RedisDB:                   0,
		SessionStore:              "redis",
		SessionStoreRedisAddr:     "localhost:6379",
		SessionStoreRedisPassword: "",
		SessionStoreRedisDB:       2,
	}
}
