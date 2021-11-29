package xredis

import (
	"context"
	"github.com/go-redis/redis/v8"
	"time"
)

var ctx = context.Background()

var Client *redis.Client

func InitRedis(addr, password string, db int) {
	Client = redis.NewClient(&redis.Options{
		Addr:         addr,
		Password:     password,
		DB:           db,
		DialTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	})
}

func CloseRedis() {
	if Client != nil {
		Client.Close()
	}
}
