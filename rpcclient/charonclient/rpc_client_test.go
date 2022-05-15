package charonclient

import (
	"context"
	"sync"
	"testing"
	"time"
)

func ClientMock(t *testing.T) {
	InitClient(":20004")

	ctx, _ := context.WithTimeout(context.Background(), 100*time.Second)

	w := sync.WaitGroup{}
	w.Add(1)
	go _ping(ctx, Client, &w)
	//go _githubOAth(ctx, Client, &w)
	//go _OAthInfo(ctx, Client, &w)
	//go _Authentication(ctx, Client, &w)
	// go _attributeWithSubProduct(ctx, Client, &w)

	w.Wait()
}
