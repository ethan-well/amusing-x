package xredis

import (
	"context"
	"testing"
)

func TestHGet(t *testing.T) {
	InitRedis("localhost:6379", "", 0)
	result := Client.HGet(context.Background(), "pluto:service:promotion:books:inventory:cache:keydddd", "6115")
	count := -1
	err := result.Scan(&count)
	if err != nil {
		panic(err)
	}

	result = Client.HGet(context.Background(), "pluto:service:promotion:books:inventory:cache:keyddd", "6115dddd")

}
