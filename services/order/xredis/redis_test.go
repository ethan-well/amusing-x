package xredis

import (
	"context"
	"testing"
	"time"
)

func TestHGet(t *testing.T) {
	InitRedis("127.0.0.1:6379", "", 0)
	ctx, _ := context.WithTimeout(context.Background(), 2*time.Second)
	result := Client.HGet(ctx, "pluto:service:promotion:books:inventory:cache:keydddd", "6115")
	count := -1
	err := result.Scan(&count)
	if err != nil {
		panic(err)
	}

	result = Client.HGet(context.Background(), "pluto:service:promotion:books:inventory:cache:keyddd", "6115dddd")
}
