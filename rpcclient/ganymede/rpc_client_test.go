package ganymede

import (
	ganymedeservice "amusingx.fit/amusingx/protos/ganymede/service"
	"context"
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestInitClient(t *testing.T) {
	InitClient("localhost:20005")

	ctx, _ := context.WithTimeout(context.Background(), 100*time.Second)

	w := sync.WaitGroup{}
	w.Add(1)
	//go _ping(ctx, Client, &w)
	go _githubOAth(ctx, Client, &w)

	w.Wait()
}

func _ping(ctx context.Context, client ganymedeservice.GanymedeServiceClient, w *sync.WaitGroup) {
	defer w.Done()

	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	select {
	case <-ctx.Done():
		fmt.Println("time out")
		return
	case <-ticker.C:
		resp, err := client.Pong(context.Background(), &ganymedeservice.BlankParams{})
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println("resp", resp)
	}
}

func _githubOAth(ctx context.Context, client ganymedeservice.GanymedeServiceClient, w *sync.WaitGroup) {
	defer w.Done()

	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	select {
	case <-ctx.Done():
		fmt.Println("time out")
		return
	case <-ticker.C:
		req := &ganymedeservice.OAuthLoginRequest{
			Provider: "github",
			Code:     "958b85c5a03c5d92f66e",
		}
		resp, err := client.OAuthLogin(context.Background(), req)
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println("resp", resp)
	}
}
