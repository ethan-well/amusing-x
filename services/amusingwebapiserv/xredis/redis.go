package xredis

import (
	"context"
	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

var Client *redis.Client

func InitRedis(addr, password string, db int) {
	Client = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})
}

func CloseRedis() {
	if Client != nil {
		Client.Close()
	}
}
