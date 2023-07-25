package user

import (
	service "amusingx.fit/amusingx/protos/user/service/user/proto"
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
	w.Add(2)
	go _ping(ctx, Client, &w)
	//go _githubOAth(ctx, Client, &w)
	//go _OAthInfo(ctx, Client, &w)
	//go _Authentication(ctx, Client, &w)
	go _IsLogin(ctx, Client, &w)

	w.Wait()
}

func _IsLogin(ctx context.Context, client service.UserServiceClient, w *sync.WaitGroup) {
	defer w.Done()

	in := &service.IsLoginRequest{SessionID: "5bd9bdc9-d2b3-4efd-9527-ac05938c38f9"}
	resp, err := client.IsLogin(ctx, in)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(resp)
}

func _ping(ctx context.Context, client service.UserServiceClient, w *sync.WaitGroup) {
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

		fmt.Println("resp", resp)
	}
}

func _githubOAth(ctx context.Context, client service.UserServiceClient, w *sync.WaitGroup) {
	defer w.Done()

	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	select {
	case <-ctx.Done():
		fmt.Println("time out")
		return
	case <-ticker.C:
		req := &service.OAuthLoginRequest{
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

func _OAthInfo(ctx context.Context, client service.UserServiceClient, w *sync.WaitGroup) {
	defer w.Done()

	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	select {
	case <-ctx.Done():
		fmt.Println("time out")
		return
	case <-ticker.C:
		req := &service.OAuthInfoRequest{
			Provider: "github",
			Service:  "pangu",
		}
		resp, err := client.OAuthInfo(context.Background(), req)
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println("resp", resp)
	}
}

func _Authentication(ctx context.Context, client service.UserServiceClient, w *sync.WaitGroup) {
	defer w.Done()

	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	select {
	case <-ctx.Done():
		fmt.Println("time out")
		return
	case <-ticker.C:
		req := &service.AuthenticationRequest{
			Server:  "pangu",
			UserId:  21,
			Actions: []string{"all-action", "no-action"},
		}
		resp, err := client.Authentication(context.Background(), req)
		if err != nil {
			fmt.Println(err)
		}

		for _, p := range resp.Result {
			fmt.Printf("action: %s, has permission: %t", p.Action, p.HasPermission)
		}
	}
}
