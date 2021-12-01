package charonclient

import (
	"amusingx.fit/amusingx/protos/charons/service"
	"context"
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestInitClient(t *testing.T) {
	InitClient(":11004")

	ctx, _ := context.WithTimeout(context.Background(), 100*time.Second)

	w := sync.WaitGroup{}
	w.Add(1)
	go _ping(ctx, Client, &w)

	w.Wait()
}

func _ping(ctx context.Context, client service.CharonServClient, w *sync.WaitGroup) {
	defer w.Done()

	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	select {
	case <-ctx.Done():
		fmt.Println("time out")
		return
	case <-ticker.C:
		resp, err := client.Pong(context.Background(), &service.BlankParams{})
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println(resp)
	}
}
