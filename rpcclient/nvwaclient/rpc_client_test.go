package ganymede

import (
	helloworldserv "amusingx.fit/amusingx/protos/nvwa/service/helloword"
	"context"
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestInitClient(t *testing.T) {
	InitClient("localhost:8080")

	ctx, _ := context.WithTimeout(context.Background(), 100*time.Second)

	w := sync.WaitGroup{}
	w.Add(1)
	go _echo(ctx, Client, &w)

	w.Wait()
}

func _echo(ctx context.Context, client helloworldserv.GreeterClient, w *sync.WaitGroup) {
	defer w.Done()

	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	select {
	case <-ctx.Done():
		fmt.Println("time out")
		return
	case <-ticker.C:
		resp, err := client.SayHello(context.Background(), &helloworldserv.HelloRequest{Name: "wei.wei"})
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println("resp", resp)
	}
}
